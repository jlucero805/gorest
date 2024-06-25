package parser_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jlucero805/gorest/lexer/tokens"
	"github.com/jlucero805/gorest/parser"
)

func TestParser(t *testing.T) {
    var tests = []struct {
        tokens []tokens.Token
        want parser.ParseResult 
    }{
        {
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "Name"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "Test 1"},
                {TokenType: tokens.IDENT, Lexem: "GET"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "http://localhost:3000/api/v1/user"},
            },
            parser.ParseResult{
                Meta: map[string]string{
                    "Name": "Test 1",
                    "GET": "http://localhost:3000/api/v1/user",
                },
            },
        },
        {
            []tokens.Token{
                {TokenType: tokens.IDENT, Lexem: "BODY"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.L_BRACKET, Lexem: ""},
                {TokenType: tokens.STRING, Lexem: "name"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "Jason"},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
            },
            parser.ParseResult{
                Meta: map[string]string{},
                Body: "{\"name\":\"Jason\"}",
            },
        },
        {
            []tokens.Token{
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
                {TokenType: tokens.STRING, Lexem: "Accept"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "text/html"},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
            },
            parser.ParseResult{
                Meta: map[string]string{},
                Body: "{\"headers\":{\"Content-Type\":\"application/json\",\"Accept\":\"text/html\"}}",
            },
        },
        {
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
                {TokenType: tokens.STRING, Lexem: "Accept"},
                {TokenType: tokens.COLON, Lexem: ":"},
                {TokenType: tokens.STRING, Lexem: "text/html"},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
                {TokenType: tokens.R_BRACKET, Lexem: ""},
            },
            parser.ParseResult{
                Meta: map[string]string{
                    "Name": "Test 1",
                    "GET": "http://localhost:3000/api/v1/user",
                },
                Body: "{\"headers\":{\"Content-Type\":\"application/json\",\"Accept\":\"text/html\"}}",
            },
        },
    }

    for _, tt := range tests {
        t.Run("Parse Test", func(t *testing.T) {
            got := parser.Parse(tt.tokens)
            fmt.Print(got.Body)
            if !reflect.DeepEqual(got, tt.want) {
                t.Fatalf("Parse failed")
            }
        })
    }
}
