package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := writter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d\n", bytesWritten)

	bytesAvailable := writter.Available()
	log.Printf("Bytes available: %d\n", bytesAvailable)

	bytesWritten, err = writter.WriteString("\nJust a random string")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := writter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	writter.Flush()
}
