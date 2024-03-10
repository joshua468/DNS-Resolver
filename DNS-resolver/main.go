package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Check if domain name argument is provided
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <domain>")
		os.Exit(1)
	}

	// Get domain name from command line argument
	domain := os.Args[1]

	// Resolve domain name to IP addresses
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print resolved IP addresses
	fmt.Printf("IP addresses for %s:\n", domain)
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
