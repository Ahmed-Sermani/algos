

from ch3.min_stack import MinStack, MinStack2
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




