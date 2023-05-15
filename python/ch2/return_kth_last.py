
# Time: O(N)
# Space: O(1)
def return_kth_last(head, k):
    node = head
    c = 1
    while node is not None:
        node = node.next
        c += 1
    position = c - k
    c = 1
    node = head
    while node is not None:
        if c == position:
            return node
        node = node.next
        c += 1
    return None

# Time: O(N)
# Space: O(N)
def return_kth_last_2(head, k):
    i = 0
    kth = None
    def _recur(head, k):
        nonlocal i, kth
        if head is None:
            return
        _recur(head.next, k)
        i += 1

        if i == k:
            kth = head
        return
    _recur(head, k) 

    return kth

def return_kth_last_3(head, k):
    node = kth = head
    c = 0
    while True:
        if node.next is None:
            return kth
        node = node.next
        c += 1
        if c >= k:
            kth = kth.next
