package lexer

import (
	"parser/lexer/errors"
	"parser/lexer/lexertoken"
	"strings"
)

func lexSection(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_MISSING_RIGHT_BRACKET)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RIGHT_BRACKET) {
			lexer.Emit(lexertoken.TOKEN_SECTION)
			return lexRightBracket
		}

		lexer.Inc()
	}
}
