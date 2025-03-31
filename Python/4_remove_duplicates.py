"""
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
1. Create an empty result list.
2. Iterate through the input list.
3. Add each element to the result only if it is not already present.
4. Avoid using `set()` to practice logic-building.
"""

def remove_duplicates(arr):


# Test Cases
print(remove_duplicates([1, 2, 3, 1, 2, 4, 5]))  # [1, 2, 3, 4, 5]
print(remove_duplicates([7, 8, 9, 7, 7, 8]))  # [7, 8, 9]
print(remove_duplicates([]))  # []
