package parser

import (
	"fmt"

	"github.com/QuodEstDubitandum/D--/lexer"
)


func (p *Parser) NoPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", lexer.ASCIIMap[t])
	p.errors = append(p.errors, msg)
}

func (p *Parser) NextTokenError(expectToken lexer.TokenType, gotToken lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", lexer.ASCIIMap[expectToken], lexer.ASCIIMap[gotToken])
	p.errors = append(p.errors, msg)
}

func (p *Parser) MissingSpaceError(afterToken lexer.TokenType) {
	msg := fmt.Sprintf("no space found after %s", lexer.ASCIIMap[afterToken])
	p.errors = append(p.errors, msg)
}