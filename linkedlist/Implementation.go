package linkedlist

import (
	"log"
)

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Insert(n int) {
	newNode := &Node{
		Val: n,
	}
	list := l.head
	l.head = newNode
	l.head.Next = list
	l.length++
}

func (l *LinkedList) Delete(v int) {
	if l.length == 0 {
		return
	}
	if l.head.Val == v {
		l.head = l.head.Next
	}
	cur := l.head
	var prev *Node
	for cur != nil {
		if cur.Val == v && cur.Next != nil {
			prev.Next = cur.Next
			l.length--
		} else {
			cur = nil
			l.length--
		}
	}
	cur = prev
}

func (l *LinkedList) PrintValues() {
	cur := l.head
	for cur != nil {
		log.Println(cur.Val)
		cur = cur.Next
	}
}

func (l *LinkedList) Size() int {
	return l.length
}

func (l *LinkedList) Exists(n int) bool {
	cur := l.head
	for cur != nil {
		if cur.Val == n {
			return true
		}
		cur = cur.Next
	}
	return false
}
