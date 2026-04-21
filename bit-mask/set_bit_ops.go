package main

import "fmt"

// IsSet checks if the bit at position pos is set.
func IsSet(num, pos int) bool {
	return (num & (1 << pos)) != 0
}

// SetBit sets the bit at position pos.
func SetBit(num, pos int) int {
	return num | (1 << pos)
}

// UnsetBit unsets the bit at position pos.
func UnsetBit(num, pos int) int {
	return num & ^(1 << pos)
}

// FlipBit toggles the bit at position pos.
func FlipBit(num, pos int) int {
	return num ^ (1 << pos)
}

func main() {
	num := 10 // Binary: 1010
	pos := 1  // pos is Zero-indexed

	// Check if a bit is set
	isSet := IsSet(num, pos)
	fmt.Printf("Is bit at position %d set in %d? %v\n", pos, num, isSet)

	// Flip a bit
	flipped := FlipBit(num, pos)
	fmt.Printf("Flipping bit at position %d of %d: %d (binary: %b)\n", pos, num, flipped, flipped)

	// Set a bit
	set := SetBit(num, 2)
	fmt.Printf("Setting bit at position 2 of %d: %d (binary: %b)\n", num, set, set)

	// Unset a bit
	unset := UnsetBit(num, 3)
	fmt.Printf("Unsetting bit at position 3 of %d: %d (binary: %b)\n", num, unset, unset)
}
