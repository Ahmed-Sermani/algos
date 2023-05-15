
from ch2 import Node


def sum_lists(n1, n2):
    sum_list = curr = None
    carry = 0

    while n1 or n2:
        res = None
        if n1 and n2:
            sum = n1.val + n2.val + carry
            carry = 1 if sum >= 10 else 0 
            res = Node(sum % 10)
        else:
            sum = n1.val + carry if n1 else n2.val + carry
            carry = 1 if sum >= 10 else 0
            res = Node(sum % 10)

        if sum_list is None:
            sum_list = curr = res
        else:
            curr.next = res
            curr = curr.next
        
        if n1:
            n1 = n1.next
        if n2:
            n2 = n2.next

    if carry != 0:
        curr.next = Node(carry)
    
    return sum_list

def sum_list_recur(n1, n2):
    def recur(n1, n2, carry):
        if n1 is None and n2 is None and carry == 0:
            return None

        value = carry
        if n1:
            value += n1.val

        if n2:
            value += n2.val

        res = Node(value % 10)
        res.next = recur(n1.next if n1 else None, n2.next if n2 else None, 1 if value >= 10 else 0)
        return res

    return recur(n1, n2, 0)

def sum_list_forward(n1, n2):
    number1 = []
    while n1:
        number1.append(str(n1.val))
        n1 = n1.next
    number1 = ''.join(number1)
    
    number2 = []
    while n2:
        number2.append(str(n2.val))
        n2 = n2.next
    number2 = ''.join(number2)

    res = int(number1) + int(number2)

    res_root = curr = None
    for c in str(res):
        if res_root is None:
            res_root = curr = Node(int(c))
        else:
            curr.next = Node(int(c))
            curr = curr.next

    return res_root

def sum_list_forward_recur(n1, n2):

    # TODO: handle padding when the two lists not equal in length

    def recur(n1, n2):
        if n1 is None and n2 is None:
            return None, 0

        more, carry = recur(n1.next if n1 else None, n2.next if n2 else None)

        value = carry

        if n1:
            value += n1.val

        if n2:
            value += n2.val

        res = Node(value % 10)
        res.next = more

        return res, 1 if value >= 10 else 0
    
    res, _ = recur(n1, n2)

    return res


