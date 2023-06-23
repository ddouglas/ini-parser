package lexer

import (
	"github.com/ddouglas/sfs-parser/lexer/lexertoken"
)

func Initialize(name, input string) *Lexer {
	return &Lexer{
		Name:   name,
		Input:  input,
		State:  lexBegin,
		Tokens: make(chan lexertoken.Token, 1),
	}
}
