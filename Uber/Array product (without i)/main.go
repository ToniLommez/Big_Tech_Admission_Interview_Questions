package main

import (
	"fmt"
	"sort"
)

/*
Given a list of numbers and a number k, return whether any two numbers
from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

// O = n²
// Theta = n²
func inside(slice []int, k int) (bool, int, int) {
	sort.Ints(slice) // O(n*log(n))
	have := false
	low := 0
	high := 0
	needle := 0
	i := 0
	for i = 0; i < len(slice)-1; i++ {
		low = i + 1
		high = len(slice) - 1
		needle = (high + low) / 2
		for low <= high {
			if slice[i]+slice[needle] == k {
				return true, needle, i
			} else if slice[i]+slice[needle] > k {
				high = needle - 1
				needle = (high + low) / 2
			} else {
				low = needle + 1
				needle = (high + low) / 2
			}
		}
	}
	return have, needle, i
}

func main() {
	array1 := []int{10, 15, 3, 7}
	array2 := []int{5, 12, 3, 9, 6, 2, 7, 3}
	array3 := []int{20, 12, 3, 78, 5}
	k1 := 17
	k2 := 13
	k3 := 55
	a1, needle1, i1 := inside(array1, k1)
	a2, needle2, i2 := inside(array2, k2)
	a3, needle3, i3 := inside(array3, k3)

	fmt.Printf("Given a list of numbers and a number k, return whether any two numbers\n")
	fmt.Printf("from the list add up to k.\n")
	fmt.Printf("For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.\n")
	fmt.Printf("Bonus: Can you do this in one pass?\n\n")

	fmt.Printf("%v\n%d + %d = %d\n%t\n\n", array1, array1[i1], array1[needle1], k1, a1)
	fmt.Printf("%v\n%d + %d = %d\n%t\n\n", array2, array2[i2], array2[needle2], k2, a2)
	fmt.Printf("%v\n%d + %d = %d\n%t\n\n", array3, array3[i3], array3[needle3], k3, a3)
}
