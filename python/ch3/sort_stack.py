
# Sort Stack: Write a program to sort a stack such that 
# the smallest items are on the top. You can use an additional temporary stack, 
# but you may not copy the elements into any other data structure (such as an array). 
# The stack supports the following operations: push, pop, peek, and isEmpty.


# Time: O(N^2)
# Space: O(N)

from ch3 import Stack


def sort_stack(stack: Stack):
    temp = Stack()
    while not stack.is_empty():
        if temp.is_empty():
            temp.push(stack.pop())
            continue

        v1 = stack.pop()
        temp_pops = 0
        while not temp.is_empty() and temp.peek() > v1:
            stack.push(temp.pop())
            temp_pops += 1

        temp.push(v1)
        while temp_pops:
            temp.push(stack.pop())
            temp_pops -= 1

    while not temp.is_empty():
        stack.push(temp.pop())

    return stack
