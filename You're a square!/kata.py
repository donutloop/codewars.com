from math import sqrt

def is_square(n):
    if n < 0:
        return False
    if 0 == n:
        return True
    root = sqrt(n)
    return root == int(root)