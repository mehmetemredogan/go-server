package array

// InArray
// Searches for a value in the array.
// Syntax: array, value
// Output;
// - Not Found: -1
// - Found: Position in array
func InArray(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}