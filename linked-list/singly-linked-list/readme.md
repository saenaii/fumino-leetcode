# Singly Linked List

> [Introduction](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1287/)
 
## Operation 
> data structure
```go
type Node struct {
    val int
    next *Node
}

type LinkedList struct {
    head *Node
}
```

### [Add Operation](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1288/)
#### Add at head
```go
func AddAtHead(val int, h *LinkedList) *LinkedList {
	node := &Node{
		val: val,
		next: h.head,
	}
	h.head = node
	return h
}
```

#### Add at tail
```go
func AddAtTail(val int, h *LinkedList) *LinkedList {
	node := &Node{val, nil}
	if h.head == nil {
		h.head = node
		return h
	}

	current := h.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
	return h
}
```

### [Delete Operation](https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1289/)

## Problems
* [Design Linked List](./707.design-linked-list.go)
