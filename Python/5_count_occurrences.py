"""
Problem 5: Count Occurrences of an Element

Given an array and a target value, count how many times the target appears in the array.

Example Input/Output:
Input: ([1, 2, 3, 2, 2, 4, 5], 2)  
Output: 3  

Input: ([7, 8, 7, 7, 9, 10], 7)  
Output: 3  

Input: ([], 5)  
Output: 0  

Hint / Algorithm:
1. Initialize a count variable to 0.
2. Iterate through the array.
3. If an element matches the target, increase the count.
4. Return the count.
"""

def count_occurrences(arr, target):


# Test Cases
print(count_occurrences([1, 2, 3, 2, 2, 4, 5], 2))  # 3
print(count_occurrences([7, 8, 7, 7, 9, 10], 7))  # 3
print(count_occurrences([], 5))  # 0
