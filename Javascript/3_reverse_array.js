/*
Problem 3: Reverse an Array

Given an array, reverse it without using built-in functions.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: [5, 4, 3, 2, 1]  

Input: [10, 20, 30]  
Output: [30, 20, 10]  

Input: []  
Output: []  

Hint / Algorithm:
1. Use two pointers: one at the beginning and one at the end.
2. Swap elements at these positions.
3. Move the pointers toward the center until they meet.
4. Avoid using `.reverse()`.
*/

function reverseArray(arr) {
    let left = 0, right = arr.length - 1;
    while (left < right) {
        let temp = arr[left];
        arr[left] = arr[right];
        arr[right] = temp;
        left++;
        right--;
    }
    return arr;
}

// Test Cases
console.log(reverseArray([1, 2, 3, 4, 5])); // [5, 4, 3, 2, 1]
console.log(reverseArray([10, 20, 30])); // [30, 20, 10]
console.log(reverseArray([])); // []
