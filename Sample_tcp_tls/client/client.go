package main

import (
	"crypto/tls"
	"io"
	"log"
)

func main() {
	// Load the server certificate (for verification)
	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/decrypted_server.key")
	if err != nil {
		log.Fatalf("Error loading certificate: %v", err)
	}

	// Set up TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}

	// Connect to the server (TCP)
	log.Println("Connecting to server on localhost:8080...")
	conn, err := tls.Dial("tcp", "localhost:8080", tlsConfig)
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	defer conn.Close()
	log.Println("Successfully connected to server.")

	// Send a simple request
	request := "Hey this is rk from client!."
	log.Printf("Sending request: %s", request)
	_, err = conn.Write([]byte(request))
	if err != nil {
		log.Fatalf("Error sending the request: %v", err)
	}
	log.Println("Request sent.")

	// Read the response from the server
	response := make([]byte, 1024)
	n, err := conn.Read(response)
	if err != nil && err != io.EOF {
		log.Fatalf("Error reading response: %v", err)
	}

	// Print the response
	log.Printf("Received response from server: %s", string(response[:n]))
}
