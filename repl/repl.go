package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Nearrivers/combo-parser/lexer"
	"github.com/Nearrivers/combo-parser/token"
)

const PROMPT string = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()

		if !scanned || scanner.Text() == "x" {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
