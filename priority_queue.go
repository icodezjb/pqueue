// Package Pqueue implements a priority queue data structure supporting arbitrary
// value types and float priorities.
//
// The reasoning behind using floats for the priorities vs. ints or interfaces
// was larger flexibility without sacrificing too much performance or code
// complexity.
//
// If you would like to use a min-priority queue, simply negate the priorities.
//
// Internally the queue is based on the standard heap package working on a
// sortable version of the block based stack.
package pqueue

import "container/heap"

// Priority queue data structure.
type Pqueue struct {
	container *sStack
}

// Creates a new priority queue.
func New() *Pqueue {
	return &Pqueue{newSstack()}
}

// Pushes a value with a given priority into the queue, expanding if necessary.
func (p *Pqueue) Push(data interface{}, priority float32)  {
	heap.Push(p.container, &item{data, priority})
}

// Pops the value with the greates priority off the stack and returns it.
// Currently no shrinking is done.
func (p *Pqueue) Pop() (interface{}, float32)  {
	item := heap.Pop(p.container).(*item)
	return item.value, item.priority
}

// Pops only the value from the queue, dropping the associated priority.
func (p *Pqueue) PopValue() interface{}  {
	return heap.Pop(p.container).(*item).value
}

// Checks whether the priority queue is empty.
func (p *Pqueue) Empty() bool  {
	return p.container.Len() == 0
}

// Returns the number of element in the priority queue.
func (p *Pqueue) Size() int {
	return p.container.Len()
}

// Clears the contents of the priority queue.
func (p *Pqueue) Reset() {
	*p = *New()
}