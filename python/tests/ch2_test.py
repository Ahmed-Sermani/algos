import pytest
from ch2 import Node
from ch2.palindrome import palindrome, palindrome_2, palindrome_3, palindrome_4
from ch2.remove_dups import remove_dups, remove_dups_2
from ch2.return_kth_last import return_kth_last, return_kth_last_2, return_kth_last_3
from ch2.delete_middle_node import delete_middle_node
from ch2.partition import partition
from ch2.sum_lists import sum_list_recur, sum_lists

def contruct_linked_list(*args):
    assert len(args) != 0
    head = Node(args[0])
    cur_node = head
    for arg in args[1:]:
        cur_node.next = Node(arg)
        cur_node = cur_node.next
    return head

def print_linked_list(head):
    while head is not None:
        print('-----------')
        print(head.val)
        head = head.next

@pytest.mark.parametrize(
    "l_in,l_out,func",
    [
        [contruct_linked_list(1, 2, 2, 3, 3, 4), contruct_linked_list(1, 2, 3, 4), remove_dups],
        [contruct_linked_list(1, 2, 2, 3, 3, 4, 4), contruct_linked_list(1, 2, 3, 4), remove_dups],
        [contruct_linked_list(1, 2, 2, 3, 3, 4), contruct_linked_list(1, 2, 3, 4), remove_dups_2],
        [contruct_linked_list(1, 2, 2, 3, 3, 4, 4), contruct_linked_list(1, 2, 3, 4), remove_dups_2],
    ]
)
def test_remove_dups(l_in, l_out, func):
    res = func(l_in)
    print_linked_list(res)
    while res is not None and l_out is not None:
        assert res.val == l_out.val
        res = res.next
        l_out = l_out.next
    assert res is None
    assert l_out is None

@pytest.mark.parametrize(
    "l_in,k,expect,func",
    [
        [contruct_linked_list(1, 2, 3, 4, 5), 2, 4, return_kth_last],
        [contruct_linked_list(1, 2), 1, 2, return_kth_last],
        [contruct_linked_list(1, 2, 3, 4, 5), 2, 4, return_kth_last_2],
        [contruct_linked_list(1, 2), 1, 2, return_kth_last_2],
        [contruct_linked_list(1, 2, 3, 4, 5), 2, 4, return_kth_last_3],
        [contruct_linked_list(1, 2), 1, 2, return_kth_last_3],
    ]
)
def test_return_kth_last(l_in, k, expect, func):
    if expect is None:
        assert func(l_in, k) is None
    assert func(l_in, k).val == expect

def create_linked_list(values):
    head = None
    tail = None
    for val in values:
        node = Node(val)
        if head is None:
            head = node
        else:
            tail.next = node
        tail = node
    return head

def linked_list_to_list(head):
    values = []
    while head is not None:
        values.append(head.val)
        head = head.next
    return values

def test_delete_middle_node():
    head = create_linked_list([1, 2, 3, 4, 5])
    node = head.next.next
    delete_middle_node(node)
    values = linked_list_to_list(head)
    assert values == [1, 2, 4, 5]

def test_partition_empty():
    assert partition(None, 5) is None

def test_partition_single():
    node = Node(3)
    assert partition(node, 5) == node

def test_partition_all_less():
    head = Node(1)
    head.next = Node(2)
    head.next.next = Node(3)
    assert partition(head, 4) == head

def test_partition_all_greater():
    head = Node(6)
    head.next = Node(7)
    head.next.next = Node(8)
    assert partition(head, 5) == head

def test_partition_mixed():
    head = Node(3)
    head.next = Node(5)
    head.next.next = Node(8)
    head.next.next.next = Node(5)
    head.next.next.next.next = Node(10)
    head.next.next.next.next.next = Node(2)
    head.next.next.next.next.next.next = Node(1)

    result = partition(head, 5)

    assert result.val == 3
    assert result.next.val == 2
    assert result.next.next.val == 1
    assert result.next.next.next.val == 5
    assert result.next.next.next.next.val == 8
    assert result.next.next.next.next.next.val == 5
    assert result.next.next.next.next.next.next.val == 10

def list_to_number(head):
    # Helper function to convert a linked list to a number
    num = 0
    factor = 1
    curr = head
    while curr:
        num += curr.val * factor
        factor *= 10
        curr = curr.next
    return num

def number_to_list(num):
    # Helper function to convert a number to a linked list
    if num == 0:
        return Node(0)
    head = curr = None
    while num > 0:
        digit = num % 10
        num //= 10
        node = Node(digit)
        if head is None:
            head = curr = node
        else:
            curr.next = node
            curr = curr.next
    return head

@pytest.mark.parametrize("n1,n2,expected", [
    (0, 0, 0),
    (123, 456, 579),
    (999, 1, 1000),
    (1234, 56789, 58023),
])
def test_sum_lists(n1, n2, expected):
    # Convert the numbers to linked lists
    n1_list = number_to_list(n1)
    n2_list = number_to_list(n2)

    # Call the sum_lists function and get the result list
    result_list = sum_lists(n1_list, n2_list)

    # Convert the result list to a number and compare with the expected value
    result = list_to_number(result_list)
    assert result == expected

@pytest.mark.parametrize("n1,n2,expected", [
    (0, 0, 0),
    (123, 456, 579),
    (999, 1, 1000),
    (1234, 56789, 58023),
])
def test_sum_lists_recur(n1, n2, expected):
    # Convert the numbers to linked lists
    n1_list = number_to_list(n1)
    n2_list = number_to_list(n2)

    # Call the sum_lists function and get the result list
    result_list = sum_list_recur(n1_list, n2_list)

    # Convert the result list to a number and compare with the expected value
    result = list_to_number(result_list)
    assert result == expected

def test_palindrome():
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')
    head.next.next.next.next = Node('m')

    assert palindrome(head)
    
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')

    assert not palindrome(head)

def test_palindrome_2():
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')
    head.next.next.next.next = Node('m')

    assert palindrome_2(head)
    
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')

    assert not palindrome_2(head)

def test_palindrome_3():
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')
    head.next.next.next.next = Node('m')

    assert palindrome_3(head)
    
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')

    assert not palindrome_3(head)

def test_palindrome_4():
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')
    head.next.next.next.next = Node('m')

    assert palindrome_4(head)
    
    head = Node('m')
    head.next = Node('a')
    head.next.next = Node('d')
    head.next.next.next = Node('a')

    assert not palindrome_4(head)

