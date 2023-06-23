package lexer

import "parser/lexer/lexertoken"

func lexRightBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.RIGHT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_RIGHT_BRACKET)
	return lexBegin
}
