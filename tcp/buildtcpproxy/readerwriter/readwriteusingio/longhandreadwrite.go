package readerwriter

//build a reader and writer as an example
//keepfor reference and do no use

import (
	"fmt"
	"log"
	"os"
)

// Fooreader defines an io.Reader to read from stdin.
type FooReader struct{}

// Read reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}

// Write writes data to Stdout
func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func useExplicitInsOuts() {
	// Instantiate reqder and writer
	var (
		reader FooReader
		writer FooWriter
	)

	// Create buffer to hold input/output
	input := make([]byte, 4096)

	// Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}
	fmt.Printf("Read %d bytes from stdin\n", s)

	//Use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("wrote %d byes to stdout\n", s)
}
