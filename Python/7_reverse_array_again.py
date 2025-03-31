"""
Problem 7: Reverse an Array

Given an array, return a new array that contains the elements in reverse order.

Example Input/Output:
Input: [1, 2, 3, 4, 5]  
Output: [5, 4, 3, 2, 1]  

Input: [7, 8, 9]  
Output: [9, 8, 7]  

Input: []  
Output: []  

Hint / Algorithm:
1. Create an empty result list.
2. Traverse the input list from the last element to the first.
3. Append each element to the result list.
"""

def reverse_array(arr):
    reversed_arr = []
    for i in range(len(arr) - 1, -1, -1):
        reversed_arr.append(arr[i])
    return reversed_arr

# Test Cases
print(reverse_array([1, 2, 3, 4, 5]))  # [5, 4, 3, 2, 1]
print(reverse_array([7, 8, 9]))  # [9, 8, 7]
print(reverse_array([]))  # []
