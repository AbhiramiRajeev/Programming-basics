/*
Problem 16: Move Zeros to End

Given an array, move all zeros to the end while maintaining the relative order of the non-zero elements.

Example Input/Output:
Input: [0, 1, 0, 3, 12]
Output: [1, 3, 12, 0, 0]

Input: [1, 0, 2, 0, 3]
Output: [1, 2, 3, 0, 0]

Hint / Algorithm:
1. Use two pointers: One for iterating and another to track the position where non-zero elements should be placed.
2. Traverse the array and swap non-zero elements with zero positions.
*/

function moveZerosToEnd(arr) {
    // Implement the function here
}

// Test Cases
console.log(moveZerosToEnd([0, 1, 0, 3, 12])); // Output: [1, 3, 12, 0, 0]
console.log(moveZerosToEnd([1, 0, 2, 0, 3])); // Output: [1, 2, 3, 0, 0]
console.log(moveZerosToEnd([0, 0, 1])); // Output: [1, 0, 0]
console.log(moveZerosToEnd([4, 5, 0, 6, 0, 7])); // Output: [4, 5, 6, 7, 0, 0]
console.log(moveZerosToEnd([0, 0, 0])); // Output: [0, 0, 0]
