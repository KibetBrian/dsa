package main

import (
	"log"
	"dsa/stack"

)

func main() {
	stack := stack.Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	log.Println(stack.Pop())
	log.Println(stack.Peek())
	log.Println(stack.Peek())
	log.Println(stack.Peek())
	
}
