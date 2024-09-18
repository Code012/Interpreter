package repl

import (
	"bufio" // buffered I/O
	"fmt"
	"io"
	"monkey-go/lexer"
	"monkey-go/token"
)

const PROMPT = ">> "

// Initialise REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// Infinite loop to interact continuosly
	for {

		fmt.Printf(PROMPT)

		// Scan next token
		scanned := scanner.Scan()
		if !scanned {
			// Exit is end of input reached
			return
		}

		// Get scanned token as string
		line := scanner.Text()

		// Create lexer for current line and then tokenise and print each token
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
