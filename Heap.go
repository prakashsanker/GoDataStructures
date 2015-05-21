package main;

import "fmt" 

func main() {
	testSlice := []int{4,1,3,2,16,9, 10, 14, 8, 7}
	fmt.Println(heapSort(testSlice))
	testSlice = []int{4,4,4,4,5}
	fmt.Println(heapSort(testSlice))


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
	fmt.Println(slice)
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
