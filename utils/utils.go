
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
