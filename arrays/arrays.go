package arrays

func main (){
	
}

func InsertBeginnning(arr [10]int, n int) [10]int{
	for i:=len(arr)-1; i>=1; i--{
		arr[i]=arr[i-1]
	}
	arr[0]=n
	return arr
}

func DeleteAtTheBeginning(arr [10]int) [10]int{
	for i:=0; i<len(arr)-1; i++{
		arr[i]=arr[i+1]
	}
	arr[len(arr)-1]=0;
	return arr
};

func DeleteAtAnyWhere(arr[10]int, n int) [10]int{
	counter := 0;
	for i:=0; i<len(arr); i++{
		if (arr[i] != n){
			arr[counter]=arr[i]
			counter++
		}
	}
	return arr;
}