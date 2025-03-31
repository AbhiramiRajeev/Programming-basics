package main

import "fmt"

/*
Problem 4: Remove Duplicates from an Array

Given an array, remove duplicate values and return the unique elements while maintaining the order.

Example Input/Output:
Input: [1, 2, 3, 1, 2, 4, 5]  
Output: [1, 2, 3, 4, 5]  

Input: [7, 8, 9, 7, 7, 8]  
Output: [7, 8, 9]  

Input: []  
Output: []  

Hint / Algorithm:
1. Create an empty slice for unique values.
2. Iterate through the input slice.
3. Add each element to the result only if it is not already present.
*/

func removeDuplicates(arr []int) []int {
	unique := []int{}
	exists := make(map[int]bool)

	for _, num := range arr {
		if !exists[num] {
			unique = append(unique, num)
			exists[num] = true
		}
	}
	return unique
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 2, 3, 1, 2, 4, 5})) // [1, 2, 3, 4, 5]
	fmt.Println(removeDuplicates([]int{7, 8, 9, 7, 7, 8})) // [7, 8, 9]
	fmt.Println(removeDuplicates([]int{})) // []
}
