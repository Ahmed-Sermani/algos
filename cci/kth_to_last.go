package cci

import (
	"errors"

	"github.com/Ahmed-Sermani/algos/structures"
)

/*
	Return kth to last: implement an algorithm to find kth to last element in singly linked list.
*/

// time O(N) where N is the number of nodes in the list.
// space O(N)
func KthToLast(l *structures.LinkedList[int], k int) (*structures.Node[int], error) {
	if l.Head == nil {
		return nil, errors.New("not found")
	}
	var memo []*structures.Node[int]
	current := l.Head
	for ; current != nil; current = current.Next {
		memo = append(memo, current)
	}
	if k > len(memo) {
		return nil, errors.New("overflow")
	}
	return memo[len(memo)-k-1], nil
}

type kthToLast struct {
	idx  int
	node *structures.Node[int]
}

// time O(N)
// space O(N)
func KthToLast2(l *structures.LinkedList[int], k int) (*structures.Node[int], error) {
	node := recusor(l.Head, k)
	if node.node == nil {
		return nil, errors.New("not found")
	}
	return node.node, nil
}

func recusor(head *structures.Node[int], k int) *kthToLast {
	if head == nil {
		return &kthToLast{}
	}

	node := recusor(head.Next, k)
	if node.idx == k {
		node.node = head
	}
	node.idx++
	return node
}

// time O(N)
// space O(1)
func KthToLast3(l *structures.LinkedList[int], k int) (*structures.Node[int], error) {
	if l.Head == nil {
		return nil, errors.New("not found")
	}
	var listLength int
	var p1, p2 *structures.Node[int]
	p1 = l.Head
	for ; p1 != nil; p1 = p1.Next {
		listLength++
	}
	p2 = l.Head
	stepsToLast := listLength - k - 1
	currStep := 0
	for ; p2 != nil; p2 = p2.Next {
		if currStep == stepsToLast {
			return p2, nil
		}
		currStep++
	}
	return nil, errors.New("overflow")
}
