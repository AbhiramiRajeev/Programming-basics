"""
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
1. Use a dictionary to count occurrences of each character.
2. Iterate through the string and return the first character with a count of 1.
3. If none are found, return '_'.
"""

def first_unique_char(s):


# Test Cases
print(first_unique_char("leetcode")) # "l"
print(first_unique_char("loveleetcode")) # "v"
print(first_unique_char("aabbcc")) # "_"
print(first_unique_char("swiss")) # "w"
print(first_unique_char("")) # "_"
