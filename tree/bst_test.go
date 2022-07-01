package tree
import (
	"testing"
)

func TestInsert(t *testing.T) {
	tree := Node{};

	for i := 0; i<10; i++ {
		tree.Insert(i)
	}
	
}