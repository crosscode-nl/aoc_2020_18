package token

import "strconv"

func appendNumberToToken(n int, t Token) Token {
	return NewNumberToken(t.Number()*10+n)
}

func Tokenize(input string) []Token {
	var tokens []Token
	for _,c := range input {
		switch c {
		case rune(Plus) : {
			tokens=append(tokens,NewOperatorToken(Plus))
		}
		case rune(Multiply) : {
			tokens=append(tokens,NewOperatorToken(Multiply))
		}
		case rune(OpenParenthesis) : {
			tokens=append(tokens,NewOperatorToken(OpenParenthesis))
		}
		case rune(CloseParenthesis) : {
			tokens=append(tokens,NewOperatorToken(CloseParenthesis))
		}
		default:
			if n,err := strconv.Atoi(string(c)); err == nil {
				if len(tokens)>0 {
					last := len(tokens)-1
					if tokens[last].Type()==Number {
						tokens[last] = appendNumberToToken(n,tokens[last])
						break
					}
				}
				tokens=append(tokens,NewNumberToken(n))
			}
		}
	}
	return tokens
}