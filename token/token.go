package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // トークンや文字が未知であることを表す
	EOF = "EOF" // end of file

	// 識別子+リテラル
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT = "INT" // 12345

	// 演算子
	ASSIGN = "="
	PLUS = "+"

	// デリミタ
	COMMA = ","
	SEMICORON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET = "LET"
)
