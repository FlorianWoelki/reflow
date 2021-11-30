package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/florianwoelki/reflow/lexer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("reflow > ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		l := lexer.New(text)
		for i := 0; i < len(text); i++ {
			token := l.NextToken()
			fmt.Printf("[type: %s, literal: %s] ", token.Type, token.Literal)
		}
		fmt.Println()

		if strings.Compare("exit", text) == 0 {
			os.Exit(1)
		}
	}
}
