
# Time: O(N^2)
# Space: O(1)
def is_unique_without_additional_datastructure(s):
    for i in range(len(s)):
        for k in range(i+1, len(s)):
            if s[i] == s[k]:
                return False
    return True

# TIME: O(N)
# Space: O(N)
def is_unique_with_hashmap(s):
    unique = set()
    for c in s:
        unique.add(c)
    if len(unique) != len(s):
        return False
    return True


# Time: O(N)
# Space: O(1)
def is_unique_with_bitvector(s):
    # assumes all ascci charactors
    vec = 0
    for c in s:
        if (vec >> ord(c) & 1) > 0:
            return False
        vec |= 1 << ord(c)

    return True


