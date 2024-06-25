package grammar_test

import (
	"testing"
	"github.com/jlucero805/gorest/lexer/grammar"
)

func TestDigit(t *testing.T) {
    var tests = []struct {
        char rune
        want bool
    }{
        {'0', true},
        {'1', true},
        {'9', true},
        {'a', false},
        {'z', false},
    }

    for _, tt := range tests {
        t.Run("Digit Test", func(t *testing.T) {
            got := grammar.Digit(tt.char)
            if got != tt.want {
                t.Fatalf("Digit failed")
            }
        })
    }
}

func TestAlpha(t *testing.T) {
    var tests = []struct {
        char rune
        want bool
    }{
        {'a', true},
        {'A', true},
        {'z', true},
        {'Z', true},
        {'0', false},
        {'9', false},
    }

    for _, tt := range tests {
        t.Run("Alpha Test", func(t *testing.T) {
            got := grammar.Alpha(tt.char)
            if got != tt.want {
                t.Fatalf("Alpha failed")
            }
        })
    }
}
