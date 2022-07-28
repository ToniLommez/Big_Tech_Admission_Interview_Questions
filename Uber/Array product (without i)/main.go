package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Good morning! Here's your coding interview problem for today.

This problem was asked by Uber.

Given an array of integers, return a new array such that each element at index i
of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

// SPACE O(n²)
// TIME  O(n)
func function(slice []int) []int {
	slice_1 := make([]int, 0)
	product := 1
	for i := 0; i < len(slice); i++ {
		product = 1
		for j := 0; j < len(slice); j++ {
			if j == i {
				continue
			} else {
				product *= slice[j]
			}
		}
		slice_1 = append(slice_1, product)
	}

	return slice_1
}

// SPACE O(n)
// TIME  O(n)
func function2(slice []int) []int {

	if len(slice) == 1 {
		slice_1 := make([]int, 0)
		slice_1 = append(slice_1, 0)
		return slice_1
	}

	slice_1 := make([]int, len(slice))

	i, temp := 0, 1
	for ; i < len(slice); i++ {
		slice_1[i] = temp
		temp *= slice[i]
	}

	i, temp = len(slice)-1, 1
	for ; i >= 0; i-- {
		slice_1[i] *= temp
		temp *= slice[i]
	}

	return slice_1
}

func IsDigitsOnly(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func main() {
	temp := make([]int, 0)
	temp = append(temp, 1)
	temp = append(temp, 2)
	temp = append(temp, 3)
	temp = append(temp, 4)
	temp = append(temp, 5)
	array1_temp := function2(temp)
	fmt.Print(array1_temp)

	fmt.Println("Given an array of integers, return a new array such that each element at index i")
	fmt.Println("of the new array is the product of all the numbers in the original array except the one at i.")
	fmt.Println("For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24].")
	fmt.Printf("If our input was [3, 2, 1], the expected output would be [2, 3, 6].\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type the array values or 'exit' to finish\n-> ")
	slice := make([]int, 0)
	scanner.Scan()
	value := scanner.Text()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	for value != "exit" {
		if IsDigitsOnly(value) && value != "" {
			temp, _ := strconv.Atoi(value)
			slice = append(slice, temp)
		} else {
			fmt.Printf("Cannot convert value: %s\n", value)
		}
		fmt.Print("-> ")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		value = scanner.Text()
	}

	array1_1 := function(slice)
	array1_2 := function2(slice)

	fmt.Println("\nmethod 1 = ", array1_1)
	fmt.Println("method 2 = ", array1_2)
	fmt.Println("\nPress ENTER to finish")
	scanner.Scan()
}
