#### priority queue

 Package Pqueue implements a priority queue data structure supporting arbitrary
 value types and float priorities.

 The reasoning behind using floats for the priorities vs. ints or interfaces
 was larger flexibility without sacrificing too much performance or code
 complexity.

 If you would like to use a min-priority queue, simply negate the priorities.

 Internally the queue is based on the standard heap package working on a
 sortable version of the block based stack.
 
 #### example
```go

package pqueue_test

import (
	"fmt"
	"github.com/icodezjb/pqueue"
)

// Insert some data into a priority queue and pop them out in prioritized order.
func Example_usage() {
	// Define some data to push into the priority queue
	prio := []float32{77.7, 22.2, 44.4, 55.5, 11.1, 88.8, 33.3, 99.9, 0.0, 66.6}
	data := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Create the priority queue and insert the prioritized data
	pq := pqueue.New()
	for i := 0; i < len(data); i++ {
		pq.Push(data[i], prio[i])
	}
	// Pop out the data and print them
	for !pq.Empty() {
		val, prio := pq.Pop()
		fmt.Printf("%.1f:%s ", prio, val)
	}
	// Output:
	// 99.9:seven 88.8:five 77.7:zero 66.6:nine 55.5:three 44.4:two 33.3:six 22.2:one 11.1:four 0.0:eight
}
```