package main

import (
	"fmt"
	"github.com/ukrainskiys/go-collections/collections"
)

func print(i int) {
	fmt.Print(i, " ")
}

func main() {
	set := collections.NewSet[int]()

	set.AddAll(collections.NewSetOfSlice[int]([]int{1, 2, 3, 4}))

	set.AddAll(collections.NewSetOfSlice[int]([]int{5, 10, 3, 4}))

	set2 := collections.NewSetOf[int](set)
	set2.ForEach(print)

	//for i := range set.Iterator() {
	//	fmt.Println(i)
	//}
}
