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


fizz_buzz(1, 100)
