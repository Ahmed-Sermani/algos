


def partition(node, x):
    
    before_start = before_end = after_start = after_end = None

    while node is not None:
        next = node.next
        node.next = None
        if node.val < x:
            if before_start is None:
                before_start = node
                before_end = before_start
            else:
                before_end.next = node
                before_end = node
        else:
            if after_start is None:
                after_start = node
                after_end = after_start
            else:
                after_end.next = node
                after_end = node

        node = next

    if before_start is None:
        return after_start

    # merge
    before_end.next = after_start

    return before_start



def partition2(node, x):

    if node is None:
        return

    head = tail = node

    while node is not None:
        next = node.next
        if node.val < x:
            node.next = head
            head = node
        else:
            tail.next = node 
            tail = node

        node = next

    tail.next = None

    return head


