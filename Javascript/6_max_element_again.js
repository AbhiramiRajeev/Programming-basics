/*
Problem 6: Find the Maximum Element in an Array

Given an array of numbers, find and return the maximum element.

Example Input/Output:
Input: [3, 1, 9, 7, 4]  
Output: 9  

Input: [-5, -1, -7, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize max as the first element.
2. Iterate through the array, updating max whenever a larger element is found.
3. Return max after the loop.
4. Avoid using Math.max() to practice logic-building.
*/

function findMax(arr) {
    if (arr.length === 0) return null; // Handle empty array case
    let max = arr[0];
    for (let num of arr) {
        if (num > max) {
            max = num;
        }
    }
    return max;
}

// Test Cases
console.log(findMax([3, 1, 9, 7, 4])); // 9
console.log(findMax([-5, -1, -7, -3])); // -1
console.log(findMax([10])); // 10
console.log(findMax([])); // null
