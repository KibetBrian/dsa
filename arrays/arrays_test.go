package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertBeginnning(t *testing.T) {
	var numbers = [10]int{1, 2, 3, 4, 5,6,7,8,9,10};
	numToBeInserted := 232;

	numbers = InsertBeginnning(numbers, numToBeInserted)
	require.Equal(t, numbers[0],numToBeInserted)
}

func TestBegginingDeletion (t *testing.T) {
	var numbers = [10]int{1, 2, 3, 4, 5,6,7,8,9,10};
	newNums := DeleteAtTheBeginning(numbers);
	
	if newNums[0]==numbers[0]{
		t.Error("Failed deletion")
	}

}

func TestDeleteAtAnyWhere(t *testing.T){
	var numbers = [10]int{1, 2, 3, 4, 5,6,7,8,9,10};
	numToBeDeleted := 5;
	newNumbs := DeleteAtAnyWhere(numbers, numToBeDeleted);

	for i:=0; i<len(newNumbs); i++{
		if newNumbs[i]==numToBeDeleted{
			t.Errorf("\nDeletion didn't work.\n OriginalNumbers: %v\n NewNumbers: %v",numbers,newNumbs)
		}
	}


}
