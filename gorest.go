package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jlucero805/gorest/lexer"
	"github.com/jlucero805/gorest/parser"
)


func main() {
    contents, _ := os.ReadFile("test.gorest")
    tokens := lexer.Lex(string(contents))
    values := parser.Parse(tokens)
    if url, ok := values.Meta["GET"]; ok {
        req, err := http.NewRequest(http.MethodGet, url, nil)
        if err != nil {
            log.Fatal(err)
        }

        res, err := http.DefaultClient.Do(req)
        if err != nil {
            log.Fatal(err)
        }

        body, err := ioutil.ReadAll(res.Body)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Print(string(body))
    }
}
