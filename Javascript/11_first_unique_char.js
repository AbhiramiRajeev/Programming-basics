/*
Problem 11: Find the First Non-Repeating Character

Given a string, find the first non-repeating character and return it. 
If there is no non-repeating character, return an underscore ('_').

Example Input/Output:
Input: "leetcode"
Output: "l"

Input: "loveleetcode"
Output: "v"

Input: "aabbcc"
Output: "_"

Hint / Algorithm:
1. Use a hash map (object in JS) to count occurrences of each character.
2. Iterate through the string and return the first character with a count of 1.
3. If none are found, return '_'.
*/

function firstUniqueChar(s) {
    let charCount = {};

    for (let char of s) {
        charCount[char] = (charCount[char] || 0) + 1;
    }

    for (let char of s) {
        if (charCount[char] === 1) {
            return char;
        }
    }

    return "_";
}

// Test Cases
console.log(firstUniqueChar("leetcode")); // "l"
console.log(firstUniqueChar("loveleetcode")); // "v"
console.log(firstUniqueChar("aabbcc")); // "_"
console.log(firstUniqueChar("swiss")); // "w"
console.log(firstUniqueChar("")); // "_"
