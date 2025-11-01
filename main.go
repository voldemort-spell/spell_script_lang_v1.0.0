package spellscriptlang

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {

		filename := os.Args[1]
		_, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s \n", err)
			os.Exit(1)
		}
	}
}
