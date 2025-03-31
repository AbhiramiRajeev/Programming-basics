package main

import "fmt"

/*
Problem 8: Check if an Array is Sorted

Given an array of numbers, determine if it is sorted in non-decreasing order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]
Output: true

Input: [3, 2, 1]
Output: false

Input: [7, 7, 8, 9]
Output: true

Hint / Algorithm:
1. Iterate through the array.
2. Compare each element with the next one.
3. If any element is greater than the next, return false.
4. If loop completes, return true.
*/

func isSorted(arr []int) bool {

}

func main() {
	fmt.Println(isSorted([]int{1, 2, 3, 4, 5})) // true
	fmt.Println(isSorted([]int{3, 2, 1}))       // false
	fmt.Println(isSorted([]int{7, 7, 8, 9}))    // true
	fmt.Println(isSorted([]int{}))              // true
	fmt.Println(isSorted([]int{10}))            // true
}
