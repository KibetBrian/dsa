package queue

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEnqueue(t *testing.T) {
	q := Queue{};
	q.Enqueue("Brian")
	q.Enqueue("Kemboi")
	q.Enqueue("Kibet")

	require.Equal(t, len(q.vals),3)
	require.Equal(t, q.vals[0], "Brian")
	require.Equal(t, q.vals[1], "Kemboi")
	require.Equal(t, q.vals[2], "Kibet")
}

func TestDequeue(t *testing.T) {
	q := Queue{};
	
	q.Enqueue("Brian")
	q.Enqueue("Kemboi")
	q.Enqueue("Kibet")

	require.Equal(t, q.vals[0], "Brian")
	q.Dequeue();
	require.Equal(t, q.vals[0], "Kemboi")
	q.Dequeue()
	require.Equal(t, q.vals[0], "Kibet")

}


func RandomNameGenerator() string{
	var alphabet = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	min := 0;
	max := len(alphabet);
	rand.Seed(time.Now().UnixNano());
	
	res := ""
	
	for i:=0; i<6; i++{
		index := rand.Intn((max-min+1)+min)
		res+=alphabet[index]
	}
	return res
}