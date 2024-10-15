package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/mdp/qrterminal/v3"
)

// GetLocalIP retrieves the local IP address of the current machine.
func GetLocalIP() (string, error) {
	// We use a dummy UDP connection to a remote address (Google's public DNS server)
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Get the local address from the connection
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// GenerateQRCode generates and displays a QR code for the provided data.
func GenerateQRCode(data string) {
	// Use default settings of qrterminal to display the QR code in the terminal
	qrterminal.Generate(data, qrterminal.L, os.Stdout)
}

func main() {
	// Parse command-line flags for port and https usage
	port := flag.Int("port", 80, "Port to append to the local address. Default is 80.")
	useHttps := flag.Bool("https", false, "Use HTTPS instead of HTTP.")
	flag.Parse()

	localIP, err := GetLocalIP()
	if err != nil {
		fmt.Printf("Error getting local IP: %v\n", err)
		return
	}

	// Determine protocol based on the flag
	protocol := "http"
	if *useHttps {
		protocol = "https"
	}

	// Format the full address
	address := fmt.Sprintf("%s://%s:%d", protocol, localIP, *port)
	fmt.Printf("Generated address: %s\n", address)

	// Generate and display the QR code
	GenerateQRCode(address)
}
