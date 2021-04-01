package main

// shortened echo server using go packages

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that echos recieved data
func echo(conn net.Conn) {
	defer conn.Close()

	//copy data from io.Reader to io.Writer via io.Copy()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
	}

	// old refactored into new
	// reader := bufio.NewReader(conn)
	// s, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatalln("Unable to read data")
	// }
	// log.Printf("read %d bytes: %s", len(s), s)

	// log.Println("writing data")
	// writer := bufio.NewWriter(conn)
	// if _, err := writer.WriteString(s); err != nil {
	// 	log.Fatalln("Unable to write data")
	// }
	// writer.Flush()
}

func main() {
	// Bind to TCP por 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}
	log.Println("listening on 0.0.0.0:20080")
	for {
		//wait for connection. create net.Conn on connection established
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		//handle the connection. using go runtine for concurrencly
		go echo(conn)
	}

}
