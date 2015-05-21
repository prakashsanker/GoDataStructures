package main;

import "fmt" 

func main() {
	testSlice := []int{5,4,3,2,1}
	fmt.Println(heapSort(testSlice))
}
type IMaxHeap interface {

	/*Insert(i interface{}) bool

	FindMax() interface{}

	ExtractMax() interface{}

	DeleteMax() */


	//Merge(toMerge MaxHeap) MaxHeap

	findMax() int 

	size() int
}

type MaxHeap struct {
	slice []int
	heapSize int
}


func (h MaxHeap) size() int {
	return h.heapSize
}

func (h MaxHeap) findMax() int{
	return h.slice[0]
}

func heapSort(slice []int) []int {
	h := BuildMaxHeap(slice)
	for i := len(slice); i >=1 ; i-- {
		first := slice[0]
		last := slice[len(slice) - 1]
		slice[0] = last
		slice[len(slice) - 1] = first
		h.heapSize--
		h.MaxHeapify(1)
	}
	return h.slice
}

func BuildMaxHeap(slice []int) MaxHeap{
	h := MaxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice)/2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	left := 2*i
	right := 2*i + 1
	largest := 0
	slice := h.slice
	fmt.Println("IN MAX HEAPIFY")
	fmt.Println(slice)
	if left <= h.size() && slice[left] > slice[i] {
		largest = left
	} else {
		largest = i
	}
	if right <= h.size() && slice[right] > slice[largest] {
		largest = right
	}
	if largest != i {
		prevLargest := slice[i]
		slice[i] = slice[largest]
		slice[largest] = slice[prevLargest]
		h.MaxHeapify(largest)
	}
}
