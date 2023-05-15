from collections import Counter

# Time: O(N)
# Space: O(N)
def palindrome_permutation(s):
    letters = Counter(s.replace(" ", "").lower())
    found_one_odd = False
    for count in letters.values():
        if found_one_odd and count | 0 == 1:
            return False
        if count & 1 == 1:
            found_one_odd = True
    
    return True

# Time: O(N)
# Space: O(1)
def palindrome_permutation_v2(s):
    bitvector = 0
    for c in s.replace(" ", ""):
        index = ord(c.lower())
        mask = 1 << index
        bitvector ^= mask
        print(bin(bitvector))
    return bitvector == 0 or bitvector & (bitvector - 1) == 0

