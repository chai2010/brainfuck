package brainfuck_test

import (
	"os"

	"github.com/chai2010/brainfuck"
)

func ExampleHi() {
	brainfuck.New("++++++++++[>++++++++++<-]>++++.+.", os.Stdin, os.Stdout).Run()
	// Output:
	// hi
}
