# recursion examples
def print_recursive(count):
    if count < 1:
        return
    print("Hello world")
    print_recursive(count - 1)


def print_array_elements(arr):
    if len(arr) == 0:
        return
    print(arr[0])
    # arr[start:] items from the beginning through the end
    arr = arr[1:]
    print_array_elements(arr)


def print_array_elements_reverse(arr):
    end_index = len(arr) - 1
    if end_index < 0:
        return
    print(arr[end_index])
    # arr[:stop] items from the beginning through stop-1
    arr = arr[:end_index]
    print_array_elements_reverse(arr)


def factorial(num):
    if num == 0 or num == 1:
        return 1
    # n!=n*(n-1)*(n-2)*(n-3)...*3*2*1
    # 5!=5*4*3*2*1=120
    return num * factorial(num - 1)


# print_recursive(10)
# print_array_elements([1, 2, 3])
# print_array_elements_reverse([1, 2, 3])
print(factorial(5))
