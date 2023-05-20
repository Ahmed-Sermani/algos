
# Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.

class MyQueue:

    def __init__(self) -> None:
        self._newest = []
        self._oldest = []


    def add(self, elem):
        if self._newest or (not self._newest and not self._oldest):
            self._newest.append(elem)
            return

        elif not self._newest and self._oldest:
            for _ in range(len(self._oldest)):
                self._newest.append(self._oldest.pop())
        self._newest.append(elem)


    def remove(self):
        if not self._newest and not self._oldest:
            raise Exception("the queue is empty!")

        elif self._oldest:
            return self._oldest.pop()

        self._shift_newest_to_oldest()
        return self._oldest.pop()

    def peek(self):
        if not self._newest and not self._oldest:
            raise Exception("the queue is empty!")

        elif self._oldest:
            return self._oldest[len(self._oldest)-1]

        self._shift_newest_to_oldest()
        return self._oldest[len(self._oldest)-1]


    def _shift_newest_to_oldest(self):
        if not self._oldest:
            for _ in range(len(self._newest)):
                self._oldest.append(self._newest.pop())


    def __repr__(self):
        return f"Queue(newest=[{','.join([ str(i) for i in self._newest])}], oldest=[{','.join([ str(i) for i in self._oldest])}])"

