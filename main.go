package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool
	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")

	flag.Parse()
	if isHost {
		// go run main.go -listen <IP>
		connIP := os.Args[2]
		fmt.Println("is host")
		runHost(connIP)
	} else {
		// go run main.go  <IP>
		connIP := os.Args[1]
		runGuest(connIP)
		// fmt.Println("is guest")
	}

}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenError := net.Listen("tcp", ipAndPort)
	if listenError != nil {
		log.Fatal("Error", listenError)
		os.Exit(1)
	}
	conn, acceeptError := listener.Accept()

	if acceeptError != nil {
		log.Fatal("Error:", acceeptError)
	}
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error:", readErr)
	}
	fmt.Println("Message recived", message)
}

func runGuest(ip string) {

}
