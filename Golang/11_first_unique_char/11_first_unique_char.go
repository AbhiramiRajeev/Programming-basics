package main

import "fmt"

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
1. Use a hash map to count occurrences of each character.
2. Iterate through the string and return the first character with a count of 1.
3. If none are found, return '_'.
*/

func firstUniqueChar(s string) string {

}

func main() {
	fmt.Println(firstUniqueChar("leetcode")) // "l"
	fmt.Println(firstUniqueChar("loveleetcode")) // "v"
	fmt.Println(firstUniqueChar("aabbcc")) // "_"
	fmt.Println(firstUniqueChar("swiss")) // "w"
	fmt.Println(firstUniqueChar("")) // "_"
}
