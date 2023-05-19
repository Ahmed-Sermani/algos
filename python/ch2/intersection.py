

# Intersection: 
# Given two (singly) linked lists, determine if the two lists intersect. 
# Return the intersecting node. Note that the intersection is defined based on reference, not value.
# That is, if the kth node of the first linked list is the exact same node (by reference) as the jth 
# node of the second linked list, then they are intersecting.

# Time: O(N+M)
# Space: O(1)
def intersection(l1, l2):
    
    if l1 is None or l2 is None:
        return None
    # two intersecting lists will have the same ending.
    c1 = c2 = 1
    t1 = l1
    while t1.next:
        t1 = t1.next
        c1 += 1

    t2 = l2
    while t2.next:
        t2 = t2.next
        c2 += 1

    if not (t1 is t2):
        return None

    l1_skip = 0 if c1 <= c2 else c1 - c2
    l2_skip = 0 if c2 <= c1 else c2 - c1

    l1_curr = l1
    l2_curr = l2

    while l1_curr and l2_curr:
        if l1_skip:
            l1_curr = l1_curr.next
            l1_skip -= 1
            continue

        if l2_skip:
            l2_curr = l2_curr.next
            l2_skip -= 1
            continue
        
        if l1_curr is l2_curr:
            return l1_curr

        l1_curr = l1_curr.next
        l2_curr = l2_curr.next

    return None
