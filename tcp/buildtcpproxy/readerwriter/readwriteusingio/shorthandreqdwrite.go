package readwriter

//build a reader and writer as an example
//keepfor reference and do no use

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Fooreader defines an io.Reader to read from stdin.
type FooReader struct{}

// FooWriter defines an io.Writer to write to Stdout
type FooWriter struct{}

func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func useIoCopy() {
	var (
		reader FooReader
		writer FooWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("unable to read/write data")
	}
}
