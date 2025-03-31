/*
Problem 14: First Non-Repeating Character

Given a string, find the first non-repeating character and return it. 
If there is no unique character, return an underscore ('_').

Example Input/Output:
Input: "aabccdeff"
Output: "b"

Input: "aabbcc"
Output: "_"

Hint / Algorithm:
1. Use a hash map to count occurrences of each character.
2. Iterate through the string again to find the first character with a count of 1.
*/

function firstNonRepeatingChar(s) {
    // Implement the function here
}

// Test Cases
console.log(firstNonRepeatingChar("aabccdeff")); // Output: "b"
console.log(firstNonRepeatingChar("aabbcc")); // Output: "_"
console.log(firstNonRepeatingChar("abcd")); // Output: "a"
console.log(firstNonRepeatingChar("aabbccddeeffg")); // Output: "g"
console.log(firstNonRepeatingChar("xxyyzz")); // Output: "_"
