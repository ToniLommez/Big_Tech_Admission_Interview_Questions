package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
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

func randomSlice() []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, 0)
	temp := 0
	for i := 0; i < 100000000; i++ {
		temp = (rand.Int() % 7000000)
		if rand.Int()%6 == 1 {
			temp *= -1
		}
		slice = append(slice, temp)
	}
	return slice
}

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
	// missing++
	return missing
}

func firstMissingPositive(slice []int) int {
	start := 0
	slice, start = withdrawOutRange(slice)
	// fmt.Printf("%v\n", slice)
	for i := start; i < len(slice); i++ {
		if abs(slice[i]) < len(slice)+1 {
			if slice[abs(slice[i])-1+start] > 0 {
				slice[abs(slice[i])-1+start] *= -1
			}
		}
	}
	// fmt.Printf("%v\n", slice)
	missing := firstMissingPositiveSorted(slice, start)
	return missing
}

func main() {
	start := time.Now().UnixMilli()
	dt := time.Now()

	/* 	array1 := []int{12, 3, 5, 8, 7, 4, 1, 9, 2}
	fmt.Printf("array = %v\n", array1)
	fmt.Printf("Smallest missing positive value = %d\n\n", firstMissingPositive(array1))

	array2 := []int{1, 2, 0}
	fmt.Printf("array = %v\n", array2)
	fmt.Printf("Smallest missing positive value = %d\n\n", firstMissingPositive(array2)) */

	array3 := randomSlice()
	// fmt.Printf("array = %v\n", array3)
	fmt.Printf("Smallest missing positive value = %d\n", firstMissingPositive(array3))

	fmt.Printf("\nArray size = 100.000.000\nArray range = 7.000.000\n")
	fmt.Printf("\nStart Execution Time = ")
	fmt.Println(dt.Format("15:04:05.000"))
	end := time.Now().UnixMilli()
	dt = time.Now()
	fmt.Printf("End Execution Time   = ")
	fmt.Println(dt.Format("15:04:05.000"))
	finaltime := float64(end-start) / 1000
	fmt.Printf("\nTotal execution Time = %.3f seconds\n\n", finaltime)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nPress ENTER to finish")
	scanner.Scan()
}
