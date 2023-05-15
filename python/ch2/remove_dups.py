
# Time: O(N)
# Space: O(N)
def remove_dups(head):
    if not head: return
    seen = set()
    prev = None 
    node = head
    while node is not None:
        if node.val in seen:
            prev.next = node.next
        else:
            seen.add(node.val)
        prev = node
        node = node.next
    return head

# Time: O(N^2)
# Space: O(1)
def remove_dups_2(head):
    cur = head

    while cur is not None:
        runner = cur.next 
        prev = cur
        while runner is not None:
            if runner.val == cur.val:
                prev.next = runner.next
            prev = runner
            runner = runner.next
        cur = cur.next
    return head
        

