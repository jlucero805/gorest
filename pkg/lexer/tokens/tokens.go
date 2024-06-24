package tokens

const (
    STRING = iota
    COLON
    BOOLEAN
    NUMBER
    L_BRACKET
    R_BRACKET
    IDENT
    COMMA
    EOF
)

type Token struct {
    TokenType int
    Lexem string
}
