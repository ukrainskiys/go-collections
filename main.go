package main

import (
	"fmt"
	"github.com/ukrainskiys/go-collections/collections"
)

func print(i int) {
	fmt.Print(i, " ")
}

func testSet() {

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

	for i := range set.Iterator() {
		fmt.Println(i)
	}

	set.Equal(set2)

	fmt.Println(set2)
}

func testArray() {
	type TEST struct {
		tests string
		testi int
	}

	ar := collections.NewList[TEST](TEST{"1", 1}, TEST{"2", 2})
	fmt.Println(ar.Contains(TEST{"1", 1}))

	arr := collections.NewList[int]()
	arr.ForEach(print)
	fmt.Println(arr.Size())

	arr.AddAll(collections.NewList[int](1, 2, 3, 4, 4, 4, 4))
	arr.ForEach(print)
	fmt.Println(arr.Size())

	arr.AddAll(collections.NewSet[int](5, 10, 3, 4))

	arr2 := collections.NewListOf[int](arr)
	arr2.ForEach(print)
	fmt.Println()

	for i := range arr2.Iterator() {
		fmt.Println(i)
	}

	fmt.Println(arr2)
}

func testQueue() {
	q := collections.NewQueue[int](5, 4, 3, 2, 1)
	fmt.Println(q.Peek())
	fmt.Println(q.Pool())
	fmt.Println(q)

	q.Clear()
	fmt.Println(q)

	q.Offer(1)
	q.Offer(1)
	q.Offer(100)
	q.Offer(1111)
	q.Pool()
	fmt.Println(q)

}

func main() {
	testQueue()
	//testSet()
	//testArray()
}
