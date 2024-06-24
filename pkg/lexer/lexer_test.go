package lexer_test

import (
	lexer "lexer"
	"lexer/tokens"
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
    var tests = []struct {
        source string
        want []tokens.Token
    }{
        {
            "Hey",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Hey"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "hey",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "hey"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "Hey1",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Hey1"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "hey1",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "hey1"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "    hey1 \t \n",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "hey1"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "Name: hey",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Name"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.IDENT, Lexem: "hey"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            "Method: \"GET\"",
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Method"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "GET"},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
        {
            `Name: "Test 1"
            GET: "http://localhost:3000/api/v1/user"
            BODY: {
                "headers": {
                    "Content-Type": "application/json",
                    "Bearer": "xyz",
                    "something": true
                }
            }`,
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Name"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "Test 1"},
                {TokenType: tokens.IDENT, Lexem: "GET"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "http://localhost:3000/api/v1/user"},
                {TokenType: tokens.IDENT, Lexem: "BODY"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.L_BRACKET, Lexem: ""},
                {TokenType: tokens.STRING, Lexem: "headers"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.L_BRACKET, Lexem: ""},
                {TokenType: tokens.STRING, Lexem: "Content-Type"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "application/json"},
                {TokenType: tokens.COMMA, Lexem: ","},
                {TokenType: tokens.STRING, Lexem: "Bearer"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "xyz"},
                {TokenType: tokens.COMMA, Lexem: ","},
                {TokenType: tokens.STRING, Lexem: "something"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.BOOLEAN, Lexem: "true"},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
                {TokenType: tokens.EOF, Lexem: ""},
            },
        },
    }

    for _, tt := range tests {
        t.Run("Lex test", func(t *testing.T) {
            got := lexer.Lex(tt.source)
            if !reflect.DeepEqual(got, tt.want) {
                t.Fatalf("Lex failed")
            }
        })
    }
}
