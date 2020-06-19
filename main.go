package utils

// Result of the Add
type Result struct {
	Sum int
}

// Add x and y together
func Add(x int, y int) (*Result, error) {
	return &Result{x + y}, nil
}
