package interpreter

import (
	"fmt"
	"robinseitz/interpreter/lexer"
	"strconv"
)

type Interpreter struct {
	text         string
	pos          int
	currentToken *lexer.Token
}

// NewInterpreter creates a new interpreter with the given text
func NewInterpreter(text string) *Interpreter {
	return &Interpreter{text: text, pos: 0}
}

// GetNextToken returns the next token and increments the counter
func (i *Interpreter) getNextToken() (*lexer.Token, error) {
	i.skipWhitespace()

	if i.pos > len(i.text)-1 {
		return lexer.NewToken(lexer.EOF, ""), nil
	}

	_, err := strconv.ParseInt(i.currentChar(), 10, 64)

	switch {
	case i.currentChar() == "+":
		i.pos++
		return lexer.NewToken(lexer.PLUS, i.currentChar()), nil
	case i.currentChar() == "-":
		i.pos++
		return lexer.NewToken(lexer.MINUS, i.currentChar()), nil
	case err == nil:
		return lexer.NewToken(lexer.INTEGER, i.integer()), nil
	default:
		return nil, fmt.Errorf("%s is an unknown token type", i.currentChar())
	}
}

// Eat takes in a token type and checks if the current token is of the same type
// if it is it advances to the next token, if not it returns an error
func (i *Interpreter) eat(t lexer.TokenType) error {
	if i.currentToken.Type == t {
		token, err := i.getNextToken()

		if err != nil {
			return err
		}

		i.currentToken = token
		return nil
	}

	return fmt.Errorf("%s incorrect token type", t)
}

func (i *Interpreter) term() int64 {
	token := i.currentToken
	i.eat(lexer.INTEGER)

	return token.Value.(int64)
}

func (i *Interpreter) Expr() (int64, error) {
	var err error
	i.currentToken, err = i.getNextToken()

	if err != nil {
		return 0, err
	}

	result := i.term()

	for i.currentToken.Type == lexer.PLUS || i.currentToken.Type == lexer.MINUS {
		token := i.currentToken
		if token.Type == lexer.PLUS {
			i.eat(lexer.PLUS)
			result += i.term()
		}

		if token.Type == lexer.MINUS {
			i.eat(lexer.MINUS)
			result -= i.term()
		}
	}

	return result, nil
}

func (i *Interpreter) skipWhitespace() {
	if i.pos > len(i.text)-1 {
		return
	}

	for txt := string(i.text[i.pos]); txt == " "; txt = i.currentChar() {
		i.pos = i.pos + 1
	}
}

func (i *Interpreter) integer() int64 {
	numString := ""

	for i.pos < len(i.text) {
		_, err := strconv.ParseInt(i.currentChar(), 10, 64)

		if err != nil {
			break
		}

		numString += i.currentChar()
		i.pos++
	}

	result, _ := strconv.ParseInt(numString, 10, 64)
	return result
}

func (i *Interpreter) currentChar() string {
	if i.pos-1 > len(i.text) {
		return ""
	}

	return string(i.text[i.pos])
}
