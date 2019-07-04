// Copyright © 2016 ChaiShushan (chaishushan{AT}gmail.com).
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package brainfuck

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Machine struct {
	mem  [30000]byte
	code string
	pos  int
	pc   int
	r    io.ByteReader
	w    io.Writer
}

func New(code string, r io.Reader, w io.Writer) *Machine {
	if r == nil {
		r = os.Stdin
	}
	if w == nil {
		w = os.Stdout
	}
	return &Machine{
		code: code,
		r:    bufio.NewReader(r),
		w:    w,
	}
}

func (p *Machine) Init(code string, r io.Reader, w io.Writer) *Machine {
	if r == nil {
		r = os.Stdin
	}
	if w == nil {
		w = os.Stdout
	}

	p.code = code
	p.Reset()
	return p
}

func (p *Machine) Reset() *Machine {
	for i, _ := range p.mem {
		p.mem[i] = 0
	}
	p.pos = 0
	p.pc = 0
	return p
}

func (p *Machine) Run() error {
	for ; p.pc != len(p.code); p.pc++ {
		switch x := p.code[p.pc]; x {
		case '>':
			p.pos++
		case '<':
			p.pos--
		case '+':
			p.mem[p.pos]++
		case '-':
			p.mem[p.pos]--
		case '[':
			if p.mem[p.pos] == 0 {
				p.loop(1)
			}
		case ']':
			if p.mem[p.pos] != 0 {
				p.loop(-1)
			}
		case '.':
			fmt.Fprintf(p.w, "%c", p.mem[p.pos])
		case ',':
			c, err := p.r.ReadByte()
			if err != nil {
				return err
			}
			p.mem[p.pos] = c
		}
	}
	return nil
}

func (p *Machine) loop(inc int) {
	for i := inc; i != 0; p.pc += inc {
		switch p.code[p.pc+inc] {
		case '[':
			i++
		case ']':
			i--
		}
	}
}
