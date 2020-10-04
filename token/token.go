package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // トークンや文字が未知であることを表す
	EOF     = "EOF"     // end of file

	// 識別子+リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 12345

	// 演算子
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	BANG = "!"
	SLASH = "/"
	ASTERISK = "*"

	LT = "<"
	GT = ">"

// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EQ = "=="
	NOT_EQ = "!="

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
