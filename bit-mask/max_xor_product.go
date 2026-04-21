package main

import "fmt"

const mod = 1_000_000_007

func main() {
	maximumXorProduct := func(a int64, b int64, n int) int {
		// Iterate over all n bits from n-1 to 0
		for i := n - 1; i >= 0; i-- {
			bit := int64(1) << i

			// if at current bit position if both a and b have same bit make them both 1 to maximize product
			if a&bit == b&bit {
				a |= bit
				b |= bit
			} else {
				// else give the smaller number 1 and make b's bit value 0.
				// Since for a constant sum product is maximized when numbers are closest
				if a > b {
					a, b = b, a
				}

				a |= bit
				b &= ^bit // since we are XOR other should get 0 as a and b are different
			}
			fmt.Printf("A: %b, B: %b \n", a, b)
		}
		a %= mod
		b %= mod

		return int(a*b) % mod

	}
	fmt.Println(maximumXorProduct(12, 5, 4))
}
