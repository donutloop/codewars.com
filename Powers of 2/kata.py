from math import pow

def powers_of_two(n):
    nums = []
    for exp in range(0, n+1):
        nums.append(int(pow(2, exp)))
    return nums