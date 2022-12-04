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
	set.ForEach(print)
	fmt.Println(set.Size())

	set.AddAll(collections.NewSet[int](1, 2, 3, 4, 4, 4, 4))
	set.ForEach(print)
	fmt.Println(set.Size())

	set.AddAll(collections.NewSet[int](5, 10, 3, 4))

	set2 := collections.NewSetOf[int](set)
	set2.ForEach(print)
	fmt.Println()

	for i := range set2.Iterator() {
		fmt.Println(i, "qwe")
	}

	//for i := range set.Iterator() {
	//	fmt.Println(i)
	//}
}
