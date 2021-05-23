def divisors(integer):
    c = []
    for num in range(2, integer):
        if integer % num == 0:
            c.append(num)

    if c:
        return c

    return "{} is prime".format(integer)