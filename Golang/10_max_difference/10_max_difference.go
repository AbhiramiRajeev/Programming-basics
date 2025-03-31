package main

import "fmt"

/*
Problem 10: Find the Maximum Difference in an Array

Given an array, find the maximum difference between any two elements 
such that the larger element appears after the smaller one.

Example Input/Output:
Input: [2, 3, 10, 6, 4, 8, 1]
Output: 8  (10 - 2)

Input: [7, 9, 5, 6, 3, 2]
Output: 2  (9 - 7)

Input: [10, 9, 8, 7]
Output: -1  (No valid pair found)

Hint / Algorithm:
1. Use a variable to track the minimum value encountered.
2. Iterate through the array and compute the difference with the minimum.
3. Keep track of the maximum difference found.
4. Return -1 if no valid difference exists.
*/

func maxDifference(arr []int) int {

}

func main() {
	fmt.Println(maxDifference([]int{2, 3, 10, 6, 4, 8, 1})) // 8
	fmt.Println(maxDifference([]int{7, 9, 5, 6, 3, 2}))    // 2
	fmt.Println(maxDifference([]int{10, 9, 8, 7}))         // -1
	fmt.Println(maxDifference([]int{1}))                   // -1
	fmt.Println(maxDifference([]int{}))                    // -1
}
