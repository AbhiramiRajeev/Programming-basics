"""
Problem 19: Longest Consecutive Sequence

Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Example Input/Output:
Input: [100, 4, 200, 1, 3, 2]
Output: 4 (because [1, 2, 3, 4] is the longest sequence)

Input: [9,1,4,7,3,-1,0,5,8,-3,6,2]
Output: 7 ([-1,0,1,2,3,4,5,6])

Hint / Algorithm:
1. Store all elements in a set for quick lookup.
2. Iterate through the array and find the start of a sequence.
3. Expand the sequence by checking consecutive numbers in the set.
"""

def longest_consecutive(nums):
    # Implement the function here
    pass

# Test Cases
print(longest_consecutive([100, 4, 200, 1, 3, 2]))  # Output: 4
print(longest_consecutive([9,1,4,7,3,-1,0,5,8,-3,6,2]))  # Output: 7
print(longest_consecutive([1,2,0,1]))  # Output: 3
print(longest_consecutive([]))  # Output: 0
print(longest_consecutive([10,20,30,40]))  # Output: 1
