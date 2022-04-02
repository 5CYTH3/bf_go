package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	t_file := strings.Split(string(file), "\n")
	ParseLine(t_file)
}

func ParseLine(lines []string) {
	var out byte = 49
	var pointer int64
	main := make([]byte, 3000)
	for _, i := range lines {
		lineArray := strings.Split(i, "")
		for i, char := range lineArray {
			switch char {
			case "<":
				pointer--
			case ">":
				pointer++
			case "+":
				main[pointer]++
			case "-":
				main[pointer]--
			case "[":
				if lineArray[i+1] == "-" {
					for main[pointer] != 0 {
						main[pointer]--
					}
				}
			case "]":
			case ".":
				out = main[pointer]
				fmt.Printf("out: %c\n", out)
			case ",":
				reader := bufio.NewReader(os.Stdin)
				b, _ := reader.ReadByte()

				main[pointer] = b
			}
		}

	}
}
