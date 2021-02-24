package main

import (
	"aoc_2020_18/token"
	"bufio"
	"fmt"
	"os"
)

type PrecedenceFunc func(stackOperator token.OperatorType, newOperator token.OperatorType) int

func convertToRPN(input []token.Token, takesPrecedence PrecedenceFunc) token.Stack {
	outputStack := &token.Stack{}
	operatorStack := &token.Stack{}
	for _, t := range input {
		switch t.Type() {
		case token.Operator:
			{
				if operatorStack.Len() == 0 {
					operatorStack.Push(t)
				} else {
					switch t.Operator() {
					case token.Plus:
						{
							handleStacks(operatorStack, takesPrecedence, token.Plus, outputStack, t)
						}
					case token.Multiply:
						{
							handleStacks(operatorStack, takesPrecedence, token.Multiply, outputStack, t)
						}
					case token.OpenParenthesis:
						{
							operatorStack.Push(t)
						}
					case token.CloseParenthesis:
						{
							for topt, ok := operatorStack.Pop(); ok && topt.Operator() != token.OpenParenthesis; topt, ok = operatorStack.Pop() {
								outputStack.Push(topt)
							}
						}
					case token.Invalid:
						{
							fmt.Println(t)
							panic("Bug, invalid operator encountered")
						}
					}
				}
			}
		case token.Number:
			{
				outputStack.Push(t)
			}
		}
	}
	for t, ok := operatorStack.Pop(); ok; t, ok = operatorStack.Pop() {
		outputStack.Push(t)
	}
	return outputStack.Flipped()
}

func handleStacks(operatorStack *token.Stack, takesPrecedence PrecedenceFunc, tokenOperator token.OperatorType, outputStack *token.Stack, t token.Token) {
	for operatorStack.Len() > 0 {
		topOperatorToken, _ := operatorStack.Pop()
		if takesPrecedence(topOperatorToken.Operator(), tokenOperator) > 0 {
			operatorStack.Push(topOperatorToken)
			break
		} else {
			outputStack.Push(topOperatorToken)
			continue
		}
	}
	operatorStack.Push(t)
}

func multiplyTokens(a token.Token, b token.Token) token.Token {
	return token.NewNumberToken(a.Number() * b.Number())
}

func sumTokens(a token.Token, b token.Token) token.Token {
	return token.NewNumberToken(a.Number() + b.Number())
}

func evalRPN(rpn token.Stack) int {
	outputStack := token.Stack{}
	for t, ok := rpn.Pop(); ok; t, ok = rpn.Pop() {
		switch t.Type() {
		case token.Number:
			{
				outputStack.Push(t)
			}
		case token.Operator:
			{
				a, _ := outputStack.Pop()
				b, _ := outputStack.Pop()
				switch t.Operator() {
				case token.Multiply:
					{
						outputStack.Push(multiplyTokens(a, b))
					}
				case token.Plus:
					{
						outputStack.Push(sumTokens(a, b))
					}
				default:
					panic("Bug: invalid operator encountered")
				}
			}
		}
	}
	resultToken, _ := outputStack.Pop()
	return resultToken.Number()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculate(lines []string, precedenceFunc PrecedenceFunc) int {
	total := 0
	for _, l := range lines {
		res := evalRPN(convertToRPN(token.Tokenize(l), precedenceFunc))
		total += res
	}
	return total
}

func fileToLines(fileName string) (lines []string) {
	f, err := os.Open(fileName)
	check(err)
	defer func() {
		check(f.Close())
	}()
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return
}

var OldPrecedenceRules = func(stackOperator token.OperatorType, newOperator token.OperatorType) int {
	if newOperator==stackOperator { return 0 }
	if newOperator==token.OpenParenthesis && stackOperator==token.Multiply { return 1 }
	if newOperator==token.OpenParenthesis && stackOperator==token.Plus { return 1 }
	if newOperator==token.Multiply && stackOperator==token.OpenParenthesis { return 1 }
	if newOperator==token.Multiply && stackOperator==token.Plus { return 0 }
	if newOperator==token.Plus && stackOperator==token.Multiply { return 0 }
	if newOperator==token.Plus && stackOperator==token.OpenParenthesis { return 1 }
	return 0
}

var NewPrecedenceRules = func(stackOperator token.OperatorType, newOperator token.OperatorType) int {
	if newOperator==stackOperator { return 0 }
	if newOperator==token.OpenParenthesis && stackOperator==token.Multiply { return -1 }
	if newOperator==token.OpenParenthesis && stackOperator==token.Plus { return -1 }
	if newOperator==token.Multiply && stackOperator==token.OpenParenthesis { return 1 }
	if newOperator==token.Multiply && stackOperator==token.Plus { return -1 }
	if newOperator==token.Plus && stackOperator==token.Multiply { return 1 }
	if newOperator==token.Plus && stackOperator==token.OpenParenthesis { return 1 }
	return 0
}

func main() {
	lines := fileToLines("calc1.txt")
	fmt.Println("Solution of part 1:")
	fmt.Printf("Total: %d\n", calculate(lines, OldPrecedenceRules))
	fmt.Println("\nSolution of part 2:")
	fmt.Printf("Total: %d\n", calculate(lines, NewPrecedenceRules))
}
