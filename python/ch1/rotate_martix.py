import math

def rotate_martic(m):
    n = len(m)

    for layer in range(n//2):
        first, last = layer, n - layer - 1
        for i in range(first, last):
            top = m[layer][i]
            m[layer][i] = m[n - i - 1][]
