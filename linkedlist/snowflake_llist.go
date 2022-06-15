package main

import "fmt"

type snowflake struct {
	code int
	next *snowflake
}

func main() {
	var tdata = []int{1, 2, 3, 4, 5, 6, 7, 8}

	var snowflakes *snowflake

	for _, v := range tdata {
		tnode := &snowflake{v, snowflakes}
		fmt.Println(tnode)
		snowflakes = tnode

	}
	fmt.Println()
	currentNode := snowflakes
	for currentNode != nil {
		fmt.Println(currentNode)
		currentNode = currentNode.next
	}

}
