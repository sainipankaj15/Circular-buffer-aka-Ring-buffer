package main

import (
	"fmt"
	"sync"
)

// CircularBuffer represents a fixed-size circular buffer for float64 values.
type CircularBuffer struct {
	buffer []float64    // The underlying slice to store the data
	size   int          // The fixed size of the buffer
	head   int          // Index where the next element will be inserted
	tail   int          // Index of the oldest element
	count  int          // Number of elements currently in the buffer
	mutex  sync.RWMutex // Mutex for thread-safety 
}

// NewCircularBuffer creates and returns a new CircularBuffer with the specified size.
func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		buffer: make([]float64, size),
		size:   size,
	}
}

// Push adds a new item to the buffer and returns the item that was removed (if any).
func (cb *CircularBuffer) Push(item float64) float64 {
	cb.mutex.Lock() // write lock
	defer cb.mutex.Unlock()

	var poppedItem float64

	if cb.count == cb.size {
		// Buffer is full, remove the oldest item
		poppedItem = cb.buffer[cb.tail]
		cb.tail = (cb.tail + 1) % cb.size
	} else {
		// Buffer is not full, increase the count
		cb.count++
	}

	// Add the new item at the head
	cb.buffer[cb.head] = item
	// Move the head forward, wrapping around if necessary
	cb.head = (cb.head + 1) % cb.size

	return poppedItem
}

// Pop removes and returns the oldest item from the buffer.
// The boolean return value indicates whether an item was successfully removed.
func (cb *CircularBuffer) Pop() (float64, bool) {
	cb.mutex.Lock() // Write lock
	defer cb.mutex.Unlock()

	if cb.count == 0 {
		// Buffer is empty
		return 0, false
	}

	item := cb.buffer[cb.tail]
	// Move the tail forward, wrapping around if necessary
	cb.tail = (cb.tail + 1) % cb.size
	cb.count--

	return item, true
}

// Average calculates and returns the average of all items in the buffer.
func (cb *CircularBuffer) Average() float64 {
	cb.mutex.RLock() // Read Lock
	defer cb.mutex.RUnlock()

	if cb.count == 0 {
		return 0
	}

	sum := 0.0
	for i := 0; i < cb.count; i++ {
		// Calculate the actual index, wrapping around if necessary
		index := (cb.tail + i) % cb.size
		sum += cb.buffer[index]
	}

	return sum / float64(cb.count)
}

// PrintBuffer prints all elements in the buffer from tail to head.
func (cb *CircularBuffer) PrintBuffer() {
	cb.mutex.RLock() // Read lock 
	defer cb.mutex.RUnlock()

	fmt.Print("Buffer contents (tail to head): ")
	if cb.count == 0 {
		fmt.Println("empty")
		return
	}

	for i := 0; i < cb.count; i++ {
		// Calculate the actual index, wrapping around if necessary
		index := (cb.tail + i) % cb.size
		fmt.Printf("%.2f ", cb.buffer[index])
	}
	fmt.Println()
}

func main() {
	// Create a new circular buffer with size 5 (smaller for demonstration)
	cb := NewCircularBuffer(5)

	// Example usage: Push 7 items into the buffer
	for i := 0; i < 7; i++ {
		popped := cb.Push(float64(i))
		fmt.Printf("Pushed %.2f", float64(i))
		if i >= 5 {
			// After 5 pushes, the buffer starts overwriting old values
			fmt.Printf(", Popped %.2f", popped)
		}
		fmt.Println()

		// Print the buffer state after each push
		cb.PrintBuffer()
		fmt.Println() // Extra line for readability
	}

	// Calculate and print the average of all items in the buffer
	fmt.Printf("Average: %.2f\n", cb.Average())
}
