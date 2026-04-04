package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// reader

	/*
		file, err := os.Open("letters.txt")
		if err != nil {
			panic(err)
		}
	*/

	file := strings.NewReader("Hello world!")

	n, err := countAlphabets(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Letters: %d\n", n)

	// writer
	fileOut, err := os.Create("writing.txt")
	if err != nil {
		panic(err)
	}
	defer fileOut.Close()

	n, err = writeString("Hello World!", fileOut)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Writing bytes: %d\n", n)
}

func countAlphabets(reader io.Reader) (int, error) {
	count := 0
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		for _, c := range buffer[:n] {
			if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') {
				count++
			}
		}

		switch err {
		case nil:
		case io.EOF:
			return count, nil
		default:
			return 0, err
		}
	}
}

func writeString(str string, writer io.Writer) (int, error) {
	n, err := writer.Write([]byte(str))
	if err != nil {
		return 0, fmt.Errorf("error occurred while writting: %w", err)
	}
	return n, nil
}
