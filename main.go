package main

import (
	"fmt"
	"os"

	"github.com/Noreen-py/Bander/core"
	"github.com/Noreen-py/Bander/interpreter/lexer"
	"github.com/Noreen-py/Bander/interpreter/token"
)

var source string = `
goto
`

func main() {
	l := *lexer.Lex(source)
	tok := l.NextToken()
	for {
		tok = l.NextToken()
		fmt.Fprintf(os.Stdin, "%+v\n", tok)
		if tok.Type == token.EOF {
			break
		}
	}

	App := core.Core{}
	App.Window()

}
