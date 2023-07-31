package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

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

func RunGuest(ip string) {

}
