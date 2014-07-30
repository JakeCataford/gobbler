package main

import (
	"database"
	"encoding/json"
	"log"
	"net"
)

func main() {
	db, err := database.Connect("localhost")
	tcp, err := net.Listen("tcp", ":3050")
	defer closeTcpConnection(tcp)

	if err != nil {
		panic(err)
	}

	waitForConnection(tcp)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	dec := json.NewDecoder(conn)

	var message map[string]interface{}
	dec.Decode(&event)

	log.Printf("\rRecieved : %s", event)
	go persistEvent(event)
}

func waitForConnection(tcp net.Listener) {
	for {
		conn, err := tcp.Accept()

		if err != nil {
			panic(err)
			continue
		}

		go handleConnection(conn)
	}
}

func closeTcpConnection(tcp net.Listener) {
	err := tcp.Close()
	if err != nil {
		panic(err)
	}
}
