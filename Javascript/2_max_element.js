/*
Problem 2: Find the Maximum Element in an Array

Given an array of numbers, find the maximum value in the array.

Example Input/Output:
Input: [1, 5, 3, 9, 2]  
Output: 9  

Input: [-5, -1, -8, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize `max` as the first element.
2. Iterate through the array.
3. If an element is greater than `max`, update `max`.
4. Return the `max` value.
5. Avoid using `Math.max()` to build logic.
*/

function findMax(arr) {
    if (arr.length === 0) return null; // Edge case: Empty array

    let max = arr[0];
    for (let i = 1; i < arr.length; i++) {
        if (arr[i] > max) {
            max = arr[i];
        }
    }
    return max;
}

// Test Cases
console.log(findMax([1, 5, 3, 9, 2])); // 9
console.log(findMax([-5, -1, -8, -3])); // -1
console.log(findMax([10])); // 10
console.log(findMax([])); // null
