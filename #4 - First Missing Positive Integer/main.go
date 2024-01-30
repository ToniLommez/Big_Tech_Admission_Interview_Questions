package main

import (
	"fmt"
)

/* This problem was asked by Stripe.

Given an array of integers, find the first missing positive integer in linear
time and constant space. In other words, find the lowest positive integer that
does not exist in the array. The array can contain duplicates and negative
numbers as well.

For example,
the input [3, 4, -1, 1] should give 2.
The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

func abs(value int) int {
	if value < 0 {
		value *= -1
	}
	return value
}

func withdrawOutRange(slice []int) ([]int, int) {
	temp := 0
	needle := 0
	for i := 1; i < len(slice); i++ {
		if slice[i] < 1 {
			temp = slice[needle]
			slice[needle] = slice[i]
			slice[i] = temp
			needle++
		}
	}
	return slice, needle
}

func firstMissingPositiveSorted(slice []int, start int) int {
	missing := 1
	for i := start; i < len(slice); i++ {
		if slice[i] > 0 {
			return missing
		}
		missing++
	}

	return missing
}

func firstMissingPositive(slice []int) int {
	start := 0
	slice, start = withdrawOutRange(slice)

	for i := start; i < len(slice); i++ {
		if abs(slice[i]) < len(slice)+1 {
			if slice[abs(slice[i])-1+start] > 0 {
				slice[abs(slice[i])-1+start] *= -1
			}
		}
	}

	missing := firstMissingPositiveSorted(slice, start)
	return missing
}

func main() {
	array1 := []int{12, 3, 5, 8, 7, 4, 1, 9, 2}
	fmt.Printf("array = %v\n", array1)
	fmt.Printf("Smallest missing positive value = %d\n\n", firstMissingPositive(array1))

	array2 := []int{1, 2, 0}
	fmt.Printf("array = %v\n", array2)
	fmt.Printf("Smallest missing positive value = %d\n\n", firstMissingPositive(array2))
}
