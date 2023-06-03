package main

import (
	"log"
	"net"
)

/*
ON client terminal in order to hit the server
curl http://localhost:1729
use -UseBasicParsing
*/
func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	if err != nil {
		log.Fatal(err)
	}

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHellow, World\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1729")

	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for client to conn")
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		log.Println("Client is connected!!")

		// Making it in seperate thread..
		// to make it multi-threadded tcp connection handling
		go do(conn)
	}

}
