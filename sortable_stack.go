package pqueue

// The size of a block of data
const blockSize = 4096

// A prioritized item in the sorted stack
type item struct {
	value interface{}
	priority float32
}

// Internal sortable stack data structure.Implement the Push and Pop ops for
// the stack (heap) functionality and the Len, Less and Swap methods for the
// interface of the heap
type sStack struct {
	size int
	capacity int
	offset int

	blocks [][]*item
	active []*item
}

// Create a new, empty stack
func newSstack() (stack *sStack) {
	stack = new(sStack)
	stack.active = make([]*item,blockSize)
	stack.blocks = [][]*item{stack.active}
	stack.capacity = blockSize
	return
}

// Pushes a value onto the stack, expanding it if necessary, Required by
// heap.Interface.
func (s *sStack) Push(data interface{}) {
	if s.offset == s.capacity {
		s.active = make([]*item, blockSize)
		s.blocks = append(s.blocks, s.active)
		s.offset = 0
	} else if s.offset == blockSize {
		s.active = s.blocks[s.size/blockSize]
		s.offset = 0
	}
	s.active[s.offset] = data.(*item)
	s.offset++
	s.size++
}

// Pops a value off the stack and returns it. Currently no shrinking is done.
// Required by heap.Interface.
func (s *sStack) Pop() (data interface{})  {
	s.size--
	s.offset--
	if s.offset < 0 {
		s.offset = blockSize - 1
		s.active = s.blocks[s.size/blockSize]
	}
	data, s.active[s.offset] = s.active[s.offset], nil
	return
}

// Returns the length of the stack. Required by sort.Interface.
func (s *sStack) Len() int {
	return s.size
}

// Compares the priority of two elements of the stack (higher is first).
// Required by sort.Interface.
func (s *sStack) Less(i, j int) bool {
	return s.blocks[i/blockSize][i%blockSize].priority > s.blocks[j/blockSize][j%blockSize].priority
}

// Swaps two elements in the stack. Required by sort.Interface.
func (s *sStack) Swap(i, j int) {
	ib, io, jb, jo := i/blockSize, i%blockSize, j/blockSize, j%blockSize
	s.blocks[ib][io], s.blocks[jb][jo] = s.blocks[jb][jo], s.blocks[ib][io]
}

// Resets the stack, effectively clearing its contents.
func (s *sStack) Reset() {
	*s = *new(sStack)
}