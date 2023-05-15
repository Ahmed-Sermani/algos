
from collections import defaultdict

from ch2 import Node


def palindrome(l):
    m = defaultdict(int)

    while l:
        m[l.val] += 1 
        l = l.next

    foundOdd = False

    for v in m.values():
        if v & 1 == 1:
            if foundOdd:
                return False
            foundOdd = True

    return True


def palindrome_2(l):
    rev = reverse_and_clone(l)

    while rev and l:
        if rev.val != l.val:
            return False
        rev = rev.next
        l = l.next

    return rev is None and l is None



def reverse_and_clone(l):
    head = None
    node = l

    while node is not None:
        nn = Node(node.val)
        nn.next = head
        head = nn
        node = node.next

    return head


def palindrome_3(l):
    fast = slow = l
    stack = []

    while fast and fast.next:
        stack.append(slow.val)
        slow = slow.next
        fast = fast.next.next

    if fast:
        slow = slow.next

    while slow:
        if stack.pop() != slow.val:
            return False
        slow = slow.next

    return True


def palindrome_4(l):
    length_ll = get_length(l)

    def recur(l, length):
        if l is None or length == 0:
            return l, True
        elif length == 1:
            return l.next, True

        sibling, is_palindrome = recur(l.next, length - 2)

        if not is_palindrome:
            return sibling, False

        return sibling.next, sibling.val == l.val

    _, is_palindrome = recur(l, length_ll)
    return is_palindrome



def get_length(l):
    c = 0 
    curr = l
    while curr:
        c += 1
        curr = curr.next
    return c

