package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleIncomingMessages2(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Received from Player 1: %s\n", text)
	}
}

func main() {
	// player2做客户端连接服务端
	conn, _ := net.Dial("tcp", "127.0.0.1:9001")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("fail to close connection, error=%v\n", err)
		}
	}(conn)

	fmt.Println("Player 2 connected to Player 1")

	// player2的接收
	go handleIncomingMessages2(conn)

	// player2的发送
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintf(conn, "%s\n", text)
		if err != nil {
			fmt.Printf("fail to send text, error=%v\n", err)
			return
		}
	}
}
