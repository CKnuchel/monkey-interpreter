package repl

import (
	"bufio"
	"fmt"
	"io"

	"monkey-interpreter/internal/lexer"
	"monkey-interpreter/internal/token"
)

// Package repl implements a Read-Eval-Print Loop (REPL) for the Monkey
// programming language. The REPL allows users to interactively enter Monkey
// code, which is then evaluated and the results printed back to the user.

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
