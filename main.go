package main

import (
	"fmt"

	"github.com/Noreen-py/Bander/core"
	"github.com/Noreen-py/Bander/interpreter/lexer"
)

var APP bool = false

func main() {
	lexed := lexer.LexedResult()
	fmt.Println(lexed)
	if APP {
		App := core.Core{}
		App.Window()
	}

}
