package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	timeout = time.Duration(3 * time.Second)
)

// Server type represents a server n¡in the docker infraestruture
type Server struct {
	Name string
	URL  string
	Port uint16
}

type Servers []Server

type ServerCh chan Server

func check(server Server, ready ServerCh, notReadyChannel ServerCh) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", server.URL, server.Port), timeout)

	if err != nil {
		notReadyChannel <- server
		return
	}

	ready <- server

	defer conn.Close()
}

// CheckServers checks if the servers are up for network communication and doesn't return the control until all of them do
func CheckServers(servers Servers) {
	readyChannel, notReadyChannel := make(ServerCh), make(ServerCh)

	defer close(readyChannel)
	defer close(notReadyChannel)

	start := time.Now()

	for _, server := range servers {
		go check(server, readyChannel, notReadyChannel)
	}

	for i := 0; i < len(servers); {
		select {
		case server := <-readyChannel:
			log.Printf("%s server ready. It takes %.2f seconds to start", server.Name, time.Until(start).Seconds())
			i++
		case server := <-notReadyChannel:
			go check(server, readyChannel, notReadyChannel)
		}
	}

	log.Println("Server checks finished")
}
