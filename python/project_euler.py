# If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these
# multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.


def multiples_of(max_value):
    multiples_list = list()
    sum_all_multiples = 0
    for i in range(1, max_value):
        if i % 3 == 0 or i % 5 == 0:
            multiples_list.append(i)
            sum_all_multiples += i
    print('Multiples: ' + multiples_list.__str__())
    print('Sum of all the multiples: ' + sum_all_multiples.__str__())


# Each new term in the Fibonacci sequence is generated by adding the previous two terms
# the first 10 terms will be:
#
# 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

def fibonacci_sequence(max_value):
    if max_value <= 2:
        return
    sequence_list = list()
    last_value = 0
    current_value = 1
    sequence_list.append(last_value)
    sequence_list.append(current_value)
    next_value = last_value + current_value
    while next_value <= max_value:
        sequence_list.append(next_value)
        last_value = current_value
        current_value = next_value
        next_value = last_value + current_value
    print('Fibonacci sequence ' + sequence_list.__str__())


def fibonacci_sequence_recursive(num) -> int:
    if num <= 1:
        return num
    return fibonacci_sequence_recursive(num - 1) + fibonacci_sequence_recursive(num - 2)


# By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the
# even-valued terms.

def even_fibonacci_terms(max_value):
    if max_value <= 2:
        return
    even_term_list = list()
    even_terms_sum = 0
    last_value = 0
    current_value = 1
    next_value = last_value + current_value
    while next_value <= max_value:
        if next_value % 2 == 0:
            even_term_list.append(next_value)
            even_terms_sum += next_value
        last_value = current_value
        current_value = next_value
        next_value = last_value + current_value
    print('Even Fibonacci terms ' + even_term_list.__str__())
    print('Even Fibonacci terms sum ' + even_terms_sum.__str__())


def prime_number(number):
    # 1 is not prime nor composite number
    if number <= 1:
        return False
    for i in range(2, number):
        if number % i == 0:
            return False
    return True


def prime_numbers(max_value):
    prime_number_list = list()
    for i in range(2, max_value + 1):
        if prime_number(i):
            prime_number_list.append(i)
    return prime_number_list


def factorials(number):
    factorial_list = list()
    for i in range(1, number + 1):
        if number % i == 0:
            factorial_list.append(i)
    return factorial_list


def prime_factorials_partial(number):
    prime_factorial_partial_list = list()
    for i in range(1, number + 1):
        if number % i == 0:
            if prime_number(i):
                prime_factorial_partial_list.append(i)
    return prime_factorial_partial_list


def prime_factorials(number):
    pass


# multiples_of(1000)
# fibonacci_sequence(100)
# even_fibonacci_terms(100)
print(fibonacci_sequence_recursive(5))
# print('Is prime number: ' + str(prime_number(2)))
# print('Prime numbers: ' + str(prime_numbers(100)))
# print('Factorials: ' + str(factorials(100)))
# print('Prime factorials partial: ' + str(prime_factorials_partial(100)))