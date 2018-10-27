package interpreter

import (
	"fmt"
)

type MultInterpreter struct {
	lexer        Lexer
	currentToken *Token
}

func NewMultInterpreter(lexer Lexer) *MultInterpreter {
	return &MultInterpreter{lexer: lexer}
}

func (m *MultInterpreter) factor() int64 {
	val := m.currentToken.Value.(int64)
	m.eat(INTEGER)

	return val
}

func (m *MultInterpreter) eat(tokenType TokenType) error {
	if m.currentToken.Type != tokenType {
		return fmt.Errorf("Invalid Token found: %v", tokenType)
	}

	var err error
	m.currentToken, err = m.lexer.GetNextToken()

	return err
}

func (m *MultInterpreter) Expr() (int64, error) {
	var err error
	m.currentToken, err = m.lexer.GetNextToken()

	if err != nil {
		return 0, err
	}

	result := m.factor()

	if err != nil {
		return 0, err
	}

	for m.currentToken.Type == MULTIPLY || m.currentToken.Type == DIVIDE || m.currentToken.Type == PLUS || m.currentToken.Type == MINUS {
		if m.currentToken.Type == MULTIPLY {
			m.eat(MULTIPLY)
			result *= m.factor()
		}

		if m.currentToken.Type == DIVIDE {
			m.eat(DIVIDE)
			result /= m.factor()
		}

		if m.currentToken.Type == PLUS {
			m.eat(PLUS)
			result += m.factor()
		}

		if m.currentToken.Type == MINUS {
			m.eat(MINUS)
			result -= m.factor()
		}
	}

	return result, nil
}
