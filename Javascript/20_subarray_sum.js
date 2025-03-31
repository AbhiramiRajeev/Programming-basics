/*
Problem 20: Subarray with Given Sum

Given an array of integers and a target sum, find a continuous subarray that adds up to the target sum.
Return the starting and ending indices (0-based). If no such subarray exists, return an empty array.

Example Input/Output:
Input: ([1, 2, 3, 7, 5], 12)
Output: [2, 4]  (subarray [3,7,5] adds up to 12)

Input: ([1, 4, 20, 3, 10, 5], 33)
Output: [2, 4]  (subarray [20, 3, 10] adds up to 33)

Input: ([1, 2, 3, 4, 5], 15)
Output: [0, 4]  (subarray [1,2,3,4,5] adds up to 15)

Input: ([5, 6, 7, 8, 9], 50)
Output: []  (No subarray found)

Hint / Algorithm:
1. Use a sliding window technique to track the sum of elements.
2. Expand the window by adding elements until the sum exceeds the target.
3. If sum exceeds the target, shrink the window from the left.
*/

function subarrayWithSum(arr, target) {
    // Implement the function here
}

// Test Cases
console.log(subarrayWithSum([1, 2, 3, 7, 5], 12)); // Output: [2, 4]
console.log(subarrayWithSum([1, 4, 20, 3, 10, 5], 33)); // Output: [2, 4]
console.log(subarrayWithSum([1, 2, 3, 4, 5], 15)); // Output: [0, 4]
console.log(subarrayWithSum([5, 6, 7, 8, 9], 50)); // Output: []
console.log(subarrayWithSum([1, 2, 3, 4, 5, 6], 6)); // Output: [0, 2]
