
# Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the 
# beginning of the loop. 
# DEFINITION 
# Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so 
# as to make a loop in the linked list.


# Time: O(N)
# Space: O(1)

def loop_detection(l):
    # The solution is to use Floyd’s Cycle Finding Algorithm
    if l is None or l.next is None:
        return None

    fast = slow = l

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            break

    # No cycle
    if fast is None or fast.next is None:
        return None

    slow = l

    while True:
        if slow is fast:
            return slow
        slow = slow.next
        fast = fast.next

