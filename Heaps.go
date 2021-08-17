/*
Normal Queue
 - First in first out

Priority Queues
 - Highest priority served first

Max Heap
 - Root : Max

Min heap
 - Root : Min

 Inserting and extracting node in the heap
  - end of / index of array -> ignore
  - depend on how high of the tree is : O(h) (h: height of tree)
  - Heap inserting & Extract -> O(log n)


   Accesible from outside
	struct : MaxHeap
	Method  - Insert
		   |
		    - Extract
	===================
	method  -- maxHeapifyUp
		   | - maxHeapifyDown
		    -- swap
	function - parent
		    |- left
			 - right
*/

package main

import "fmt"

// Maxheap struct has a slice that holds the array
type MaxHeap struct{
	array []int 
}

// Insert adds an element to the heap
func ( h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// When the array is empty 
	if len(h.array) == 0 {
		fmt.Println("cannot extract beacause array length is 0")
		return -1 
	}
	// take out the last index and put it in the root 
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyUp(0)

	return extracted
}

// maxHeapifyUp will heapify bottom top
func(h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index]  {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyUp will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int){

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex{
		if l == lastIndex { //when left child index is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger 
			childToCompare = r
		}
		//compare array value of current index to larger chold and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}


// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index 
func left (i int) int {
	return 2 * i  + 1
}

// get the right child index 
func right (i int) int {
	return 2 * i  + 2
}

// swap keys in the array 
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

/*
func main() {
	m := &MaxHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 30, 3, 4, 6, 11, 34}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

	fmt.Println(right(3), left(4))
}
*/