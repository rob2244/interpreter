package ast

import (
	"fmt"
	"robinseitz/interpreter/lexer"
)

type Parser struct {
	lexer        lexer.Lexer
	currentToken *lexer.Token
}

func NewParser(lexer lexer.Lexer) *Parser {
	token, _ := lexer.GetNextToken()
	return &Parser{lexer, token}
}

func (p *Parser) Parse() interface{} {
	return p.expr()
}

func (p *Parser) expr() interface{} {
	node := p.term()

	for p.currentToken.Type == lexer.PLUS || p.currentToken.Type == lexer.MINUS {
		token := p.currentToken

		if token.Type == lexer.PLUS {
			p.eat(lexer.PLUS)
		}

		if token.Type == lexer.MINUS {
			p.eat(lexer.MINUS)
		}

		node = NewBinOp(token, node, p.term())
	}

	return node
}

func (p *Parser) term() interface{} {
	node := p.factor()

	for p.currentToken.Type == lexer.MULTIPLY || p.currentToken.Type == lexer.DIVIDE {
		token := p.currentToken

		if token.Type == lexer.MULTIPLY {
			p.eat(lexer.MULTIPLY)
		}

		if token.Type == lexer.DIVIDE {
			p.eat(lexer.DIVIDE)
		}

		node = NewBinOp(token, node, p.factor())
	}

	return node
}

func (p *Parser) factor() interface{} {
	if p.currentToken.Type == lexer.INTEGER {
		token := p.currentToken
		p.eat(lexer.INTEGER)
		return NewNum(token)
	}

	p.eat(lexer.LPAREN)
	node := p.expr()
	p.eat(lexer.RPAREN)

	return node
}

func (p *Parser) eat(tokenType lexer.TokenType) error {
	if tokenType != p.currentToken.Type {
		return fmt.Errorf("Unexpected token found %s", tokenType)
	}

	p.currentToken, _ = p.lexer.GetNextToken()

	return nil
}
