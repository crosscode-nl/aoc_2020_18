package token

type OperatorType rune

const (
	Invalid  OperatorType = '~'
	Plus = '+'
	Multiply = '*'
	OpenParenthesis = '('
	CloseParenthesis = ')'
)

type Type int

const (
	Unknown Type = iota
	Number
	Operator
)

type Token struct {
	tokenType Type
	operator  OperatorType
	number    int
}

func (t Token)Operator() OperatorType {
	return t.operator
}

func (t Token)Number() int {
	return t.number
}

func (t Token)Type() Type {
	return t.tokenType
}

func NewOperatorToken(operator OperatorType) Token {
	return Token{
		tokenType: Operator,
		operator: operator,
		number:   0,
	}
}

func NewUnknownToken() Token {
	return Token {}
}

func NewNumberToken(number int) Token {
	return Token{
		tokenType: Number,
		operator: Invalid,
		number:   number,
	}
}

