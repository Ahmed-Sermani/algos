package main

import "fmt"

type entry struct {
	v string
	t int
}

type bstNode struct {
	val *entry
	l   *bstNode
	r   *bstNode
}

type bst struct {
	root    *bstNode
	closest *bstNode
}

func (b *bst) insert(val *entry) {
	b.recInsert(b.root, val)
}

func (b *bst) recInsert(node *bstNode, value *entry) *bstNode {
	if b.root == nil {
		b.root = &bstNode{
			val: value,
		}
		return b.root
	}
	if node == nil {
		return &bstNode{val: value}
	}
	if value.t <= node.val.t {
		node.l = b.recInsert(node.l, value)
	}
	if value.t > node.val.t {
		node.r = b.recInsert(node.r, value)
	}
	return node
}

func (b *bst) find(timestamp int) *entry {
	b.closest = b.root
	node := b.recFind(b.root, timestamp)
	if node == nil {
		cls := b.closest
		b.closest = nil
		if cls != nil {
			return cls.val
		}
		return nil
	}
	b.closest = nil
	return node.val
}

func (b *bst) recFind(node *bstNode, t int) *bstNode {
	if node == nil {
		return nil
	}
	if node.val.t == t {
		b.closest = b.root
		return b.root
	}
	if t < node.val.t {
		b.closest = node.l
		return b.recFind(node.l, t)
	}
	if node.r != nil {
		b.closest = node.r
	}
	return b.recFind(node.r, t)
}

type Map map[string]*bst

type TimeMap struct {
	store Map
}

func Constructor() TimeMap {
	return TimeMap{make(Map)}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	etr := &entry{value, timestamp}
	if _, ok := tm.store[key]; ok {
		tm.store[key].insert(etr)
	} else {
		tm.store[key] = &bst{root: &bstNode{val: etr}, closest: nil}
	}
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	b := tm.store[key]
	etr := b.find(timestamp)
	if etr == nil {
		return ""
	}
	return etr.v
}

func main() {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	fmt.Println(tm.Get("foo", 1))
	fmt.Println(tm.Get("foo", 3))
	tm.Set("foo", "bar2", 4)
	fmt.Println(tm.Get("foo", 4))
}
