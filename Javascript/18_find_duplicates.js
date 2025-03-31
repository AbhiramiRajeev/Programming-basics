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

function findDuplicates(nums) {
    // Implement the function here
}

// Test Cases
console.log(findDuplicates([4,3,2,7,8,2,3,1])); // Output: [2,3]
console.log(findDuplicates([1,1,2])); // Output: [1]
console.log(findDuplicates([1,2,3,4,5])); // Output: []
console.log(findDuplicates([5,4,3,3,2,1,5])); // Output: [3,5]
console.log(findDuplicates([7,7,7,7,7,7,7])); // Output: [7]
