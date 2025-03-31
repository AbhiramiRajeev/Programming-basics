package main

import (
	"fmt"
)

/*
Problem 3: Reverse an Array

Given an array, reverse it without using built-in functions.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: [5, 4, 3, 2, 1]  

Input: [10, 20, 30]  
Output: [30, 20, 10]  

Input: []  
Output: []  

Hint / Algorithm:
1. Use two pointers: one at the beginning and one at the end.
2. Swap elements at these positions.
3. Move the pointers toward the center until they meet.
*/

func reverseArray(arr []int) []int {

}

func main() {
	fmt.Println(reverseArray([]int{1, 2, 3, 4, 5})) // [5, 4, 3, 2, 1]
	fmt.Println(reverseArray([]int{10, 20, 30})) // [30, 20, 10]
	fmt.Println(reverseArray([]int{})) // []
}
