/*
Problem 1: Remove All Occurrences of an Element in an Array

Given an array and a value, remove all occurrences of the value from the array and return the modified array.

Example Input/Output:
Input: array = [1, 2, 3, 4, 2, 2, 5], element = 2  
Output: [1, 3, 4, 5]  

Input: array = [4, 4, 4, 4], element = 4  
Output: []  

Input: array = [1, 3, 5, 7], element = 2  
Output: [1, 3, 5, 7]  

Hint / Algorithm:
1. Initialize an empty list.
2. Iterate over the given array.
3. If an element is NOT equal to the given value, add it to the new list.
4. Return the new list.
5. Avoid using inbuilt functions like `filter()` for better logic-building practice.
*/

function removeOccurrences(arr, element) {

}

// Test Cases
console.log(removeOccurrences([1, 2, 3, 4, 2, 2, 5], 2)); // [1, 3, 4, 5]
console.log(removeOccurrences([4, 4, 4, 4], 4)); // []
console.log(removeOccurrences([1, 3, 5, 7], 2)); // [1, 3, 5, 7]
console.log(removeOccurrences([], 5)); // []
