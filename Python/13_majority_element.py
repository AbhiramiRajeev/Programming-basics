"""
Problem 13: Find the Majority Element

Given an array of size n, find the majority element. The majority element is 
the element that appears more than ⌊n / 2⌋ times. You may assume that the array 
is non-empty and that the majority element always exists.

Example Input/Output:
Input: [3, 2, 3]
Output: 3

Input: [2, 2, 1, 1, 1, 2, 2]
Output: 2

Hint / Algorithm:
1. Use a dictionary to count occurrences and find the majority element.
2. Alternatively, use the Boyer-Moore Voting Algorithm.
"""

def majority_element(nums):
    # Implement the function here
    pass

# Test Cases
print(majority_element([3, 2, 3])) 
print(majority_element([2, 2, 1, 1, 1, 2, 2])) 
print(majority_element([1, 1, 1, 2, 3, 1, 1])) 
print(majority_element([4])) 
print(majority_element([5, 5, 5, 5, 5, 5, 5])) 
