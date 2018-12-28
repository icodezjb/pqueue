#### priority queue

 Package Pqueue implements a priority queue data structure supporting arbitrary
 value types and float priorities.

 The reasoning behind using floats for the priorities vs. ints or interfaces
 was larger flexibility without sacrificing too much performance or code
 complexity.

 If you would like to use a min-priority queue, simply negate the priorities.

 Internally the queue is based on the standard heap package working on a
 sortable version of the block based stack.