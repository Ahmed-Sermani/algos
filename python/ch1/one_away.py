

# Time: O(N) where N the length of the shorter string
# Space: O(1)
def one_away(s1, s2):
    if len(s1) == len(s2):
        return check_replacement(s1, s2)
    elif len(s1) + 1 == len(s2):
        return check_insert(s1, s2)
    elif len(s1) - 1 == len(s2):
        return check_insert(s2, s1)
    
    return False

def check_replacement(s1, s2):
    found_diff = False
    for i in range(len(s1)):
        if s1[i] != s2[i]:
            if found_diff:
                return False
            found_diff = True
    return True

def check_insert(s1, s2):
    i1 = 0
    i2 = 0
    while i1 < len(s1) and i2 < len(s2):
        if s1[i1] != s2[i2]:
            if i1 != i2:
                return False
            i2 += 1
        else:
            i1 += 1
            i2 += 1

    return True
            
