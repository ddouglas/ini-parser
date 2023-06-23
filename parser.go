package parser

import (
	"fmt"
	"strings"

	"github.com/ddouglas/sfs-parser/lexer/lexer"
	"github.com/ddouglas/sfs-parser/lexer/lexertoken"
)

func isEof(t lexertoken.Token) bool {
	return t.Type == lexertoken.TOKEN_EOF
}

func Parse(filename, input string) File {

	o := File{
		Name:     filename,
		Sections: make([]Section, 0),
	}

	var token lexertoken.Token
	var value string

	section := Section{}
	var key string

	fmt.Println("starting lexer and parser....")

	l := lexer.Initialize(o.Name, input)

	for {
		token = l.NextToken()

		if isEof(token) {
			o.Sections = append(o.Sections, section)
			break
		}

		value = token.Value

		if token.Type != lexertoken.TOKEN_VALUE {
			value = strings.TrimSpace(value)
		}

		switch token.Type {
		case lexertoken.TOKEN_SECTION:
			if len(section.Pairs) > 0 {
				o.Sections = append(o.Sections, section)
			}

			key = ""

			section.Name = value
			section.Pairs = make([]KeyValue, 0)
		case lexertoken.TOKEN_KEY:
			key = value
		case lexertoken.TOKEN_VALUE:
			section.Pairs = append(section.Pairs, KeyValue{
				Key:   key,
				Value: value,
			})
			key = ""
		}
	}

	fmt.Println("Parser has completed")
	return o

}

type File struct {
	Name     string    `json:"name"`
	Sections []Section `json:"sections"`
}

type Section struct {
	Name  string     `json:"name"`
	Pairs []KeyValue `json:"pairs"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
