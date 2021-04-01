package readerwriter

import (
	"io"
	"log"
	"net"
)

// echo is a handler function that echos recieved data
func echo(conn net.Conn) {
	defer conn.Close()

	//Create a buffer to store recieved data
	b := make([]byte, 512)
	for {
		//Recieve data via conn.Read into buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("client disconnected")
			break
		}
		if err != nil {
			log.Println("unexpected error")
			break
		}
		log.Printf("Recieved %d byes: %s\n", size, string(b))

		//send data via conn.Write
		log.Println("Writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
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
