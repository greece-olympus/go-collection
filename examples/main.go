package main

import (
	"fmt"
	collection "github.com/greece-olympus/go-collection"
)

func main() {
	col := collection.NewCollection([]int{1, 2, 3, 4})
	fmt.Println(col.Count())
}
