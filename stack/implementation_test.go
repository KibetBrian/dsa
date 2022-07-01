package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	stack := Stack{}
	total := 10;

	//Push values into a stack
	for i := 0; i < total; i++ {
		stack.Push(i)
	}

	size := len(stack.values)

	require.Equal(t, size, total)
}

func TestPop (t *testing.T){
	stack := Stack{}
	total := 5
	//Inset values upto total - 1
	for i := 0; i < total; i++ {
		stack.Push(i)
	}
	
	poped := stack.Pop();
	size := len(stack.values)
	
	require.Equal(t, poped, total-1)
	require.Equal(t, size, total-1)
}

func TestPeek (t *testing.T) {
	stack := Stack{}
	total := 5

	//Inset values upto total - 1
	for i := 0; i < total; i++ {
		stack.Push(i)
	}
	peeked, _ := stack.Peek()
	size := len(stack.values)
	require.Equal(t, peeked, total-1);
	require.Equal(t, size, t)
}