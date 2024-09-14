# Circular Buffer (Ring Buffer)

A thread-safe implementation of a **Circular Buffer** (also known as a **Ring Buffer**) in Go, designed for efficient data storage and retrieval in a fixed-size buffer. The buffer automatically overwrites the oldest values when it's full, ensuring continuous data processing without requiring resizing.

## Features

- **Fixed Size**: The buffer has a fixed capacity defined during initialization.
- **Thread-Safe**: Operations are safe for concurrent access, using `sync.RWMutex` to ensure safe reads and writes.
- **Overwrite on Full**: When the buffer is full, new values overwrite the oldest ones.
- **Efficient Pushing and Popping**: Supports pushing new elements and popping the oldest element.
- **Average Calculation**: Computes the average of all the current elements in the buffer.
- **Print Buffer Contents**: Prints all current elements in the buffer in order from oldest to newest.

## Data Structure

```go
type CircularBuffer struct {
    buffer []float64    // The underlying slice to store the data
    size   int          // The fixed size of the buffer
    head   int          // Index where the next element will be inserted
    tail   int          // Index of the oldest element
    count  int          // Number of elements currently in the buffer
    mutex  sync.RWMutex // Mutex for thread-safety 
}
