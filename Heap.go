package main;

//MaxHeapify runs in O(log n)
//Build Max heap runs in O(n)
//Heap sort runs in O(nlogn)
//Increase key runs in O(log n)
//Extract Max runs in O(const + log n) = O(log n)

import (
		"errors"
		"fmt"
)

func main() {
	testSlice := []int{4,1,3,2,16,9, 10, 14, 8, 7}
	fmt.Println(heapSort(testSlice))
	//fmt.Println(heapSort(testSlice))
	// h := BuildMaxHeap(testSlice)

	// h.increaseKey(8, 15)
	// fmt.Println(h.slice)
	// heapSort1(testSlice)

}

type MaxHeap struct {
	slice []int
	heapSize int
}


func (h MaxHeap) size() int {
	return h.heapSize
}

func (h MaxHeap) max() int{
	return h.slice[0]
}

func (h MaxHeap) extractMax() int {
	if h.size() < 1 {
		err := errors.New("Heap underflow")
		fmt.Println(err)
	}
	//remove and return largest element
	largestElem := h.slice[0]
	h.slice[0] = h.slice[h.size()-1]
	h.MaxHeapify(0)
	return largestElem
}

func (h MaxHeap) increaseKey(i int, key int) {
	if key < h.slice[i] {
		err := errors.New("New key smaller than current key")
		fmt.Println(err)
	}
	h.slice[i] = key
	index := i
	parentVal := h.slice[(i+1)/2]
	for parentVal < key && index > 0 {
		h.slice[(index+1)/2], h.slice[index] = h.slice[index], h.slice[(index+1)/2]
		index = (index +1)/2
		parentVal = h.slice[(index+1)/2]
	}
}


func heapSort(slice []int) []int {
	h := BuildMaxHeap(slice)
	for i := len(h.slice) - 1; i > 0 ; i-- {
		first := h.slice[0]
		last := h.slice[i]
		h.slice[0] = last
		h.slice[i] = first
		h.heapSize--
		h.MaxHeapify(0)

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
	left := 2*i + 1
	right := 2*i + 2
	largest := i
	slice := h.slice

	if left < h.size() {
		if slice[left] > slice[i] {
			largest = left
		} else {
			largest = i
		}
	}
	if right < h.size() {
		if slice[right] > slice[largest] {
			largest = right
		}
	}
	if largest != i {
		slice[i], slice[largest] = slice[largest], slice[i]
		h.MaxHeapify(largest)
	}
}
