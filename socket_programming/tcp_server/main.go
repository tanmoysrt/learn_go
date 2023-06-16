package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func main()  {
	socket, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		panic(err)
	}
	defer socket.Close()
	fmt.Println("Server is running on port 3000")
	for {
		client, err := socket.Accept()
		if err != nil {
			panic(err)
		}
		go handlRead(client)
		go handleWrite(client)
	}
}



func handlRead(client net.Conn){
	var buf [128]byte
	for {
		n, err := client.Read(buf[:])
		if err != nil {
			break
		}
		msg := string(buf[:n])
		if strings.TrimSpace(msg) == "exit" {
			fmt.Println("Client is exiting...")
			break
		}
		fmt.Print(msg)
	}
	client.Close()
}

func handleWrite(client net.Conn){
	// var msg string
	// for {
	// 	fmt.Scanln(&msg)
	// 	msg = msg + "\n"
	// 	_, err := client.Write([]byte(msg))
	// 	if err != nil {
	// 		break
	// 	}
	// }

	// Disable write buffering
	if tcpConn, ok := client.(*net.TCPConn); ok {
		if err := tcpConn.SetWriteBuffer(0); err != nil {
			panic(err)
		}
	}

	start := time.Now()

	for i := 0; i < 40000; i++ {
		client.Write([]byte(strconv.Itoa(i)+"\n"))
	}

	fmt.Println("Time taken: ", time.Since(start))

	client.Close()
}