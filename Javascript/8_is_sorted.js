/*
Problem 8: Check if an Array is Sorted

Given an array of numbers, determine if it is sorted in non-decreasing order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: true  

Input: [3, 2, 1]  
Output: false  

Input: [7, 7, 8, 9]  
Output: true  

Hint / Algorithm:
1. Iterate through the array.
2. Compare each element with the next one.
3. If any element is greater than the next, return false.
4. If loop completes, return true.
*/

function isSorted(arr) {
    for (let i = 0; i < arr.length - 1; i++) {
        if (arr[i] > arr[i + 1]) {
            return false;
        }
    }
    return true;
}

// Test Cases
console.log(isSorted([1, 2, 3, 4, 5])); // true
console.log(isSorted([3, 2, 1])); // false
console.log(isSorted([7, 7, 8, 9])); // true
console.log(isSorted([])); // true
console.log(isSorted([10])); // true
