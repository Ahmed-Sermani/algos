package structures

type Queue[T any] struct {
	data []T
	size int
	head int
	tail int
}

func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		data: make([]T, size+1),
		size: size + 1,
	}
}

func (q *Queue[T]) isfull() bool {
	return (q.tail+1 == q.head) || (q.head == 0 && q.tail == q.size-1)
}

func (q *Queue[T]) isempty() bool {
	return q.tail == q.head
}

func (q *Queue[T]) Enqueue(x T) {
	if q.isfull() {
		panic("overflow!")
	}

	q.data[q.tail] = x
	if q.tail == q.size-1 {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *Queue[T]) Dequeue() T {
	if q.isempty() {
		panic("underflow")
	}

	x := q.data[q.head]

	if q.head == q.size-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x
}
