package main

import (
	"fmt"
	"os"
	"spell_script_lang/lexer"
	"spell_script_lang/repl"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		// File mode
		filename := os.Args[1]
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Lexing file: %s\n", filename)
		fmt.Println(strings.Repeat("=", 40))

		l := lexer.NewLexer(string(data))
		for {
			tok := l.NextToken()
			fmt.Printf("Position %s - %s: %s\n", tok.Pos, tok.Type, tok.Literal)
			if tok.Type == lexer.EOF {
				break
			}
		}
	} else {
		// REPL mode
		fmt.Println("Welcome to My Language REPL!")
		fmt.Println("Type your code or 'exit' to quit")
		fmt.Println(strings.Repeat("-", 40))
		repl.Start()
	}
}
