package parser

import (
  "monkey/ast"
  "monkey/lexer"
  "monkey/token"
)

type Parser struct {
  l *lexer.Lexer
  
  curToken token.Token
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser { l: l }
  
  // read 2 tokens
  p.nextToken()
  p.nextToken()
  
  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.nextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}