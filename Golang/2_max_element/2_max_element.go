package main

import (
	"fmt"
)

/*
Problem 2: Find the Maximum Element in an Array

Given an array of numbers, find the maximum value in the array.

Example Input/Output:
Input: [1, 5, 3, 9, 2]  
Output: 9  

Input: [-5, -1, -8, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize `max` as the first element.
2. Iterate through the array.
3. If an element is greater than `max`, update `max`.
4. Return the `max` value.
5. Avoid using built-in functions like `max()` for better logic building.
*/

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1 // Edge case: Empty array
	}

	max := arr[0]
	for _, num := range arr[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	// Test Cases
	fmt.Println(findMax([]int{1, 5, 3, 9, 2})) // 9
	fmt.Println(findMax([]int{-5, -1, -8, -3})) // -1
	fmt.Println(findMax([]int{10})) // 10
	fmt.Println(findMax([]int{})) // -1 (or return an error instead)
}
