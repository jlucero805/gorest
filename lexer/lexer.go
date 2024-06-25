package lexer

import (
	"unicode"

	"github.com/jlucero805/gorest/lexer/grammar"
	"github.com/jlucero805/gorest/lexer/tokens"
)

type Lexer struct {
	chars  []rune
	start  int
	end    int
	tokens []tokens.Token
}

func Lex(source string) []tokens.Token {
    lexer := Lexer{
        chars: []rune(source),
        start: 0,
        end: 0,
        tokens: []tokens.Token{},
    }
    lexer.lex()
    return lexer.tokens
}

func (l *Lexer) lex() {
    for !l.isEof() {
        if grammar.Initial(l.currentChar()) {
            l.incCur()
            l.lexIdent()
        } else if l.currentChar() == ':' {
            l.tokens = append(l.tokens, tokens.Token{
                TokenType: tokens.COLON,
                Lexem: ":",
            })
            l.incSync()
        } else if l.currentChar() == '{' {
            l.tokens = append(l.tokens, tokens.Token{
                TokenType: tokens.L_BRACKET,
                Lexem: "",
            })
            l.incSync()
        } else if l.currentChar() == '}' {
            l.tokens = append(l.tokens, tokens.Token{
                TokenType: tokens.R_BRACKET,
                Lexem: "",
            })
            l.incSync()
        } else if l.currentChar() == ',' {
            l.tokens = append(l.tokens, tokens.Token{
                TokenType: tokens.COMMA,
                Lexem: ",",
            })
            l.incSync()
        } else if unicode.IsSpace(l.currentChar()) {
            l.incSync()
        } else if l.currentChar() == '"' {
            l.incSync()
            l.lexString()
        }
    }
    l.tokens = append(l.tokens, tokens.Token{
        TokenType: tokens.EOF,
        Lexem: "",
    })

}

func (l *Lexer) incCur() {
    l.end += 1
}

func (l *Lexer) currentChar() rune {
    return l.chars[l.end]
}

func (l *Lexer) isEof() bool {
    return l.end >= len(l.chars)
}

func (l *Lexer) lexIdent() {
    for !l.isEof() && grammar.Subsequent(l.currentChar()) {
        l.incCur()
    }
    lexem :=  l.slice()
    if lexem == "true" || lexem == "false" {
        l.tokens = append(l.tokens, tokens.Token{
            TokenType: tokens.BOOLEAN,
            Lexem: lexem,
        })
    } else {
        l.tokens = append(l.tokens, tokens.Token{
            TokenType: tokens.IDENT,
            Lexem: lexem,
        })
        l.sync()
    }
}

func (l *Lexer) sync() {
    l.start = l.end
}

func (l *Lexer) incSync() {
    l.end += 1
    l.start = l.end
}


func (l *Lexer) lexString() {
    for !l.isEof() && !(l.currentChar() == '"') {
        l.incCur()
    }
    l.tokens = append(l.tokens, tokens.Token{
        TokenType: tokens.STRING,
        Lexem: l.slice(),
    })
    l.incSync()
}

func (l *Lexer) slice() string {
    return string(l.chars[l.start:l.end])
}
