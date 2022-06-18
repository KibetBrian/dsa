package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	list := LinkedList{}
	var i int
	for i = 0; i < 10; i++ {
		list.Insert(i)
	}
	require.Equal(t, i, list.length)
}

func TestDelete(t *testing.T) {
	list := LinkedList{}

	for i := 0; i < 10; i++ {
		list.Insert(i)
	}

	for i := 0; i < 11; i++ {
		list.Delete(i)
	}

	for i := 0; i < 10; i++ {
		list.Insert(i)
	}

	for i := 11; i >=0; i-- {
		list.Delete(i)
	}

	require.Equal(t, list.length, 0)
}

func TestSize(t *testing.T) {
	list := LinkedList{}
	var i int
	for i = 0; i < 5; i++ {
		list.Insert(i)
	}

	size := list.Size()
	require.Equal(t, i, size)
}

func TestExists(t *testing.T) {
	list := LinkedList{}
	var numbersToExist = []int{22, 14, 42}
	var numberToNotExist = []int{2,5,7}

	for _, v := range numbersToExist {
		list.Insert(v)
	}

	for i:=0; i<len(numberToNotExist); i++{
		if list.Exists(numberToNotExist[i]){
			t.Errorf("Found a non inserted number. Number: %v,", numberToNotExist[i])
		}
	}

	for _, v := range numbersToExist {
		if !list.Exists(v) {
			t.Errorf("Number missing: %v", v)
		}
	}
}
