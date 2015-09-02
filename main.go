package main

import (
	"fmt"
	"worldTree"
)

func main() {
	var Nordrassil worldTree.TreeNode
	Nordrassil.WorldTreeInit()
	value := Nordrassil.GetNodeValue()
	fmt.Println(value)
}
