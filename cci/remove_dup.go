package cci

/*
	remove duplicate from linked list
*/
import "github.com/Ahmed-Sermani/algos/structures"

// time O(N)
// space O(N)
func RemoveDup(l *structures.LinkedList[int]) {
	if l.Head == nil {
		return
	}

	m := map[int]struct{}{}
	curr := l.Head
	for curr != nil {
		if _, duplicated := m[curr.Data]; duplicated {
			curr.Prev.Next = curr.Next
			if curr.Next != nil {
				curr.Next.Prev = curr.Prev
			}
		} else {
			m[curr.Data] = struct{}{}
		}
		curr = curr.Next
	}
}

// time O(N^2)
// space O(1)
func RemoveDup2(l *structures.LinkedList[int]) {
	if l.Head == nil {
		return
	}
	var current, runner *structures.Node[int]
	current = l.Head
	for current != nil {
		runner = current.Next
		for runner != nil {
			if current.Data == runner.Data {
				runner.Prev.Next = runner.Next
				if runner.Next != nil {
					runner.Next.Prev = runner.Prev
				}
			}
			runner = runner.Next
		}
		current = current.Next
	}
}
