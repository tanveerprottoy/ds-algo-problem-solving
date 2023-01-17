from typing import List


# Given an array of integers, return indices of the two numbers such that they add up to a specific target.
# You may assume that each input would have exactly one solution, and you may not use the same element twice.
# Example:
# Given nums = [2, 7, 11, 15], target = 9,
# Because nums[0] + nums[1] = 2 + 7 = 9,
# return [0, 1].
def two_sum(nums: List[int], target: int) -> List[int]:
    index_map = dict()
    index_list = list()
    for i in range(len(nums)):
        value = nums[i]
        required = target - value
        if len(index_map) > 0 and required in index_map:
            previous_index = index_map.get(required)
            index_list.append(previous_index)
            index_list.append(i)
        index_map[value] = i
    return index_list


# You are given two non-empty linked lists representing two non-negative integers.
# The digits are stored in reverse order and each of their nodes contain a single digit.
# Add the two numbers and return it as a linked list.
# You may assume the two numbers do not contain any leading zero, except the number 0 itself.
# Example:
# Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
# Output: 7 -> 0 -> 8
# Explanation: 342 + 465 = 807.
class ListNode:
    def __init__(self, x, next_node = None):
        self.val = x
        self.next = next_node

    def __repr__(self):
        return "<ListNode val:%s next:%s>" % (self.val, self.next)

    def __str__(self):
        return "val is %s, next is %s" % (self.val, self.next)


def add_two_numbers(l1: ListNode, l2: ListNode) -> ListNode:
    temp0 = l1
    temp1 = l2
    result = None
    carry = -1
    while temp0 is not None or temp1 is not None or carry >= 0:
        sum_val = 0 if carry < 0 else carry
        carry = -1
        if temp0 is not None:
            sum_val += temp0.val
            temp0 = temp0.next
        if temp1 is not None:
            sum_val += temp1.val
            temp1 = temp1.next
        if sum_val >= 10:
            carry = sum_val // 10
            sum_val = sum_val % 10
        curr = ListNode(sum_val)
        if result is None:
            result = curr
        else:
            temp2 = result
            while temp2.next is not None:
                temp2 = temp2.next
            temp2.next = curr
    return result


def add_two_numbers_recursive(l1: ListNode, l2: ListNode, carry: int) -> ListNode:
    if l1 is None:
        return l2
    if l2 is None:
        return l1
    sum_val = l1.val + l2.val + carry
    carry = 0
    if sum_val >= 10:
        carry = sum_val // 10
        sum_val %= 10
        list_node = ListNode(sum_val)
        list_node.next = add_two_numbers_recursive(l1.next, l2.next, carry)
        return list_node
    else:
        list_node = ListNode(sum_val)
        list_node.next = add_two_numbers_recursive(l1.next, l2.next, carry)
        return list_node


def longest_subsequence(s: str):
    substring = {}
    for i in range(len(s)):
        substring[s[i]] = i
    print(substring)


# Given a string, find the length of the longest substring without repeating characters.
# Example 1:
# Input: "abcabcbb"
# Output: 3
# Explanation: The answer is "abc", with the length of 3.
def length_of_longest_substring(string: str) -> int:
    substring = ''
    substring_temp = ''
    length_str = len(string)
    for i in range(length_str):
        value = string[i]
        if substring_temp != '':
            remainder_count = length_str - 1 - i
            if substring_temp.__contains__(value):
                # repeating value found
                # check if excluding previous one creates equal substring
                # if so reassign substring to the excluded one
                matched_index = substring_temp.index(value)
                temp_sub = substring_temp[matched_index + 1:]
                if len(temp_sub) + 1 == len(substring_temp):
                    substring_temp = temp_sub
                    substring_temp += value
                elif len(temp_sub) + 1 + remainder_count >= len(substring):
                    # a longer or equal substring can be found
                    substring_temp = temp_sub
                    substring_temp += value
                else:
                    # longest substring already found
                    # break the loop
                    break
            else:
                substring_temp += value
        else:
            substring_temp = value
        if len(substring_temp) >= len(substring):
            substring = substring_temp
    print(substring)
    return len(substring)


# Given a string, find the length of the longest substring without repeating characters.
# Example 1:
# Input: "abcabcbb"
# Output: 3
# Explanation: The answer is "abc", with the length of 3.
# implementation using sliding window technique
# whenever repeating char found, previous values are removed from set
def length_of_longest_substring_set(string: str) -> int:
    length_arr = len(string)
    fast = slow = 0
    result_set = set()
    result = 0
    while fast < length_arr and slow < length_arr:
        value = string[fast]
        # try to extend the range[i, j]
        if value not in result_set:
            # this portion extends the fast range
            result_set.add(value)
            fast += 1
            result = max(result, fast - slow)
        else:
            # this portion extends the slow range
            result_set.remove(string[slow])
            slow += 1
    return result


def length_of_longest_substring_map(string: str) -> int:
    slow = 0
    result_dict = dict()
    result = 0
    for fast in range(len(string)):
        value = string[fast]
        if value in result_dict:
            # invalidate previous substring, slow points to the last matched value
            slow = max(result_dict[value], slow)
        # + 1 is added because length is being calculated, max_index + 1 is the length
        result = max(result, fast - slow + 1)
        result_dict[value] = fast + 1
    return result


def length_of_longest_substring_map_str(string: str) -> int:
    length_arr = len(string)
    slow = 0
    result_dict = dict()
    result = 0
    result_string = temp_str = ''
    for fast in range(length_arr):
        value = string[fast]
        if value in result_dict:
            slow = max(result_dict[value], slow)
            temp_str = value
        if value not in result_dict:
            temp_str += value
        # + 1 is added because length is being calculated, max_index + 1 is the length
        result = max(result, fast - slow + 1)
        result_dict[value] = fast + 1
        if len(temp_str) > len(result_string):
            result_string = temp_str
    print(result_string)
    return result


def length_of_longest_substring_set_str(string: str) -> int:
    length_arr = len(string)
    fast = slow = 0
    result_set = set()
    result = 0
    result_string = temp_str = ''
    val: str
    while fast < length_arr and slow < length_arr:
        val = string[fast]
        # try to extend the range[i, j]
        if val not in result_set:
            # this portion extends the fast range
            result_set.add(val)
            fast += 1
            result = max(result, fast - slow)
            # for printing result string
            temp_str += val
        else:
            # this portion extends the slow range
            val = string[slow]
            result_set.remove(val)
            slow += 1
            temp_str = temp_str.replace(val, '')
        if len(temp_str) > len(result_string):
            result_string = temp_str
    print(result_set)
    print(result_string)
    return result


def length_of_longest_palindromic_substring(string: str) -> int:
    return -1
    # length_arr = len(string)
    # fast = slow = 0
    # result_set = set()
    # result = 0
    # while fast < length_arr and slow < length_arr:
    #     value = string[fast]
    #     # try to extend the range[i, j]
    #     if value not in result_set:
    #         # this portion extends the fast range
    #         result_set.add(value)
    #         fast += 1
    #         result = max(result, fast - slow)
    #     else:
    #         # this portion extends the slow range
    #         result_set.remove(string[slow])
    #         slow += 1
    # return result


# There are two sorted arrays nums1 and nums2 of size m and n respectively.
# Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
# You may assume nums1 and nums2 cannot be both empty.
# Example 1:
# nums1 = [1, 3]
# nums2 = [2]
# The median is 2.0
# shortcut inefficient way
def median_sorted_arrays(nums1: List[int], nums2: List[int]) -> float:
    if nums1 is None and nums2 is None:
        return 0.0
    merged_list = nums1 + nums2
    merged_list.sort()
    length_merged_list = len(merged_list)
    middle_point: int
    median: float
    if length_merged_list % 2 == 0:
        middle_point = int((length_merged_list + 1) / 2)
        # middle point found but index will but middle_point - 1, next will be middle_point
        median = (merged_list[middle_point - 1] + merged_list[middle_point]) / 2
    else:
        middle_point = int((length_merged_list + 1) / 2)
        median = merged_list[middle_point - 1]
    return median


# Given a string s, return the longest palindromic substring in s.
# A string is called a palindrome string
# if the reverse of that string is the same as the original string.
def longest_palindrome(s: str) -> str:
    if len(s) == 0:
        return ""
    output = ""
    temp = ""
    start = 0
    full_length = len(s)
    end = full_length - 1
    while start < full_length:
        if s[start] == s[end]:
            temp = temp + s[start]
        else:
            if len(temp) > len(output):
                output = temp
        start += 1
        end -= 1
    if len(temp) > len(output):
        output = temp
    if len(output) == 0:
        return s[0]
    return output


def longest_palindrome_stack(s):
    if len(s) == 0:
        return ""
    stack = []
    output = ""
    start = 0
    full_length = len(s)
    while start < full_length:
        val = s[start]
        if len(stack) > 0:
            output = output + stack.pop()
            if not output:
                stack.append(val)
    return output


# print(two_sum([2, 7, 11, 15], 9))
# print(

#     add_two_numbers(
#         ListNode(2, ListNode(4, ListNode(3))),
#         ListNode(5, ListNode(6, ListNode(4)))
#     )
# )
# print(
#     add_two_numbers_recursive(
#         ListNode(2, ListNode(4, ListNode(3))),
#         ListNode(5, ListNode(6, ListNode(4))),
#         0
#     )
# )
# longest_subsequence("pwwkew")
# length_of_longest_substring("anviaj")
# length_of_longest_substring("pwwkew")
# length_of_longest_substring("ruowzgiooobpple")
# length_of_longest_substring("hkcpmprxxxqw")
# length_of_longest_substring("hkcpmprxxxqw")
# print(length_of_longest_substring_set("pwwkew"))
print(length_of_longest_substring_map_str("pwwkew"))
print(length_of_longest_substring_map_str(" "))
# print(median_sorted_arrays([1, 3], [2]))

# print(longest_palindrome("babad"))
# print(longest_palindrome("cbbd"))
# print(longest_palindrome("ababa"))
# print(longest_palindrome("ab"))
# print(longest_palindrome("abba"))
print(longest_palindrome_stack("abb"))
