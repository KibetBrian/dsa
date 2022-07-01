package tree

import "log"

type Node struct {
	val        int
	leftChild  *Node
	rightChild *Node
}


func (n *Node) Insert(val int) {
	newNode := &Node{val: val}

	if n.val > val {
		if n.leftChild == nil {
			n.leftChild = newNode
			return
		}
		n.leftChild.Insert(val)
	} else {
		if n.rightChild == nil {
			n.rightChild = newNode
			return
		}
		n.rightChild.Insert(val)
	}
}

func (t *Node) Search(n int) bool {
	if t == nil {
		return false
	}

	if t.val > n {
		return t.leftChild.Search(n)
	}
	if t.val < n {
		return t.rightChild.Search(n)
	}

	return true

}

func (t *Node) InOder(){
	if t == nil {
		return
	}
	t.leftChild.InOder()
	log.Println("Value: ", t.val)
	t.rightChild.InOder()
}

func (t *Node) PreOrder() {
	if t == nil {
		return
	}
	log.Println(t.val)
	t.leftChild.PreOrder()
	t.rightChild.PreOrder()
	return
}

func (t *Node) PostOrder() {
	if t==nil{
		return
	}
	t.leftChild.PostOrder()
	t.rightChild.PostOrder()
	log.Println(t.val)
}
