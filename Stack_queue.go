package main

import "fmt"

/*
	Stack : Last in First out(LIFO) structure
		- PUSH : putting an item or data into the stck
		- POP : take it out

	Queue : First in First out (FIFO) structure
	 	- eg. Line in the supermarket
		- ENQUEUE : add an item into a queue
		- DEQUEUE : take it out
*/

// Stack represents stack that hold a slice
type Stack struct {
	items []int 
}

// Queue represents stack that hold a slice
type Queue struct {
	items []int 
}

// Push -> add value at the end 
func (s *Stack) Push (i int) {
	s.items = append(s.items, i)
}

// Pop -> remove a value at the end 
//and returns the removed value 
func (s *Stack) Pop() int {
	l := len(s.items) - 1

	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

// Enqueue -> adds a value at the end
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue -> removes a value at the front
// and returns the removes value
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:] //items[0] 제거 -> items[1, ..., n] 이런 느낌으로다 
	return toRemove
}