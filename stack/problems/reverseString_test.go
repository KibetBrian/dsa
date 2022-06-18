package problems

import (
	"testing"

	"github.com/stretchr/testify/require"
)


func TestReverseString (t *testing.T) {
	name := "kibet"
	rev := "tebik"
	reversed := Reverse(name)
	require.Equal(t, reversed, rev)
}