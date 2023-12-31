package main

import (
	"bufio"
	"flag"
	"fmt"
	"lib/lib"
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
		lib.RunHost(connIP)
	} else {
		// go run main.go  <IP>
		fmt.Println("is guest")
		connIP := os.Args[1]
		lib.RunGuest(connIP)
	}

}

const port = "8080"

//RunHost takes an IP address
//and listen for connections on that IP

func RunHost(ip string) {
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

// RunGust take destination IP
// as argument and conect to that IP
func RunGuest(ip string) {

}
