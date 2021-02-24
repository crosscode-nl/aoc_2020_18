package token

import "fmt"

type Stack struct {
	tokens []Token
}

func (ts *Stack) Len() int {
	return len(ts.tokens)
}

func  (ts *Stack) Tokens() []Token {
	return ts.tokens
}

func (ts *Stack) Push(token Token) {
	ts.tokens = append(ts.tokens, token)
}

func (ts *Stack) Print() {
	for _,t:=range ts.tokens {
		switch t.Type() {
			case Number : {
				fmt.Printf("%d ",t.number)
			}
			case Operator: {
				fmt.Printf("%s ",string(t.Operator()))
			}
		}
	}
	fmt.Print("\n")
}

func (ts *Stack) Flipped() Stack {
	newStack := Stack{}
	for t,ok:=ts.Pop(); ok; t,ok=ts.Pop() {
		newStack.Push(t)
	}
	return newStack
}

func (ts *Stack) last() int {
	return len(ts.tokens)-1
}

func (ts *Stack) Peek() (Token, bool) {
	if len(ts.tokens)==0 { return NewUnknownToken(), false }
	return ts.tokens[ts.last()], true
}

func (ts *Stack) Pop() (Token, bool) {
	token, ok := ts.Peek()
	if ok {
		ts.tokens = ts.tokens[:ts.last()]
	}
	return token, ok
}
