package cci

import "github.com/Ahmed-Sermani/algos/structures"

/*
	Delete Middle Node:
	Implement an algorithm to delete a node in the middle
	(i.e., any node but the first and last node, not necessarily the exact middle)
	of a singly linked list, given only access to that node.
*/

func DeleteMiddleNode(node *structures.Node[int]) {
	if node == nil || node.Next == nil {
		panic("cant be deleted")
	}
	node.Data = node.Next.Data
	node.Next = node.Next.Next
}
