

# Stack of Plates: Imagine a (literal) stack of plates. If the stack gets too high, it might topple. 
# Therefore, in real life, we would likely start a new stack when the previous stack exceeds some 
# threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be 
# composed of several stacks and should create a new stack once the previous one exceeds capacity. 
# SetOfStacks. push() and SetOfStacks. pop() should behave identically to a single stack 
# (that is, pop () should return the same values as it would if there were just a single stack). 
# FOLLOW UP 
# Implement a function popAt(int index) which performs a pop operation on a specific substack.


class SetOfPlates:
    
    def __init__(self, threshold):
        self.threshold = threshold
        self.__stacks = []

    def push(self, elem):

        if len(self.__stacks) == 0 or self._is_full():
            self._add_stack()

        self._curr_stack().append(elem)

    def pop(self):
        if len(self.__stacks) == 0:
            raise Exception("stack in empty")
        popped = self._curr_stack().pop() 
        if self._is_empty():
            self._remove_stack()

        return popped

    def pop_at(self, i):

        # we assume that the substack should always be on the max capacity after the popping.

        if len(self.__stacks) == 0:
            raise Exception("stack in empty")

        return self.left_shift(i, True)
    
    def left_shift(self, index, remove_top):
        stack = self.__stacks[index]

        removed_item = stack.pop(-1 if remove_top else 0)

        if len(stack) == 0:
            self.__stacks.pop(index)
        elif len(self.__stacks) > index + 1:
            v = self.left_shift(index + 1, False)
            stack.append(v)

        return removed_item
    
    
    def _curr_stack(self):
        return self.__stacks[len(self.__stacks) - 1]

    def _add_stack(self):
        self.__stacks.append([])

    def _remove_stack(self):
        self.__stacks.pop()

    def _is_full(self):
        return len(self.__stacks[len(self.__stacks) - 1]) == self.threshold

    def _is_empty(self):
        return len(self.__stacks[len(self.__stacks) - 1]) == 0

