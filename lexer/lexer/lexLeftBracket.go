package lexer

import "github.com/ddouglas/sfs-parser/lexer/lexertoken"

func lexLeftBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.LEFT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_LEFT_BRACKET)
	return lexSection
}
