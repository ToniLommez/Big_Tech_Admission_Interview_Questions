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
func inside(slice []int, k int) bool {
	sort.Ints(slice) // O(n*log(n))
	have := false
	prevent := 10
	fmt.Print("\n\n", slice, "\n")
	for i := 0; i < len(slice); i-- {
		low := i + 1
		high := len(slice) - 1
		needle := int(((high + 1 - low) / 2) + i + 1)
		fmt.Printf("new\n")
		for low < high {
			fmt.Printf("low = %d, needle = %d, high = %d, %d + %d = %d? ", low, needle, high, slice[i], slice[needle], k)
			if slice[i]+slice[needle] == k {
				fmt.Printf("yes\n")
				return true
			} else if slice[i]+slice[needle] > k {
				fmt.Printf("%d greater\n", slice[i]+slice[needle])
				high = needle
				needle = (high+low)/2 + i + 1
			} else {
				fmt.Printf("%d smaller\n", slice[i]+slice[needle])
				low = needle
				needle = (high+low)/2 + i + 1

				prevent--
				if prevent == 0 {
					return have
				}
			}
		}
	}

	/* 	for low <= high {
		median := (low + high) / 2

		if slice[median] < k {
			low = median + 1
		} else {
			high = median - 1
		}
	} */
	fmt.Printf("saiu!\n")
	return have
}

func main() {
	array1 := []int{10, 15, 3, 7}
	array2 := []int{5, 12, 3, 9, 6, 2, 7, 3}
	array3 := []int{20, 12, 3, 78, 5}
	k1 := 17
	k2 := 13
	k3 := 55
	a1 := inside(array1, k1)
	a2 := inside(array2, k2)
	a3 := inside(array3, k3)
	fmt.Printf("%t\n", a1)
	fmt.Printf("%t\n", a2)
	fmt.Printf("%t\n", a3)
}
