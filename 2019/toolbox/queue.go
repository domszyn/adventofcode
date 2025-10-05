package toolbox

import "fmt"

// IntQueue is an efficient, slice-based queue for integers.
type Queue[T any] struct {
	items []T
}

// NewIntQueue creates a new queue.
func NewQueue[T any](elems ...T) *Queue[T] {
	if len(elems) == 0 {
		return &Queue[T]{items: []T{}}

	}

	items := make([]T, len(elems))
	copy(items, elems)
	return &Queue[T]{items: items}
}

// Enqueue adds an element to the back of the queue.
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the element from the front of the queue.
// Returns an error if the queue is empty.
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T

	if len(q.items) == 0 {
		return zero, fmt.Errorf("queue is empty")
	}

	item := q.items[0]
	q.items[0] = zero
	q.items = q.items[1:]

	// **The key to efficiency**: Periodically shrink the underlying array.
	// If capacity is more than double the length, we're wasting space.
	if cap(q.items) > 2*len(q.items) {
		newItems := make([]T, len(q.items))
		copy(newItems, q.items)
		q.items = newItems
	}

	return item, nil
}

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.Size() == 0
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return len(q.items)
}
