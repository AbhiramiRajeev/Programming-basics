package main

import "fmt"

/*
Problem 15: Remove Duplicates from Sorted Array

Given a sorted array, remove the duplicates **in place** such that each unique 
element appears only once and return the new length of the array.

Example Input/Output:
Input: [1, 1, 2]
Output: 2 (array becomes [1, 2])

Input: [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]
Output: 5 (array becomes [0, 1, 2, 3, 4])

Hint / Algorithm:
1. Use two pointers: One for iteration and another to track the position of unique elements.
2. If the current element is different from the previous one, move it forward.
*/

func removeDuplicates(nums []int) int {
    // Implement the function here
    return 0
}

func main() {
    fmt.Println(removeDuplicates([]int{1, 1, 2}))  // Output: 2
    fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))  // Output: 5
    fmt.Println(removeDuplicates([]int{1, 2, 3, 4, 5}))  // Output: 5
    fmt.Println(removeDuplicates([]int{1, 1, 1, 1, 1}))  // Output: 1
    fmt.Println(removeDuplicates([]int{}))  // Output: 0
}
