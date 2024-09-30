/* repl is a function that starts a Read-Eval-Print Loop (REPL) for the given interpreter.
It continuously reads input from the user, evaluates it using the interpreter, and prints the result.
The loop terminates when the user inputs an exit command or an error occurs.*/

package repl

import (
	"bufio"
	"dumblang/dumblang/lexer"
	"dumblang/dumblang/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

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
