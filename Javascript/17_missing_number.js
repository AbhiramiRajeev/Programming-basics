/*
Problem 17: Find Missing Number

Given an array containing `n` distinct numbers taken from the range 0 to `n`, find the missing number.

Example Input/Output:
Input: [3, 0, 1]
Output: 2

Input: [9,6,4,2,3,5,7,0,1]
Output: 8

Hint / Algorithm:
1. The sum of first `n` natural numbers is `(n * (n + 1)) / 2`.
2. Find the sum of the given array and subtract it from the expected sum.
*/

function findMissingNumber(nums) {
    // Implement the function here
}

// Test Cases
console.log(findMissingNumber([3, 0, 1])); // Output: 2
console.log(findMissingNumber([9,6,4,2,3,5,7,0,1])); // Output: 8
console.log(findMissingNumber([0,1])); // Output: 2
console.log(findMissingNumber([1,2,3,4,5,6,7,8,9,10,12])); // Output: 11
console.log(findMissingNumber([0])); // Output: 1
