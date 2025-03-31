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

function maxDifference(arr) {
    if (arr.length < 2) return -1;
    
    let minVal = arr[0];
    let maxDiff = -1;

    for (let i = 1; i < arr.length; i++) {
        if (arr[i] > minVal) {
            maxDiff = Math.max(maxDiff, arr[i] - minVal);
        } else {
            minVal = arr[i];
        }
    }

    return maxDiff;
}

// Test Cases
console.log(maxDifference([2, 3, 10, 6, 4, 8, 1])); // 8
console.log(maxDifference([7, 9, 5, 6, 3, 2])); // 2
console.log(maxDifference([10, 9, 8, 7])); // -1
console.log(maxDifference([1])); // -1
console.log(maxDifference([])); // -1
