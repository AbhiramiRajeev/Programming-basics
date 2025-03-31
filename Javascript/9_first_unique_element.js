/*
Problem 9: Find the First Non-Repeating Element

Given an array of numbers, find the first element that does not repeat.

Example Input/Output:
Input: [4, 5, 1, 2, 0, 4, 5, 1]  
Output: 2  

Input: [7, 8, 7, 8, 9]  
Output: 9  

Input: [3, 3, 3]  
Output: -1  

Hint / Algorithm:
1. Use a dictionary or object to count occurrences.
2. Iterate again to find the first element with count 1.
3. Return that element or -1 if all are repeating.
*/

function firstNonRepeating(arr) {
    let count = {};
    for (let num of arr) {
        count[num] = (count[num] || 0) + 1;
    }
    for (let num of arr) {
        if (count[num] === 1) return num;
    }
    return -1;
}

// Test Cases
console.log(firstNonRepeating([4, 5, 1, 2, 0, 4, 5, 1])); // 2
console.log(firstNonRepeating([7, 8, 7, 8, 9])); // 9
console.log(firstNonRepeating([3, 3, 3])); // -1
console.log(firstNonRepeating([])); // -1
