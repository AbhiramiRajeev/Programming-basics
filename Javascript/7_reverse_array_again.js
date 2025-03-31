/*
Problem 7: Reverse an Array

Given an array, return a new array that contains the elements in reverse order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: [5, 4, 3, 2, 1]  

Input: [7, 8, 9]  
Output: [9, 8, 7]  

Input: []  
Output: []  

Hint / Algorithm:
1. Create an empty result array.
2. Traverse the input array from the last element to the first.
3. Append each element to the result array.
*/

function reverseArray(arr) {
    let reversed = [];
    for (let i = arr.length - 1; i >= 0; i--) {
        reversed.push(arr[i]);
    }
    return reversed;
}

// Test Cases
console.log(reverseArray([1, 2, 3, 4, 5])); // [5, 4, 3, 2, 1]
console.log(reverseArray([7, 8, 9])); // [9, 8, 7]
console.log(reverseArray([])); // []
