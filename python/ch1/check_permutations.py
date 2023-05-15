
# Time: O(N^2)
# Space: O(1)
def check_permutation_brute_force(s1, s2):
    if len(s1) != len(s2):
        return False
    for c in s1:
        if c not in s2:
            return False
    return True

# Time: O(N * LogN)
# Space: O(1) assuming the sort done in place. which not possible in python
# since strings are imutable
def check_permutation_sort(s1, s2):
    if len(s1) != len(s2):
        return False
    s1 = sorted(s1)
    s2 = sorted(s2)
    for i in range(len(s1)):
        if s1[i] != s2[i]:
            return False
    return True


# Time: O(N)
# Space: O(N)
def check_permutation_hashmap(s1, s2):
    if len(s1) != len(s2):
        return False
    unique = set()
    for c in s1:
        unique.add(c)
    for c in s2:
        if c not in unique:
            return False
    return True

# assumes input all assci charactors
# Time: O(N)
# Space: O(1)
def check_permutation_vector(s1, s2):
    if len(s1) != len(s2):
        return False
    vec = [False] * 128 
    for c in s1:
        vec[ord(c)] = True

    for c in s2:
        if vec[ord(c)] is False:
            return False

    return True

# assumes input all assci charactors
# Time: O(N)
# Space: O(1)
def check_permutation_bitvector(s1, s2):
    if len(s1) != len(s2):
        return False
    vec = 0
    for c in s1:
        vec |= 1 << ord(c)

    for c in s2:
        if (vec >> ord(c) & 1) == 0:
            return False
    return True
