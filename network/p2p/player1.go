package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleIncomingMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Received from Player 2: %s\n", text)
	}
}

func main() {
	// player1做服务端监听
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	defer listener.Close()

	fmt.Println("Player 1 waiting for a match on 127.0.0.1:9001")

	conn, _ := listener.Accept()
	defer conn.Close()

	fmt.Println("Player 1 connected to Player 2")

	// player1的接收
	go handleIncomingMessages(conn)

	// player1的发送
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintf(conn, "%s\n", text)
	}
}
