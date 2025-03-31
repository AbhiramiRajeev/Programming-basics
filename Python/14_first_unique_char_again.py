"""
Problem 14: First Non-Repeating Character

Given a string, find the first non-repeating character and return it. 
If there is no unique character, return an underscore ('_').

Example Input/Output:
Input: "aabccdeff"
Output: "b"

Input: "aabbcc"
Output: "_"

Hint / Algorithm:
1. Use a dictionary to store character counts.
2. Iterate through the string again to find the first character with a count of 1.
"""

def first_non_repeating_char(s):
    # Implement the function here
    pass

# Test Cases
print(first_non_repeating_char("aabccdeff"))  # Output: "b"
print(first_non_repeating_char("aabbcc"))  # Output: "_"
print(first_non_repeating_char("abcd"))  # Output: "a"
print(first_non_repeating_char("aabbccddeeffg"))  # Output: "g"
print(first_non_repeating_char("xxyyzz"))  # Output: "_"
