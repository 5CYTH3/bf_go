package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	out     byte
	pointer int64
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	t_file := strings.Split(string(file), "\n")
	ParseLine(t_file)
}

func ParseLine(lines []string) {
	var skipClosingLoop int = 0
	var loopStart int = -1
	var loopEnd int = -1
	// loop_stack := make([]byte, 3000)
	main := make([]byte, 3000)
	for _, i := range lines {
		lineArray := strings.Split(i, "")
		for i, char := range lineArray {
			switch char {
			case "<":
				pointer--
				i++
			case ">":
				pointer++
				i++
			case "+":
				main[pointer]++
				i++
			case "-":
				main[pointer]--
				i++
			case ".":
				out = main[pointer]
				fmt.Printf("out: %c\n", out)
				i++
			case "[":
				skipClosingLoop += 1
			case "]":
				if skipClosingLoop != 0 {
					skipClosingLoop -= 1
					continue
				}
				if loopStart == loopEnd {
					loopStart = -1
					loopEnd = -1
					continue
				}
				loop := lineArray[loopStart:loopEnd]
				for main[pointer] > 0 {
					executeWith(main, loop)
				}
			case ",":
				reader := bufio.NewReader(os.Stdin)
				b, _ := reader.ReadByte()
				main[pointer] = b
			}

		}

	}

}

func executeWith(main []byte, code []string) {

	for _, char := range code {
		if char == "+" {
			pointer++
		} else if char == "-" {
			pointer--
		} else if char == ">" {
			main[pointer]++
		} else if char == "<" {
			main[pointer]--
		} else if char == "." {
			out = main[pointer]
			fmt.Printf("out: %c\n", out)
		} else if char == "," {
			reader := bufio.NewReader(os.Stdin)
			b, _ := reader.ReadByte()
			main[pointer] = b
		} else if char == "[" {
			executeWith(main, code)
		} else if char == "]" {
			continue
		}
	}
}
