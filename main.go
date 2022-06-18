package main

import (
	"log"
	"dsa/stack/problems"

)

func main() {
	name := "brian"	
	reversed := problems.Reverse(name)
	log.Println(name, reversed)
}
