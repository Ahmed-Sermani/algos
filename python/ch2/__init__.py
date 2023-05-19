
class Node:
    next: 'Node'
    def __init__(self, val=None):
        self.val = val
        self.next = None
        self.prev = None


    def __repr__(self) -> str:
        return f"Node({self.val})"
