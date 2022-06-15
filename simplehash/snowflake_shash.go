package main

import "fmt"

var MAX int = 100

type snowflake struct {
	snowflakes []int
	next       *snowflake
}

func hash(elems []int) int {
	var code int
	for _, v := range elems {
		code += v
	}
	return code % MAX
}

func main() {
	var tdata = [][]int{{1, 2, 3, 4, 5, 6},
		{2, 3, 4, 5, 6, 7},
		{3, 4, 5, 6, 7, 2}} //will have same "hash" as second element in array causing collision

	var snowflakes [100]*snowflake //data structure to hold hashmap; collisions will be stored in linked list node

	for _, v := range tdata {
		snowflake_code := hash(v)
		tnode := &snowflake{v, snowflakes[snowflake_code]}
		fmt.Println(tnode)
		snowflakes[snowflake_code] = tnode
	}

	fmt.Println()
	for _, v := range snowflakes {
		cnode := v
		for cnode != nil {
			fmt.Println(cnode)
			cnode = cnode.next
		}
	}

}
