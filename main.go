package main

import (
	"dsa/linkedlist"
	"fmt"
	"log"
)

func main(){
	fmt.Println("Hello DSA!")
	l := linkedlist.LinkedList{}
	
	for i:=0; i<3; i++{
		l.Insert(i)
	}
	
	for i:=0; i<4; i++{
		l.Delete(i)
		log.Printf("Executed Deleting: %v", i)
		log.Printf("Length: %v", l.Size())
	}
	l.PrintValues()
}