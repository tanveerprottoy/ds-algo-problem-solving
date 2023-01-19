def fizz_buzz(start, end):
    for i in range(start, end + 1):
        divisible_by_3 = i % 3 == 0
        divisible_by_5 = i % 5 == 0
        if divisible_by_3 and divisible_by_5:
            print("FizzBuzz")
        elif divisible_by_3:
            print("Fizz")
        elif divisible_by_5:
            print("Buzz")
        else:
            print(i)


def reverse_string(string):
    reversed = ""
    # for i = 0; i >= 0 in range(len(string)):
    #     divisible_by_3 = i % 3 == 0
    #     divisible_by_5 = i % 5 == 0
    #     if divisible_by_3 and divisible_by_5:
    #         print("FizzBuzz")
    #     elif divisible_by_3:
    #         print("Fizz")
    #     elif divisible_by_5:
    #         print("Buzz")
    #     else:
    #         print(i)

def largest_smallest_number_in_array(arr, arrSize):
        result = [-1, -1]
        # Corner case
        if arrSize <= 0:
            return result
        for val in arr:
            # check large
            if result[0] < val:
                result[0] = val
            if result[1] == -1 or result[1] > val:
                result[1] = val
        return result


fizz_buzz(1, 100)

arr = [20, 5, 4, 6, 9]

print(largest_smallest_number_in_array(arr, len(arr)))
