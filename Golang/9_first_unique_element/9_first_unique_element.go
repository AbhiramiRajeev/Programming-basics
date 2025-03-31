package main

import "fmt"

/*
Problem 9: Find the First Non-Repeating Element

Given an array of numbers, find the first element that does not repeat.

Example Input/Output:
Input: [4, 5, 1, 2, 0, 4, 5, 1]
Output: 2

Input: [7, 8, 7, 8, 9]
Output: 9

Input: [3, 3, 3]
Output: -1

Hint / Algorithm:
1. Use a map to count occurrences.
2. Iterate again to find the first element with count 1.
3. Return that element or -1 if all are repeating.
*/

func firstNonRepeating(arr []int) int {

}

func main() {
	fmt.Println(firstNonRepeating([]int{4, 5, 1, 2, 0, 4, 5, 1})) // 2
	fmt.Println(firstNonRepeating([]int{7, 8, 7, 8, 9}))          // 9
	fmt.Println(firstNonRepeating([]int{3, 3, 3}))                // -1
	fmt.Println(firstNonRepeating([]int{}))                       // -1
}
