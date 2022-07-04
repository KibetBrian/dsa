package main

import (
	"dsa/graphs"
)

func main() {
	names := []string{"Brian","Kemboi","Kibet"}
	myGraph := graphs.Graph{}

	for i := 0; i <len(names); i++ {
		myGraph.Insert(names[i])
	}
	myGraph.PrintValues()

}

