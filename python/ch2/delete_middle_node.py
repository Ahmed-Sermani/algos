

def delete_middle_node(node):
    node.val = node.next.val
    node.next = node.next.next
        

