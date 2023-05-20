

from ch3 import Stack
from ch3.min_stack import MinStack, MinStack2
from ch3.queue_via_stacks import MyQueue
from ch3.sort_stack import sort_stack
from ch3.stack_of_plates import SetOfPlates


def test_min_stack():

    stack = MinStack()
    stack.push(5)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 5

    stack = MinStack()
    stack.push(5)
    stack.push(3)
    assert stack.peek().value == 3
    assert stack.peek().min == 3
    stack.pop()
    assert stack.peek().value == 5
    assert stack.peek().min == 5

    stack = MinStack()
    stack.push(5)
    stack.push(3)
    assert stack.min() == 3
    stack.pop()
    assert stack.min() == 5

def test_min_stack_2():
    stack = MinStack2()
    stack.push(5)
    stack.push(3)
    assert stack.pop() == 3
    assert stack.pop() == 5

    stack = MinStack2()
    stack.push(5)
    stack.push(3)
    assert stack.peek() == 3
    assert stack.min() == 3
    stack.pop()
    assert stack.peek() == 5
    assert stack.min() == 5

    stack = MinStack2()
    stack.push(5)
    stack.push(3)
    assert stack.min() == 3
    stack.pop()
    assert stack.min() == 5

def test_set_of_plates():
    stack = SetOfPlates(2)
    stack.push(1)
    stack.push(2)
    stack.push(3)
    assert len(stack._SetOfPlates__stacks) == 2
    assert stack._SetOfPlates__stacks[0] == [1, 2]
    assert stack._SetOfPlates__stacks[1] == [3]

    stack = SetOfPlates(2)
    stack.push(1)
    stack.push(2)
    stack.push(3)
    popped = stack.pop()

    assert popped == 3
    assert len(stack._SetOfPlates__stacks) == 1
    assert stack._SetOfPlates__stacks[0] == [1, 2]

    stack = SetOfPlates(2)
    stack.push(1)
    stack.push(2)
    stack.push(3)
    stack.push(4)
    stack.push(5)
    stack.push(6)

    popped = stack.pop_at(0)

    assert popped == 2
    assert len(stack._SetOfPlates__stacks) == 3
    assert stack._SetOfPlates__stacks[0] == [1, 3]
    assert stack._SetOfPlates__stacks[1] == [4, 5]
    assert stack._SetOfPlates__stacks[2] == [6]
    
    popped = stack.pop_at(0)
    
    assert popped == 3
    assert len(stack._SetOfPlates__stacks) == 2
    assert stack._SetOfPlates__stacks[0] == [1, 4]
    assert stack._SetOfPlates__stacks[1] == [5, 6]


def test_queue_via_stacks():
    q = MyQueue()
    q.add(1)
    q.add(2)
    q.add(3)

    assert q.remove() == 1
    assert q.remove() == 2
    q.add(4)
    assert q.remove() == 3
    assert q.remove() == 4

    q = MyQueue()
    q.add(1)
    q.add(2)

    assert q.peek() == 1
    assert q.remove() == 1
    assert q.peek() == 2


def test_sort_stack():
    stack = Stack()
    sorted_stack = sort_stack(stack)
    assert sorted_stack.is_empty()

    stack = Stack()
    stack.push(5)
    sorted_stack = sort_stack(stack)
    assert not sorted_stack.is_empty()
    assert sorted_stack.pop() == 5
    assert sorted_stack.is_empty()

    stack = Stack()
    stack.push(5)
    stack.push(3)
    stack.push(1)
    stack.push(4)
    stack.push(2)
    sorted_stack = sort_stack(stack)
    assert not sorted_stack.is_empty()
    assert sorted_stack.pop() == 1
    assert sorted_stack.pop() == 2
    assert sorted_stack.pop() == 3
    assert sorted_stack.pop() == 4
    assert sorted_stack.pop() == 5
    assert sorted_stack.is_empty()
