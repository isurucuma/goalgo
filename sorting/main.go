package main

import (
	"fmt"
	"mysort/bubble"
	"mysort/quick"
)

func main() {
	items := []int{5, 3, 4, 1, 2, 2, 2, 1}
	fmt.Println("before sorting: ", items)
	bubble.Sort(items)
	fmt.Println("after bubble sorting: ", items)
	items = []int{5, 3, 4, 1, 2, 2, 2, 1}
	fmt.Println("before sorting: ", items)
	quick.Sort(items)
	fmt.Println("after quick sorting: ", items)
}
