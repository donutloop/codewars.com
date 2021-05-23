def array_diff(a, b):
    c = []
    b = set(b)
    for num in a:
        if num in b:
            continue
        c.append(num)
    return c