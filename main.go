package main

import (
	"fmt"
	"os"
)

type Parser struct {
	ip      int
	program []byte
	cursor  uint32
}

func main() {
	file, _ := os.ReadFile(os.Args[1])
	t_file := string(file)
	var p = Parser{
		ip:      0,
		cursor:  0,
		program: make([]byte, 3000),
	}
	for p.ip < len(t_file) {
		t := t_file[p.ip]
		p.evalToken(rune(t))
		p.ip++
	}
}

func (p Parser) evalToken(t rune) {
	switch t {
	case '.':
		fmt.Printf("%d", p.program[p.cursor])
	case '+':
		p.program[p.cursor]++
	case '-':
		p.program[p.cursor]--
	case '>':
		p.cursor++
	case '<':
		p.cursor--
	}
}
