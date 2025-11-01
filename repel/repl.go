package repl

import (
	"bufio"
	"fmt"
	"os"
	"spell_script_lang/lexer"
)

const BANNER = "Welcome to Magic World."

const PROMPT = ">>> "

// start funtion REPL
func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for {
			token := l.NextToken()
			if token.Type == lexer.EOF {
				break
			}
			fmt.Printf("%s: %s \n", token.Type, token.Literal)
		}
	}
}
