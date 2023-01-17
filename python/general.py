# dividing in integers by ten gets rid of the last digit, while obtaining modulo ten
# gives the last digit
# this function returns the digits in reverse order
from typing import List


def extract_digits_mathematical(num: int):
    digits = list()
    while num >= 10:
        digits.append(num % 10)
        num //= 10
    # last element will be in num
    digits.append(num)
    print(digits)


def extract_digits(num: int, digits):
    if num < 10:
        digits.append(num)
    else:
        # double slash is for int division
        # recursion for extracting next digit
        extract_digits(num // 10, digits)
        digits.append(num % 10)


def factorial(num):
    if num == 0 or num == 1:
        return 1
    # n!=n*(n-1)*(n-2)*(n-3)...*3*2*1
    # 5!=5*4*3*2*1=120
    return num * factorial(num - 1)


def find_median(nums: List[int]):
    # to find middle point count the total element in a set and add 1
    # and divide by 2
    # example: count=7, 7+1=8/2=4, 4 is middle point here
    # for odd number of total element count, the value in middle point is the median
    # for even number of total element count, the middle point will be fractional num,
    # the median will be fractional number's whole number + next number
    # example: count=8, 8+1=9/2=4.5, sum of position 4 and 5 numbers is the median
    # sort the number list first
    # not efficient solution
    nums.sort()
    length_list = len(nums)
    middle_point: int
    median: float
    if length_list % 2 == 0:
        middle_point = int((length_list + 1) / 2)
        # middle point found but index will but middle_point - 1, next will be middle_point
        median = (nums[middle_point - 1] + nums[middle_point]) / 2
    else:
        middle_point = int((length_list + 1) / 2)
        median = nums[middle_point - 1]
    return median


def is_palindrome_arr(arr) -> bool:
    start = 0
    length = len(arr)
    end = length - 1
    while start < end and end >= 0:
        if arr[start] != arr[end]:
            return False
        start += 1
        end -= 1
    return True


def is_palindrome_str(string) -> bool:
    start = 0
    length = len(string)
    end = length - 1
    while start < end and end >= 0:
        if string[start] != string[end]:
            return False
        start += 1
        end -= 1
    return True


def is_palindrome_split(s) -> bool:
    start = 0
    full_length = len(s)
    half_length = len(s) // 2
    end = full_length - 1
    while start < half_length:
        if s[start] != s[end]:
            return False
        start += 1
        end -= 1
    return True


# extract_digits_mathematical(10)
# digit_list = list()
# extract_digits(10, digit_list)
# print(digit_list)
# print(factorial(8))
# print(find_median([1, 2, 3, 4]))
print(is_palindrome_arr(["n", "o", "o", "n"]))
print(is_palindrome_split("civic"))
print(is_palindrome_split("abbc"))
