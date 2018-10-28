package interpreter

import (
	"fmt"
)

type MultInterpreter struct {
	lexer        Lexer
	currentToken *Token
}

func NewMultInterpreter(lexer Lexer) *MultInterpreter {
	token, _ := lexer.GetNextToken()
	return &MultInterpreter{lexer: lexer, currentToken: token}
}

func (m *MultInterpreter) factor() int64 {
	if m.currentToken.Type == INTEGER {
		val := m.currentToken.Value.(int64)

		m.eat(INTEGER)
		return val
	}

	m.eat(LPAREN)
	result, _ := m.Expr()
	m.eat(RPAREN)

	return result
}

func (m *MultInterpreter) eat(tokenType TokenType) error {
	if m.currentToken.Type != tokenType {
		return fmt.Errorf("Invalid Token found: %v", tokenType)
	}

	var err error
	m.currentToken, err = m.lexer.GetNextToken()

	return err
}

func (m *MultInterpreter) term() int64 {
	result := m.factor()

	for m.currentToken.Type == MULTIPLY || m.currentToken.Type == DIVIDE {
		if m.currentToken.Type == MULTIPLY {
			m.eat(MULTIPLY)
			result *= m.factor()
		}

		if m.currentToken.Type == DIVIDE {
			m.eat(DIVIDE)
			result /= m.factor()
		}
	}

	return result
}

func (m *MultInterpreter) Expr() (int64, error) {
	result := m.term()

	for m.currentToken.Type == PLUS || m.currentToken.Type == MINUS {
		if m.currentToken.Type == PLUS {
			m.eat(PLUS)
			result += m.term()
		}

		if m.currentToken.Type == MINUS {
			m.eat(MINUS)
			result -= m.term()
		}
	}

	return result, nil
}
