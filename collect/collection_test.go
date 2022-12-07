package collect

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	ten      = 10
	hundred  = 100
	thousand = 1_000
	million  = 1_000_000
)

func TestCollectionWithSlice_Add(t *testing.T) {
	err := checkCollectionOf(createCollectionOf(million), []int{million})
	if err != nil {
		t.Error(err)
	}
}

func TestCollectionWithSlice_AddAll(t *testing.T) {
	type test struct {
		adding *collectionWithSlice[int]
		check  []int
	}

	coll := createCollectionOf(thousand)

	tests := []test{
		{adding: createCollectionOf(ten), check: []int{thousand, ten}},
		{adding: createCollectionOf(hundred), check: []int{thousand, ten, hundred}},
		{adding: createCollectionOf(thousand), check: []int{thousand, ten, hundred, thousand}},
		{adding: createCollectionOf(million), check: []int{thousand, ten, hundred, thousand, million}},
		{adding: createCollectionOf(million), check: []int{thousand, ten, hundred, thousand, million, million}},
	}

	for idx, test := range tests {
		coll.AddAll(test.adding)
		err := checkCollectionOf(coll, test.check)
		if err != nil {
			t.Error(fmt.Sprintf("test index=%d", idx), err)
		}
	}
}

func TestCollectionWithSlice_AddAllSlice(t *testing.T) {
	type test struct {
		adding []int
		check  []int
	}

	coll := createCollectionOf(thousand)

	tests := []test{
		{adding: *createCollectionOf(ten).data, check: []int{thousand, ten}},
		{adding: *createCollectionOf(hundred).data, check: []int{thousand, ten, hundred}},
		{adding: *createCollectionOf(thousand).data, check: []int{thousand, ten, hundred, thousand}},
		{adding: *createCollectionOf(million).data, check: []int{thousand, ten, hundred, thousand, million}},
		{adding: *createCollectionOf(million).data, check: []int{thousand, ten, hundred, thousand, million, million}},
	}

	for idx, test := range tests {
		coll.AddAllSlice(test.adding)
		err := checkCollectionOf(coll, test.check)
		if err != nil {
			t.Error(fmt.Sprintf("test index=%d", idx), err)
		}
	}
}

func TestCollectionWithSlice_Size(t *testing.T) {
	type test struct {
		coll *collectionWithSlice[int]
		size int
	}

	coll := createCollectionOf(0)

	tests := []test{
		{coll: createCollectionOf(ten), size: ten},
		{coll: createCollectionOf(hundred), size: ten + hundred},
		{coll: createCollectionOf(thousand), size: ten + hundred + thousand},
		{coll: createCollectionOf(million), size: ten + hundred + thousand + million},
		{coll: createCollectionOf(million), size: ten + hundred + thousand + million + million},
	}

	for _, test := range tests {
		coll.AddAll(test.coll)
		if coll.Size() != test.size {
			t.Errorf("expected error, containing=%d, got=%d", coll.Size(), test.size)
		}
	}
}

func TestCollectionWithSlice_Clear(t *testing.T) {
	coll := createCollectionOf(ten)
	coll.Clear()
	if len(*coll.data) != 0 {
		t.Errorf("expected error, containing=%d, got=%d", 0, len(*coll.data))
	}

	coll.AddAll(createCollectionOf(million))
	coll.Clear()
	if len(*coll.data) != 0 {
		t.Errorf("expected error, containing=%d, got=%d", 0, len(*coll.data))
	}
}

func TestCollectionWithSlice_Contains(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	coll := createCollectionOf(million)
	for i := 0; i < thousand; i++ {
		ran := rand.Intn(million)
		if !coll.Contains(ran) {
			t.Errorf("expected error, not containing=%d", ran)
		}
	}

	if coll.Contains(million + 1) {
		t.Errorf("expected error, containing=%d", million+1)
	}
}

func TestCollectionWithSlice_ContainsAll(t *testing.T) {
	tests := []int{ten, hundred, thousand, hundred}
	coll := createCollectionOf(million)
	for _, count := range tests {
		randomSlice := generateRandomSlice(count)
		if !coll.ContainsAll(NewList(randomSlice...)) {
			t.Errorf("expected error, not containing")
		}
		if coll.ContainsAll(NewList(append(randomSlice, million+1)...)) {
			t.Errorf("expected error, not containing")
		}
	}
}

func TestCollectionWithSlice_ContainsAllSlice(t *testing.T) {
	tests := []int{ten, hundred, thousand, hundred}
	coll := createCollectionOf(million)
	for _, count := range tests {
		randomSlice := generateRandomSlice(count)
		if !coll.ContainsAllSlice(randomSlice) {
			t.Errorf("expected error, not containing")
		}
		if coll.ContainsAllSlice(append(randomSlice, million+1)) {
			t.Errorf("expected error, not containing")
		}
	}
}

func TestCollectionWithSlice_Remove(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	coll := createCollectionOf(million)
	var deleted []int
	for i := 0; i < thousand; i++ {
		deletedElement := rand.Intn(million)
		coll.Remove(deletedElement)
		deleted = append(deleted, deletedElement)
	}

	for _, d := range deleted {
		if coll.Contains(d) {
			t.Errorf("expected error, not deleted %d", d)
		}
	}
}

func TestCollectionWithSlice_RemoveAll(t *testing.T) {
	coll := createCollectionOf(million)
	deleted := generateRandomSlice(thousand)
	res := coll.RemoveAll(NewList(deleted...))
	if !res {
		t.Errorf("expected error, list wasn't modified")
	}

	for _, d := range deleted {
		if coll.Contains(d) {
			t.Errorf("expected error, not deleted %d", d)
		}
	}
}

func TestCollectionWithSlice_RemoveAllSlice(t *testing.T) {
	coll := createCollectionOf(million)
	deleted := generateRandomSlice(thousand)
	res := coll.RemoveAllSlice(deleted)
	if !res {
		t.Errorf("expected error, list wasn't modified")
	}

	for _, d := range deleted {
		if coll.Contains(d) {
			t.Errorf("expected error, not deleted %d", d)
		}
	}
}

func TestCollectionWithSlice_RemoveIf(t *testing.T) {
	coll := createCollectionOf(thousand)
	res := coll.RemoveIf(func(int2 int) bool {
		return int2%2 == 0
	})
	if !res {
		t.Errorf("expected error, list wasn't modified")
	}

	if coll.Size() != thousand/2 {
		t.Errorf("expected error, incorrect size collection after RemoveIf %d", thousand/2)
	}

	for val := range coll.Iterator() {
		if val%2 == 0 {
			t.Errorf("expected error, not deleted %d", val)
		}
	}
}

func TestCollectionWithSlice_IsEmpty(t *testing.T) {
	coll := createCollectionOf(0)
	if !coll.IsEmpty() {
		t.Errorf("expected error, list not empty")
	}
	coll.Add(1)
	if coll.IsEmpty() {
		t.Errorf("expected error, list empty")
	}
	coll.Clear()
	if !coll.IsEmpty() {
		t.Errorf("expected error, list not empty")
	}
}

func generateRandomSlice(count int) []int {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < count/10; i++ {
		arr = append(arr, rand.Intn(count))
	}
	return arr
}

func createCollectionOf(size int) *collectionWithSlice[int] {
	var arr []int
	coll := &collectionWithSlice[int]{&arr}
	for i := 0; i < size; i++ {
		coll.Add(i)
	}
	return coll
}

func checkCollectionOf(coll *collectionWithSlice[int], counts []int) error {
	idx := 0
	for _, count := range counts {
		for i := 0; i < count; i++ {
			val := (*coll.data)[idx]
			if val != i {
				return errors.New(fmt.Sprintf("expected error case=%d, containing=%d, got=%d", count, i, val))
			}
			idx++
		}
	}
	return nil
}
