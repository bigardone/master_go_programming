package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")

	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err) // the same as above
	}

	err = os.Truncate("a.txt", 0) // empty file

	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
	if err != nil {
		log.Fatal(err)
	}
	p("File name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())
	p("Last modified", fileInfo.ModTime())
	p("Is directory", fileInfo.IsDir())
	p("Permisions", fileInfo.Mode())

	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			p("The file does not exist")
		}
	}

	// oldPath := "a.txt"
	// newPath := "aaa.txt"
	// err = os.Rename(oldPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}
}
