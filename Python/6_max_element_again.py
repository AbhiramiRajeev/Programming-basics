"""
Problem 6: Find the Maximum Element in an Array

Given an array of numbers, find and return the maximum element.

Example Input/Output:
Input: [3, 1, 9, 7, 4]  
Output: 9  

Input: [-5, -1, -7, -3]  
Output: -1  

Input: [10]  
Output: 10  

Hint / Algorithm:
1. Initialize max as the first element.
2. Iterate through the array, updating max whenever a larger element is found.
3. Return max after the loop.
4. Avoid using max() to practice logic-building.
"""

def find_max(arr):
    if not arr:
        return None  # Handle empty array case
    max_num = arr[0]
    for num in arr:
        if num > max_num:
            max_num = num
    return max_num

# Test Cases
print(find_max([3, 1, 9, 7, 4]))  # 9
print(find_max([-5, -1, -7, -3]))  # -1
print(find_max([10]))  # 10
print(find_max([]))  # None
