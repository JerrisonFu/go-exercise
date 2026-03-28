package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

type Message struct {
	Type    uint8
	Length  uint32
	Payload []byte
}

func StartTCPServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	addr, _ := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return
	}

	conn.Write(buffer[:n])
}

func TCPClient(port int, message string) (string, error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return "", err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(message))
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func UDPEchoServer(port int) string {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buffer)
	return string(buffer[:n])
}

func UDPEchoClient(port int, message string) string {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf("127.0.0.1:%d", port))
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	conn.Write([]byte(message))
	buffer := make([]byte, 1024)
	n, _, _ := conn.ReadFromUDP(buffer)
	return string(buffer[:n])
}

func NetworkInterfaces() []string {
	interfaces, _ := net.Interfaces()
	names := make([]string, 0, len(interfaces))
	for _, iface := range interfaces {
		names = append(names, iface.Name)
	}
	return names
}

func ResolveHost() string {
	ips, _ := net.LookupHost("localhost")
	if len(ips) > 0 {
		return ips[0]
	}
	return ""
}

func PingLikeCheck(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func EncodeMessage(msg Message) []byte {
	data := make([]byte, 5+len(msg.Payload))
	data[0] = byte(msg.Type)
	binary.BigEndian.PutUint32(data[1:5], msg.Length)
	copy(data[5:], msg.Payload)
	return data
}

func DecodeMessage(data []byte) Message {
	return Message{
		Type:    data[0],
		Length:  binary.BigEndian.Uint32(data[1:5]),
		Payload: data[5:],
	}
}
