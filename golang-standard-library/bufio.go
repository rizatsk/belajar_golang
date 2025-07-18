package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Reader on file
	input := strings.NewReader("This is long string\nwith new line \n")
	reader := bufio.NewReader(input) // Tujuan read

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// Writer
	writer := bufio.NewWriter(os.Stdout) // Tujuan write
	writer.WriteString("Hello world\n")
	writer.WriteString("Selamat belajar golang\n")
	writer.Flush()
}
