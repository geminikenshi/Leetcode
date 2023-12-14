# 20. Valid Parentheses

LeetCode : [problem](https://leetcode.com/problems/valid-parentheses/description/)

This problem is easy to solve with the stack data structure

Every closing bracket has to match the closest opening bracket or it is invalid.

Example:

- `"([)]"`
  This is invalid
- `"((((()[]{{}}))))"`
  This is valid

## Version 2

Use an object as a hashmap, resulting in much cleaner code.
