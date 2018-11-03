package interpreter

import (
	"fmt"
	"robinseitz/interpreter/lexer"
)

type MultInterpreter struct {
	lexer        lexer.Lexer
	currentToken *lexer.Token
}

func NewMultInterpreter(lexer lexer.Lexer) *MultInterpreter {
	token, _ := lexer.GetNextToken()
	return &MultInterpreter{lexer: lexer, currentToken: token}
}

func (m *MultInterpreter) factor() int64 {
	if m.currentToken.Type == lexer.INTEGER {
		val := m.currentToken.Value.(int64)

		m.eat(lexer.INTEGER)
		return val
	}

	m.eat(lexer.LPAREN)
	result, _ := m.Expr()
	m.eat(lexer.RPAREN)

	return result
}

func (m *MultInterpreter) eat(tokenType lexer.TokenType) error {
	if m.currentToken.Type != tokenType {
		return fmt.Errorf("Invalid Token found: %v", tokenType)
	}

	var err error
	m.currentToken, err = m.lexer.GetNextToken()

	return err
}

func (m *MultInterpreter) term() int64 {
	result := m.factor()

	for m.currentToken.Type == lexer.MULTIPLY || m.currentToken.Type == lexer.DIVIDE {
		if m.currentToken.Type == lexer.MULTIPLY {
			m.eat(lexer.MULTIPLY)
			result *= m.factor()
		}

		if m.currentToken.Type == lexer.DIVIDE {
			m.eat(lexer.DIVIDE)
			result /= m.factor()
		}
	}

	return result
}

func (m *MultInterpreter) Expr() (int64, error) {
	result := m.term()

	for m.currentToken.Type == lexer.PLUS || m.currentToken.Type == lexer.MINUS {
		if m.currentToken.Type == lexer.PLUS {
			m.eat(lexer.PLUS)
			result += m.term()
		}

		if m.currentToken.Type == lexer.MINUS {
			m.eat(lexer.MINUS)
			result -= m.term()
		}
	}

	return result, nil
}
