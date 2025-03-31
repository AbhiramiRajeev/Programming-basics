"""
Problem 2: Find the Maximum Element in an Array

Given an array of numbers, find the maximum value in the array.

Example Input/Output:
Input: [1, 5, 3, 9, 2]  
Output: 9  

Input: [-5, -1, -8, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize `max` as the first element.
2. Iterate through the array.
3. If an element is greater than `max`, update `max`.
4. Return the `max` value.
5. Avoid using `max()` to build logic.
"""

def find_max(arr):

# Test Cases
print(find_max([1, 5, 3, 9, 2]))  # 9
print(find_max([-5, -1, -8, -3]))  # -1
print(find_max([10]))  # 10
print(find_max([]))  # None
