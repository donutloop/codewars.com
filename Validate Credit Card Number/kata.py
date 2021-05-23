def validate(n):
    n = str(n)
    n = list(map(int, n))
    k = len(n)
    if k % 2 == 0:
        for idx, val in enumerate(n):
            if idx % 2 == 0:
                n[idx] = n[idx] * 2
    else:
        for idx, val in enumerate(n):
            if idx % 2 == 1:
                n[idx] = n[idx] * 2


    sum = 0
    for val in n:
        if val > 9:
            val = val - 9
        sum += val
    return sum % 10 == 0