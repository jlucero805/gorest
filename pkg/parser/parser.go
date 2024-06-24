package parser

import (
	"errors"
	"lexer/tokens"
	"log"
)

type Parser struct {
    Tokens []tokens.Token
    Current int
    Meta map[string]string
    Body string
}

func Parse(tokens []tokens.Token) Parser {
    parser := Parser{
        Tokens: tokens,
        Current: 0,
        Meta: make(map[string]string),
        Body: "",
    }

    return parser
}

func (p *Parser) parse() {
    for !p.isEof() {
        if p.current().TokenType == tokens.IDENT {
            if p.current().Lexem == "BODY" {
            } else {
                ident, err := p.parseIdent()
                if err != nil {
                    log.Fatal(err)
                }

                _, err = p.parseColon()
                if err != nil {
                    log.Fatal(err)
                }

                str, err := p.parseString()
                if err != nil {
                    log.Fatal(err)
                }

                p.Meta[ident.Lexem] = str.Lexem
            }
        }
    }
}

func (p *Parser) parseIdent() (tokens.Token, error) {
    if p.current().TokenType == tokens.IDENT {
        token := p.current()
        p.inc()
        return token, nil
    }
    return tokens.Token{}, errors.New("Failed to parse ident")
}

func (p *Parser) parseColon() (tokens.Token, error) {
    if p.current().TokenType == tokens.COLON {
        token := p.current()
        p.inc()
        return token, nil
    }

    return tokens.Token{}, errors.New("Failed to parse colon")
}

func (p *Parser) parseString() (tokens.Token, error) {
    if p.current().TokenType == tokens.STRING {
        token := p.current()
        p.inc()
        return token, nil
    }
    return tokens.Token{}, errors.New("Failed to parse string")
}

func (p *Parser) inc() {
    p.Current += 1
}

func (p *Parser) current() tokens.Token {
    return p.Tokens[p.Current]
}

func (p *Parser) isEof() bool {
    return p.Current >= len(p.Tokens)
}
