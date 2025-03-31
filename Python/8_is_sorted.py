"""
Problem 8: Check if an Array is Sorted

Given an array of numbers, determine if it is sorted in non-decreasing order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: true  

Input: [3, 2, 1]  
Output: false  

Input: [7, 7, 8, 9]  
Output: true  

Hint / Algorithm:
1. Iterate through the array.
2. Compare each element with the next one.
3. If any element is greater than the next, return False.
4. If loop completes, return True.
"""

def is_sorted(arr):


# Test Cases
print(is_sorted([1, 2, 3, 4, 5]))  # True
print(is_sorted([3, 2, 1]))  # False
print(is_sorted([7, 7, 8, 9]))  # True
print(is_sorted([]))  # True
print(is_sorted([10]))  # True
