package main

import "fmt"

/*
Problem 6: Find the Maximum Element in an Array

Given an array of numbers, find and return the maximum element.

Example Input/Output:
Input: [3, 1, 9, 7, 4]  
Output: 9  

Input: [-5, -1, -7, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize max as the first element.
2. Iterate through the array, updating max whenever a larger element is found.
3. Return max after the loop.
*/

func findMax(arr []int) int {
	if len(arr) == 0 {
		return -1 // Return -1 to indicate an empty array case
	}
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(findMax([]int{3, 1, 9, 7, 4}))  // 9
	fmt.Println(findMax([]int{-5, -1, -7, -3})) // -1
	fmt.Println(findMax([]int{10}))             // 10
	fmt.Println(findMax([]int{}))               // -1
}
