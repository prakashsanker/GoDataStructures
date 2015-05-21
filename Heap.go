package main;

import "fmt" 

func main() {
	// testSlice := []int{1,1,1,1,1}
	// fmt.Println(heapSort(testSlice))
	// testSlice = []int{5,1,1,1,1}
	// fmt.Println(heapSort(testSlice))
	testSlice := []int{5,4,3,2,1}
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
	largest := i
	slice := h.slice
	// fmt.Println("left")
	// fmt.Println(left)
	// fmt.Println("right")
	// fmt.Println(right)
	// fmt.Println("i")
	// fmt.Println(i)
	// fmt.Println(h.size())

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
	// fmt.Println("largest")
	// fmt.Println(largest)
	if largest != i {
		prevLargest := slice[i]
		slice[i] = slice[largest]
		slice[largest] = prevLargest
		// fmt.Println(slice)
		// fmt.Println("====")
		h.MaxHeapify(largest)
		//h.MaxHeapify(largest)
	}
}
