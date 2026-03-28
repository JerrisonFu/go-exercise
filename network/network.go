package network

import (
	"fmt"
	"sync"
)

func RunAll() {
	println("=== HTTP Server ===")
	server := CreateServer()
	defer server.Close()

	println("Server URL:", server.URL)

	println("\n=== HTTP Client GET ===")
	body, _ := HttpClientExample(server.URL)
	fmt.Println("Response:", body)

	println("\n=== HTTP Client POST ===")
	postBody, _ := HttpPostExample(server.URL)
	fmt.Println("POST Response:", postBody)

	println("\n=== TCP Server & Client ===")
	var wg sync.WaitGroup
	wg.Add(1)
	go StartTCPServer(18080, &wg)
	tcpResp, _ := TCPClient(18080, "Hello TCP")
	fmt.Println("TCP Response:", tcpResp)

	println("\n=== Network Interfaces ===")
	ifaces := NetworkInterfaces()
	fmt.Println("Interfaces:", ifaces)

	println("\n=== DNS Resolution ===")
	host := ResolveHost()
	fmt.Println("localhost resolves to:", host)

	println("\n=== Connection Check ===")
	reachable := PingLikeCheck("127.0.0.1", 18080)
	fmt.Println("Port 18080 reachable:", reachable)

	println("\n=== Message Encoding ===")
	msg := Message{Type: 1, Length: 5, Payload: []byte("hello")}
	encoded := EncodeMessage(msg)
	decoded := DecodeMessage(encoded)
	fmt.Printf("Encoded %d bytes -> Decoded: Type=%d, Payload=%s\n",
		len(encoded), decoded.Type, string(decoded.Payload))
}
