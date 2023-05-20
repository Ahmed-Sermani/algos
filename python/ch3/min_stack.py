

# Stack Min: How would you design a stack which, 
# in addition to push and pop, has a function min which returns the minimum element? 
# Push, pop and min should all operate in 0(1) time.

from dataclasses import dataclass
import sys

@dataclass
class NodeWithMin:
    value: int
    min: int

class MinStack:
    def __init__(self):
        self.__stack = []

    def push(self, elm):
        new_min = min(elm, self.min())
        self.__stack.append(NodeWithMin(elm, new_min))

    def pop(self):
        if len(self.__stack) == 0:
            raise Exception("the stack is empty")
        return self.__stack.pop().value

    def peek(self):
        if len(self.__stack) == 0:
            raise Exception("the stack is empty")

        return self.__stack[len(self.__stack) - 1]

    def min(self):
        if len(self.__stack) == 0:
            return sys.maxsize
        return self.peek().min


class MinStack2:
    def __init__(self):
        self.__stack = []
        self.mins = []


    def push(self, elem):

        if not self.mins or elem <= self.mins[len(self.mins) - 1]:
            self.mins.append(elem)

        self.__stack.append(elem)

    def pop(self):
        poped = self.__stack.pop()
        if poped == self.min():
            self.mins.pop()

        return poped

    def min(self):
        if len(self.__stack) == 0:
            return sys.maxsize
        return self.mins[len(self.mins) - 1]

        
    def peek(self):
        if len(self.__stack) == 0:
            raise Exception("the stack is empty")

        return self.__stack[len(self.__stack) - 1]


