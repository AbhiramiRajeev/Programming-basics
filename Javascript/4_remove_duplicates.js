/*
Problem 4: Remove Duplicates from an Array

Given an array, remove duplicate values and return the unique elements while maintaining the order.

Example Input/Output:
Input: [1, 2, 3, 1, 2, 4, 5]  
Output: [1, 2, 3, 4, 5]  

Input: [7, 8, 9, 7, 7, 8]  
Output: [7, 8, 9]  

Input: []  
Output: []  

Hint / Algorithm:
1. Create an empty result array.
2. Iterate through the input array.
3. Add each element to the result only if it is not already present.
4. Avoid using `Set()` to practice logic-building.
*/

function removeDuplicates(arr) {
    let unique = [];
    for (let num of arr) {
        if (!unique.includes(num)) {
            unique.push(num);
        }
    }
    return unique;
}

// Test Cases
console.log(removeDuplicates([1, 2, 3, 1, 2, 4, 5])); // [1, 2, 3, 4, 5]
console.log(removeDuplicates([7, 8, 9, 7, 7, 8])); // [7, 8, 9]
console.log(removeDuplicates([])); // []
