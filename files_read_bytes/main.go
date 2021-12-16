package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s/n", byteSlice)

	file, err = os.Open("files_read_bytes/main.go")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data as string: %s\n", data)
	log.Printf("Number of bytes read: %d\n", len((data)))

	data, err = ioutil.ReadFile("test.txt") // Handles openning and closing
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}
