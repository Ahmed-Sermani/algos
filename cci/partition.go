package cci

/*
    Partition: partition a linked list around a value x, such that all nodes less than x come before all nodes
    greater than x.
*/

import "github.com/Ahmed-Sermani/algos/structures"

func Partition(head *structures.Node[int], pivot int) *structures.Node[int] {
    if head == nil {
        return head
    }
    var partition1, partition2, head1, head2 *structures.Node[int]
    
    curr := head
    for curr != nil {
        if curr.Data < pivot {
            if partition1 == nil {
                partition1 = &structures.Node[int]{Data: curr.Data}
                head1 = partition1
            } else {
                partition1.Next = &structures.Node[int]{Data: curr.Data}
                partition1 = partition1.Next
            }
        } else {
            if partition2 == nil {
                partition2 = &structures.Node[int]{Data: curr.Data}
                head2 = partition2
            } else {
                partition2.Next = &structures.Node[int]{Data: curr.Data}
                partition2 = partition2.Next
            }
        }
    }
    partition1.Next = head2
    *head = *head1
    return head
}

func Partition2(node *structures.Node[int], pivot int) *structures.Node[int] {
    head := node
    tail := node

    for node != nil {
        next := node.Next
        if node.Data < pivot {
            node.Next = head
            head = node
        } else {
            tail.Next = node
            tail = node
        }
        node = next
    }
    tail.Next = nil
    return head
}
