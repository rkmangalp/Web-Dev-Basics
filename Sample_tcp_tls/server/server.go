package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn *tls.Conn) {
	// close the connection when function ends
	defer conn.Close()
	fmt.Println("Handling connection")

	// Read data from the connection
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		log.Printf("Error reading the data: %v", err)
		return
	}

	// Parse the request (simple string parsing)
	request := string(data[:n])
	log.Printf("Received request: %s", request)

	// Respond with a simple message
	response := fmt.Sprintf("Hey Rk! Hi from server")
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("error sending the response: %v", err)
	}
}

func main() {

	// Load server certificate and private key
	certs, err := tls.LoadX509KeyPair("certs/server.crt", "certs/decrypted_server.key")
	if err != nil {
		log.Fatalf("Error loading certificates: %v", err)
	}

	// Set up TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certs},
	}

	// Create TCP listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting serve: %v", err)
	}
	defer listener.Close()

	log.Println("server listening on port 8080")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: %v", err)
			continue
		}

		// Upgrade to TLS connection
		tlsConn := tls.Server(conn, tlsConfig)

		// Handle the connection
		go handleConnection(tlsConn)
	}

}
