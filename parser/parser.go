package parser

import (
	"errors"
	"fmt"
	"log"

	"github.com/jlucero805/gorest/lexer/tokens"
	"github.com/jlucero805/gorest/parser/types"
)

type Parser struct {
    Tokens []tokens.Token
    Current int
    Meta map[string]string
    Body string
}

type ParseResult struct {
    Meta map[string]string
    Body string
}

func Parse(tokens []tokens.Token) ParseResult {
    parser := Parser{
        Tokens: tokens,
        Current: 0,
        Meta: make(map[string]string),
        Body: "",
    }

    parser.parse()

    return ParseResult{
        Meta: parser.Meta,
        Body: parser.Body,
    } 
}

func (p *Parser) parse() {
    for !p.isEof() {
        fmt.Print(p.current())
        if p.current().TokenType == tokens.IDENT {
            if p.current().Lexem == "BODY" {
                p.inc()

                _, err := p.parseColon()
                if err != nil {
                    log.Fatal(err)
                }

                object, err := p.parseJson()

                p.Body = object.Stringify()
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

func (p *Parser) parseJson() (types.Value, error) {
    if p.current().TokenType == tokens.STRING {
        t := types.String{Value: p.current().Lexem}
        p.inc()
        return t, nil
    } else if p.current().TokenType == tokens.NUMBER {
        t := types.Number{Value: p.current().Lexem}
        p.inc()
        return t, nil
    } else if p.current().TokenType == tokens.BOOLEAN {
        boolean := true
        if p.current().Lexem == "false" {
            boolean = false
        }
        p.inc()
        return types.Bool{Value: boolean }, nil
    } else if p.current().TokenType == tokens.L_BRACKET {
        p.inc()
        fields := make(map[string]types.Value)
        for p.current().TokenType != tokens.R_BRACKET {
            str, err := p.parseString()
            if err != nil {
                log.Fatal(err)
            }

            _, err = p.parseColon()
            if err != nil {
                log.Fatal(err)
            }

            object, err := p.parseJson()
            if err != nil {
                log.Fatal(err)
            }

            fields[str.Lexem] = object

            if p.current().TokenType == tokens.COMMA {
                p.inc()
            } else {
                break;
            }
        }
        p.inc()
        return types.Object{Fields: fields}, nil
    }
    return types.Null{}, errors.New("Error")
}

func (p *Parser) peek() int {
    if (p.Current + 1) >= len(p.Tokens) {
        return tokens.EOF
    } else {
        return p.Tokens[p.Current + 1].TokenType
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
    return p.Current >= len(p.Tokens) || p.current().TokenType == tokens.EOF
}
