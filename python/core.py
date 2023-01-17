# Python Program to find the L.C.M. of two input number
# The Least Common Multiple
def compute_lcm(x, y):
    # choose the greater number
    if x > y:
        greater = x
    else:
        greater = y
    while True:
        if (greater % x == 0) and (greater % y == 0):
            lcm = greater
            break
        greater += 1
    return lcm


def multiples(value, start, stop):
    m = []
    for i in range(start, stop):
        m.append(value * i)
    return m


def factors(value, start, stop):
    m = []
    for i in range(start, stop):
        if i % value == 0:
            m.append(i)
    return m

# print(compute_lcm(3, 2))
