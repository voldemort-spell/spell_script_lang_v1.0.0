package spellscriptlang

import (
	"fmt"
	"os"
	"spell_script_lang/lexer"
	repl "spell_script_lang/repel"
)

func main() {
	if len(os.Args) > 1 {

		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s \n", err)
			os.Exit(1)
		}

		l := lexer.NewLexer(string(data))
		for {
			token := l.NextToken()
			if token.Type == lexer.EOF {
				break
			}
			fmt.Printf("%s: %s \n", token.Type, token.Literal)
		}
	} else {
		repl.Start()
	}
}
