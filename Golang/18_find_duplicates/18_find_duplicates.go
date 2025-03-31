package main

import "fmt"

/*
Problem 18: Find All Duplicates in an Array

Given an array of integers where some elements appear twice and others once, return a list of all duplicate numbers.

Example Input/Output:
Input: [4,3,2,7,8,2,3,1]
Output: [2,3]

Input: [1,1,2]
Output: [1]

Hint / Algorithm:
1. Use a hash map to track occurrences of numbers.
2. Iterate through the array and add numbers that appear more than once to the result list.
*/

func findDuplicates(nums []int) []int {
    // Implement the function here
    return []int{}
}

func main() {
    fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1})) // Output: [2,3]
    fmt.Println(findDuplicates([]int{1,1,2})) // Output: [1]
    fmt.Println(findDuplicates([]int{1,2,3,4,5})) // Output: []
    fmt.Println(findDuplicates([]int{5,4,3,3,2,1,5})) // Output: [3,5]
    fmt.Println(findDuplicates([]int{7,7,7,7,7,7,7})) // Output: [7]
}
