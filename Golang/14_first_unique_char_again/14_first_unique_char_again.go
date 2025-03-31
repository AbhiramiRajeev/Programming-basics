package main

import "fmt"

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

func firstNonRepeatingChar(s string) rune {
    // Implement the function here
    return '_'
}

func main() {
    fmt.Println(string(firstNonRepeatingChar("aabccdeff"))) // Output: "b"
    fmt.Println(string(firstNonRepeatingChar("aabbcc")))    // Output: "_"
    fmt.Println(string(firstNonRepeatingChar("abcd")))      // Output: "a"
    fmt.Println(string(firstNonRepeatingChar("aabbccddeeffg"))) // Output: "g"
    fmt.Println(string(firstNonRepeatingChar("xxyyzz")))    // Output: "_"
}
