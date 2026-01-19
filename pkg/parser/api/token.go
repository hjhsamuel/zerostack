package api

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenIdent
	TokenString
	TokenBool

	TokenAssign   // =
	TokenAt       // @
	TokenColon    // :
	TokenLBrace   // {
	TokenRBrace   // }
	TokenLParen   // (
	TokenRParen   // )
	TokenBacktick // `
	TokenComment  // //
	TokenNewLine
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

var customGroupKeywords = []string{
	"api",
	"router",
}
