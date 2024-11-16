package main

import (
	"os"

	"github.com/Nearrivers/combo-parser/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
