package main

import "fmt"

/*
Problem 19: Longest Consecutive Sequence

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Example Input/Output:
Input: [100, 4, 200, 1, 3, 2]
Output: 4 (because [1, 2, 3, 4] is the longest sequence)

Input: [9,1,4,7,3,-1,0,5,8,-3,6,2]
Output: 7 ([-1,0,1,2,3,4,5,6])

Hint / Algorithm:
1. Store all elements in a set for quick lookup.
2. Iterate through the array and find the start of a sequence.
3. Expand the sequence by checking consecutive numbers in the set.
*/

func longestConsecutive(nums []int) int {
    // Implement the function here
    return 0
}

func main() {
    fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // Output: 4
    fmt.Println(longestConsecutive([]int{9,1,4,7,3,-1,0,5,8,-3,6,2})) // Output: 7
    fmt.Println(longestConsecutive([]int{1,2,0,1})) // Output: 3
    fmt.Println(longestConsecutive([]int{})) // Output: 0
    fmt.Println(longestConsecutive([]int{10,20,30,40})) // Output: 1
}
