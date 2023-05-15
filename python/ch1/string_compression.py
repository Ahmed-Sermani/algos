# Time: O(N)
# Space: O(N)
def string_compression(s):
    buf = list()
    c = 1
    i = 0
    for i in range(1, len(s)):
        if s[i] == s[i - 1]:
            c += 1
        else:
            buf.append(s[i - 1])
            buf.append(str(c))
            c = 1
    buf.append(s[i])
    buf.append(str(c))
    return ''.join(buf) if len(buf) < len(s) else s

