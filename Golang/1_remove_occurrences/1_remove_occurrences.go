package main

import (
	"fmt"
)

/*
Problem 1: Remove All Occurrences of an Element in an Array

Given an array and a value, remove all occurrences of the value from the array and return the modified array.

Example Input/Output:
Input: array = [1, 2, 3, 4, 2, 2, 5], element = 2  
Output: [1, 3, 4, 5]  

Input: array = [4, 4, 4, 4], element = 4  
Output: []  

Input: array = [1, 3, 5, 7], element = 2  
Output: [1, 3, 5, 7]  

Hint / Algorithm:
1. Create a new slice to store the result.
2. Iterate over the given slice.
3. If an element is NOT equal to the given value, append it to the new slice.
4. Return the new slice.
5. Avoid using built-in filtering functions for better logic-building practice.
*/

func removeOccurrences(arr []int, element int) []int {
	var result []int
	for _, num := range arr {
		if num != element {
			result = append(result, num)
		}
	}
	return result
}

func main() {
	// Test Cases
	fmt.Println(removeOccurrences([]int{1, 2, 3, 4, 2, 2, 5}, 2)) // [1, 3, 4, 5]
	fmt.Println(removeOccurrences([]int{4, 4, 4, 4}, 4))           // []
	fmt.Println(removeOccurrences([]int{1, 3, 5, 7}, 2))           // [1, 3, 5, 7]
	fmt.Println(removeOccurrences([]int{}, 5))                    // []
}
