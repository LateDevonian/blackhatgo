package main

//port forwarding to a proxy website
//page40
//run curl -i -X GET https://avrilejean.com to confirm you can comnnect to the website via proxy

import (
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "avrilejean.net:80")
	if err != nil {
		log.Fatalln("unable to connect to our unreachable host")
	}
	defer dst.Close()

	//run in goroutine to preventi io.Copy form blocking
	go func() {
		//copy our sources output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// copy ourdstinations output back to the source
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	//listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go handle(conn)
	}
}
