package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	// last element percolate up to proper place in heap
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	last := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0.")
		return -1
	}
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp percolates index up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		// switch parent and larger child
		h.swap(parent(index), index)
		// now considering new parent for next iteration
		index = parent(index)
	}
}

// maxHeapifyDown will percolate index down
func (h *MaxHeap) maxHeapifyDown(index int) {
	last := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// case where index has at least one child
	for l <= last {
		if l == last { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current idx to larger child and swap if small
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get parent node
func parent(i int) int {
	return (i - 1) / 2
}

// left child
func left(i int) int {
	return 2*i + 1
}

// right child
func right(i int) int {
	return 2*i + 2
}

// swap 2 keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
