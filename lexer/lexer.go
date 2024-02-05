package lexer

import (
	"io"
)

type Lexer struct {
	reader io.Reader
	buffer []byte
	currentPosition int
	evalPosition int
	ch byte
}

func NewLexer(r io.Reader) *Lexer {
	l := &Lexer{reader: r}
	l.nextChar()
	return l
}

func (l *Lexer) nextChar() {
    if l.isBufferFull() {
        l.loadBuffer()
    }

    if l.evalPosition < len(l.buffer) {
        l.ch = l.buffer[l.evalPosition]
    } else {
        l.ch = 0 // ASCII code for EOF
    }

    l.currentPosition = l.evalPosition
    l.evalPosition++
}

func (l *Lexer) peekChar() byte {
	if l.isBufferFull() {
		l.loadBuffer()
	} 
	
	if l.evalPosition < len(l.buffer) {
		return l.buffer[l.evalPosition]
	} else {
		return 0
	}
}

func (l *Lexer) isBufferFull() bool {
	return l.evalPosition >= len(l.buffer)
}

func (l *Lexer) loadBuffer() {
    buf := make([]byte, 1024)
    bytesRead, err := l.reader.Read(buf)
    if err != nil && err != io.EOF {
        panic(err)
    }
    l.buffer = buf[:bytesRead]
    l.evalPosition = 0
}

func (l *Lexer) NextToken() Token {
	var tok Token

	switch l.ch {
	case ' ':
		tok = Token{SPACE, string(l.ch), ""}
	case '\t':
		tok = Token{TAB, string(l.ch), ""}
	case '\n':
		tok = Token{NEWLINE, string(l.ch), ""}
	
	case '(':
		tok = Token{LRBRKT, string(l.ch), ""}
	case ')':
		tok = Token{RRBKRT, string(l.ch), ""}
	case '{':
		tok = Token{LCBRKT, string(l.ch), ""}
	case '}':
		tok = Token{RCBRKT, string(l.ch), ""}
	case '[':
		tok = Token{LSBRKT, string(l.ch), ""}
	case ']':
		tok = Token{RSBRKT, string(l.ch), ""}

	case '$':
		tok = Token{DOLLAR, string(l.ch), ""}
	case '%':
		tok = Token{PERCENT, string(l.ch), ""}
	
	case ';':
		tok = Token{SEMICOLON, string(l.ch), ""}
	case ',':
		tok = Token{COMMA, string(l.ch), ""}

	case '!':
		tok = l.readExlamation()
	case '=':
		tok = l.readAssign()
	case '>':
		tok = l.readGreater()
	case '<':
		tok = l.readSmaller()
	case '+':
		tok = l.readPlus()
	case '-':
		tok = l.readMinus()
	case '*':
		tok = l.readAsterisk()
	case '/':
		tok = l.readSlash()

	case 0:
		tok = Token{EOF, "", ""}

	default:
		if isLetter(l.ch) {
			tok = l.readIdentifier()
		} else if isDigit(l.ch) {
			tok = l.readNumber()
		} else {
			tok = Token{ILLEGAL, string(l.ch), "Invalid character"}
		}
	}

	l.nextChar()
	return tok
}

func (l *Lexer) readExlamation() Token {
	nextChar := l.peekChar()
	if nextChar == '=' {
		l.nextChar()
		return Token{NOT_EQUAL, "!=", ""}
	}
	return Token{EXCLAMATION, string(l.ch), ""}
}

func (l *Lexer) readAsterisk() Token {
	nextChar := l.peekChar()
	if nextChar == '*' {
		l.nextChar()
		return Token{DOUBLE_ASTERISK, "**", ""}
	}
	return Token{ASTERISK, string(l.ch), ""}
}

func (l *Lexer) readPlus() Token {
	nextChar := l.peekChar()
	if nextChar == '=' {
		l.nextChar()
		return Token{PLUS_EQUAL, "+=", ""}
	}
	return Token{PLUS, string(l.ch), ""}
}

func (l *Lexer) readMinus() Token {
	nextChar := l.peekChar()
	if nextChar == '>' {
		l.nextChar()
		return Token{MAPS_TO, "->", ""}
	}
	if nextChar == '=' {
		l.nextChar()
		return Token{MINUS_EQUAL, "-=", ""}
	}
	return Token{MINUS, string(l.ch), ""}
}

func (l *Lexer) readSmaller() Token {
	nextChar := l.peekChar()
	if nextChar == '=' {
		l.nextChar()
		return Token{SOR_EQUAL, "<=", ""}
	}
	return Token{SMALLER, string(l.ch), ""}
}

func (l *Lexer) readGreater() Token {
	nextChar := l.peekChar()
	if nextChar == '=' {
		l.nextChar()
		return Token{GOR_EQUAL, ">=", ""}
	}
	return Token{GREATER, string(l.ch), ""}
}


func (l *Lexer) readAssign() Token {
	nextChar := l.peekChar()
	if nextChar == '=' {
		l.nextChar()
		return Token{EQUAL, "==", ""}
	}
	return Token{ASSIGN, string(l.ch), ""}
}

func (l *Lexer) readSlash() Token {
	nextChar := l.peekChar()
	if nextChar == '/' {
		l.nextChar()
		return Token{COMMENT, "//", ""}
	}
	return Token{SLASH, string(l.ch), ""}
}

// TODO
func (l *Lexer) readIdentifier() Token {
	position := l.currentPosition
	var oldBufferString string

	for {
		if l.isBufferFull() {
			oldBufferString = string(l.buffer[position:l.currentPosition+1])
		}

		if isLetter(l.peekChar()){
			l.nextChar()
		} else {
			break
		}
	}

	if oldBufferString != "" {
		tokenType := determineIdentifierType(oldBufferString + string(l.buffer[:l.evalPosition]))
		return Token{tokenType, oldBufferString + string(l.buffer[:l.evalPosition]), ""}
	}
	
	tokenType := determineIdentifierType(string(l.buffer[position:l.evalPosition]))
	return Token{tokenType, string(l.buffer[position:l.evalPosition]), ""}
}

// TODO
func (l *Lexer) readNumber() Token {
	position := l.currentPosition
	var numberString string
	isFloat := false

	for {
		if l.isBufferFull() {
			numberString = string(l.buffer[position:l.currentPosition+1])
		}

		if l.peekChar() == '.' {
			if isFloat {
				return Token{ILLEGAL, string(l.buffer[position:l.evalPosition]), "Invalid number"}
			}
			isFloat = true
			l.nextChar()
		} else if isNumber(l.peekChar()) {
			l.nextChar()
		} else {
			break
		}
	}

	if numberString != "" {
		numberString += string(l.buffer[:l.evalPosition])
	}else{
		numberString = string(l.buffer[position:l.evalPosition])
	}

	if isFloat {
		tokenType := determineFloatType(numberString)
		return Token{tokenType, numberString, ""}
	}

	tokenType := determineIntegerType(numberString)
	return Token{tokenType, numberString, ""}
}

func determineIdentifierType(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return ID
}

// determine integer tokenType (UINT8, INT16, INT32, INT64)
// TODO
func determineIntegerType(number string) TokenType {
	return INT
}

// determine float tokenType (FLOAT32, FLOAT64)
// TODO
func determineFloatType(number string) TokenType {
	return FLOAT
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isNumber(ch byte) bool {
	return isDigit(ch) || ch == '.'
}