package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		// we can write or read connection
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// conn is writer interface.
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		// $ telnet localhost 8080

		conn.Close()
	}
}
