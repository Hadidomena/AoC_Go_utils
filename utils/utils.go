
package utils

import (
	"fmt"
	"os"
	"strings"
)

type Queue[T any] []T

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(value T) {
	*q = append(*q, value)
}

// Dequeue removes and returns the front element from the queue.
// If the queue is empty, it returns the zero value of T and false.
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(*q) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := (*q)[0]
	*q = (*q)[1:]
	return value, true
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}

// LoadFile takes the path to the file as an input
// and returns input file split into lines
func LoadFile(input string) []string {
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	content := string(data)
	lines := strings.Split(content, "\r\n")
	return lines
}
// ConvertToRuneBoard takes in an array of strings,
// and converts it into separate characters to allow easier modifications
func ConvertToRuneBoard(board []string) [][]rune {
	runeBoard := make([][]rune, len(board))
	for i, line := range board {
		runeBoard[i] = []rune(line)
	}
	return runeBoard
}

// PriorityQueue is a generic priority queue implementation based on the heap interface.
// It allows storing and prioritizing elements based on a custom comparison function.
type PriorityQueue[T any] struct {
	items []T
	less  func(a, b T) bool
}

// Len returns the number of items in the priority queue.
// Implements the sort.Interface.
func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

// Less determines the priority of items in the queue.
// Implements the sort.Interface.
func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i], pq.items[j])
}

// Swap exchanges the positions of two items in the queue.
// Implements the sort.Interface.
func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// Push adds a new item to the priority queue.
// Implements the heap.Interface.
func (pq *PriorityQueue[T]) Push(x any) {
	pq.items = append(pq.items, x.(T))
}

// Pop removes and returns the last item from the priority queue.
// Implements the heap.Interface.
func (pq *PriorityQueue[T]) Pop() any {
	n := len(pq.items)
	item := pq.items[n-1]
	pq.items = pq.items[:n-1]
	return item
}

// NewPriorityQueue creates a new priority queue with a custom less function.
// The less function defines the priority of elements (return true if a should be 
// prioritized before b).
func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: []T{},
		less:  less,
	}
}
