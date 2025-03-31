package main

import "fmt"

/*
Problem 7: Reverse an Array

Given an array, return a new array that contains the elements in reverse order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: [5, 4, 3, 2, 1]  

Input: [7, 8, 9]  
Output: [9, 8, 7]  

Input: []  
Output: []  

Hint / Algorithm:
1. Create an empty result slice.
2. Traverse the input slice from the last element to the first.
3. Append each element to the result slice.
*/

func reverseArray(arr []int) []int {
	reversed := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func main() {
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5})) // [5, 4, 3, 2, 1]
	fmt.Println(reverseArray([]int{7, 8, 9}))       // [9, 8, 7]
	fmt.Println(reverseArray([]int{}))             // []
}
