package main

import "fmt"

/*
Problem 12: Find the Intersection of Two Arrays

Given two arrays, find their intersection (common elements).

Example Input/Output:
Input: [1, 2, 2, 1], [2, 2]
Output: [2]

Input: [4, 9, 5], [9, 4, 9, 8, 4]
Output: [4, 9]

Hint / Algorithm:
1. Store elements of one array in a set.
2. Iterate through the second array and check for common elements.
3. Return unique common elements.
*/

func arrayIntersection(arr1, arr2 []int) []int {
	set := make(map[int]bool)
	result := []int{}

	for _, num := range arr1 {
		set[num] = true
	}

	for _, num := range arr2 {
		if set[num] {
			result = append(result, num)
			delete(set, num)
		}
	}

	return result
}

func main() {
	fmt.Println(arrayIntersection([]int{1, 2, 2, 1}, []int{2, 2})) // [2]
	fmt.Println(arrayIntersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // [4, 9]
}
