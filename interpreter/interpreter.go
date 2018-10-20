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
	if i.pos > len(i.text)-1 {
		return NewToken(EOF, ""), nil
	}

	char := string(i.text[i.pos])
	_, err := strconv.ParseInt(char, 10, 64)

	switch {
	case char == "+":
		i.pos++
		return NewToken(PLUS, char), nil
	case char == " ":
		return NewToken(WHITESPACE, char), nil
	case err == nil:
		i.pos++
		return NewToken(INTEGER, char), nil
	default:
		return nil, fmt.Errorf("%s is an unknown token type", char)
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

func (i *Interpreter) Expr() (int64, error) {
	nextToken, err := i.getNextToken()

	if err != nil {
		return 0, err
	}

	i.currentToken = nextToken
	left, _ := strconv.ParseInt(i.currentToken.Value, 10, 64)

	err = i.eat(INTEGER)

	if err != nil {
		return 0, err
	}

	err = i.eat(PLUS)

	if err != nil {
		return 0, err
	}

	right, _ := strconv.ParseInt(i.currentToken.Value, 10, 64)
	err = i.eat(INTEGER)

	if err != nil {
		return 0, err
	}

	return left + right, nil
}
