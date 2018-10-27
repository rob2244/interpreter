package interpreter

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	text         string
	pos          int
	currentToken *Token
}

// NewInterpreter creates a new interpreter with the given text
func NewInterpreter(text string) *Interpreter {
	return &Interpreter{text: text, pos: 0}
}

// GetNextToken returns the next token and increments the counter
func (i *Interpreter) getNextToken() (*Token, error) {
	i.skipWhitespace()

	if i.pos > len(i.text)-1 {
		return NewToken(EOF, ""), nil
	}

	_, err := strconv.ParseInt(i.currentChar(), 10, 64)

	switch {
	case i.currentChar() == "+":
		i.pos++
		return NewToken(PLUS, i.currentChar()), nil
	case i.currentChar() == "-":
		i.pos++
		return NewToken(MINUS, i.currentChar()), nil
	case err == nil:
		return NewToken(INTEGER, i.integer()), nil
	default:
		return nil, fmt.Errorf("%s is an unknown token type", i.currentChar())
	}
}

// Eat takes in a token type and checks if the current token is of the same type
// if it is it advances to the next token, if not it returns an error
func (i *Interpreter) eat(t TokenType) error {
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
	i.eat(INTEGER)

	return token.Value.(int64)
}

func (i *Interpreter) Expr() (int64, error) {
	var err error
	i.currentToken, err = i.getNextToken()

	if err != nil {
		return 0, err
	}

	result := i.term()

	for i.currentToken.Type == PLUS || i.currentToken.Type == MINUS {
		token := i.currentToken
		if token.Type == PLUS {
			i.eat(PLUS)
			result += i.term()
		}

		if token.Type == MINUS {
			i.eat(MINUS)
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
