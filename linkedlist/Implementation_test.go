package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertion(t *testing.T) {
	list := LinkedList{}
	var i int
	for i = 0; i < 10; i++ {
		list.Insert(i)
	}
	require.Equal(t, i, list.length)
}

func TestDeletion(t *testing.T){
	list := LinkedList{}
	
	for i:=0; i<10; i++{
		list.Insert(i)
	}
	for i:=0; i<11; i++{
		list.Delete(i)
	}
	require.Equal(t, list.length, 0)
}
