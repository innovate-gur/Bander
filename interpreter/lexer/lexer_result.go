package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Noreen-py/Bander/interpreter/token"
)

var fileData string = ""
var errCantReadFile = errors.New("cannot read current file")
var lexed_result []token.Token

func ReadFile() error {

	file, err := os.Open("main.bdr")
	if err != nil {
		return errCantReadFile
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fileData += string(fileScanner.Text()) + "\n"
	}
	return nil
}

func LexedResult() []token.Token {
	var l = &Lexer{}
	tok := l.NextToken()
	for {
		tok = l.NextToken()
		fmt.Println(tok)
		lexed_result = append(lexed_result, tok)

		if tok.Type == token.EOF {
			break
		}

	}
	return lexed_result
}
