package problems
import (
	"dsa/stack"
)

func Reverse(s string) string {
	stack := stack.Stack{};

	for i:=0; i<len(s); i++{
		stack.Push(int(s[i]))
	}
	reversed := ""
	for i:=0; i<len(s); i++{
		reversed+=string(byte(stack.Pop()))
	}
	return reversed
}