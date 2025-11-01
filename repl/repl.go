package repl

import (
	"bufio"
	"fmt"
	"os"
	"spell_script_lang/lexer"
	"strings"
)

const PROMPT = ">>> "

func Start() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(PROMPT)
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		l := lexer.NewLexer(line)
		for {
			tok := l.NextToken()
			if tok.Type == lexer.EOF {
				break
			}
			fmt.Printf("%s: %s\n", tok.Type, tok.Literal)
		}
	}
}
