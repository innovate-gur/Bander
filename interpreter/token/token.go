package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	Undefined  = "UNDEFINED"
	Identifier = "IDENTIFIER"
	Number     = "NUMBER"
	SOF        = "SOF"
	EOF        = "EOF"
	EOL        = "EOL"

	LeftParen  = "("
	RightParen = ")"
	ONSTART    = "ONSTART"
	MOVE       = "MOVE"
	GOTO       = "GOTO"
)

var keywords = map[string]TokenType{
	"onstart": ONSTART,
	"move":    MOVE,
	"goto":    GOTO,
}

func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
