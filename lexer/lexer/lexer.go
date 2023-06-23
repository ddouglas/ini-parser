package lexer

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/ddouglas/sfs-parser/lexer/lexertoken"
)

type Lexer struct {
	Name   string
	Input  string
	Tokens chan lexertoken.Token
	State  LexFn

	Start int
	Pos   int
	Width int
}

func (this *Lexer) Backup() {
	this.Pos -= this.Width
}

func (this *Lexer) CurrentInput() string {
	return this.Input[this.Start:this.Pos]
}

func (this *Lexer) Dec() {
	this.Pos--
}

func (this *Lexer) Emit(tokenType lexertoken.TokenType) {
	this.Tokens <- lexertoken.Token{
		Type:  tokenType,
		Value: this.Input[this.Start:this.Pos],
	}
	this.Start = this.Pos
}

func (this *Lexer) Errorf(format string, args ...any) LexFn {
	this.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

func (this *Lexer) Inc() {
	this.Pos++
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Emit(lexertoken.TOKEN_EOF)
	}
}

func (this *Lexer) InputToEnd() string {
	return this.Input[this.Pos:]
}

func (this *Lexer) IsEOF() bool {
	return this.Pos >= len(this.Input)
}

func (this *Lexer) IsWhitespace() bool {
	ch, _ := utf8.DecodeRuneInString(this.Input[this.Pos:])
	return unicode.IsSpace(ch)
}

func (this *Lexer) Next() rune {
	if this.Pos >= utf8.RuneCountInString(this.Input) {
		this.Width = 0
		return lexertoken.EOF
	}

	result, width := utf8.DecodeRuneInString((this.Input[this.Pos:]))

	this.Width = width
	this.Pos += this.Width
	return result
}

func (this *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		case token := <-this.Tokens:
			return token
		default:
			this.State = this.State(this)
		}
	}
}

func (this *Lexer) Peek() rune {
	r := this.Next()
	this.Backup()
	return r
}

func (this *Lexer) Run() {
	for state := lexBegin; state != nil; {
		state = state(this)
	}

	this.Shutdown()
}

func (this *Lexer) Shutdown() {
	close(this.Tokens)
}

func (this *Lexer) SkipWhitespace() {
	for {
		ch := this.Next()

		if !unicode.IsSpace(ch) {
			this.Dec()
			break
		}

		if ch == lexertoken.EOF {
			this.Emit(lexertoken.TOKEN_EOF)
			break
		}
	}
}
