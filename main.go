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
	var (
		out     byte
		pointer int64
		isLoop  bool
		loop_p  int64
	)

	loop_stack := make([]byte, 3000)
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

			case "]":
				loop_p = int64(loop_stack[len(loop_stack)-1])
				loop_stack = loop_stack[:len(loop_stack)-1]
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
