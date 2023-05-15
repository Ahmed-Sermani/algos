
# Time: O(N)
# Space: O(N)
def urlify(s):
    b = list(s)
    right = left = len(b) - 1
    saw_first_chr = False
    while left >= 0:
        if saw_first_chr:
            if b[left] == ' ':
                b[right] = '0'
                b[right - 1] = '2'
                b[right - 2] = '%'
                right -= 3
            else:
                b[right] = b[left]
                right -= 1
        left -= 1
        if left >= 0 and b[left] != ' ':
            saw_first_chr = True
    
    return ''.join(b)


