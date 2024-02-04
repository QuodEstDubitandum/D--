package parser


func (p *Parser) NoPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) NextTokenError(expectToken lexer.TokenType, gotToken lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", lexer.ASCIIMap[expectToken], lexer.ASCIIMap[gotToken])
	p.errors = append(p.errors, msg)
}