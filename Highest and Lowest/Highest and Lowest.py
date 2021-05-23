def high_and_low(numbers):
    numbers = numbers.split(" ")

    max = None
    min = None
    for num in numbers:
        num = int(num)
        if max == None or max < num:
            max = num
        if min == None or min > num:
            min = num

    return str(max) + " " + str(min)