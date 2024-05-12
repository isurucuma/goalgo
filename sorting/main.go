package main

import (
	"fmt"
	"mysort/bubble"
)

func main() {
	items := []int{5, 3, 4, 1, 2, 2, 2, 1}
	bubble.Sort(items)
	fmt.Println(items)
}
