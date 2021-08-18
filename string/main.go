package main

import (
	"fmt"
	"string/stringManipulation"
)

func main() {
	
	str := "Matematika"

	fmt.Println(stringManipulation.DeleteIndex(str, 3))

	fmt.Println(stringManipulation.CountChar(str, 'a'))

	fmt.Println(stringManipulation.Reverse(str))
}