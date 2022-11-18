package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// for writing bytes
	size, err := f.Write([]byte("writing to file"))

	// for writing string
	// size, err := f.WriteString("writing to file")
	if err != nil {
		panic(err)
	}

	fmt.Printf("File created of sucess! The size to file in Bytes is :%d \n\n", size)
	f.Close()

	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("The data of File.Txt is: %s\n\n", file)

	fileBuffer, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	// for large files
	reader := bufio.NewReader(fileBuffer)
	buffer := make([]byte, 5)

	for {
		bytesForBuffer, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:bytesForBuffer]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
