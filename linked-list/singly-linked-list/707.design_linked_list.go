package singly_linked_list

type MyLinkedList struct {
	head *Node
	length int
	last *Node
}

type Node struct {
	val int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}

	current := this.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	node := &Node{
		val:  val,
		next: this.head,
	}
	this.head = node

	if this.length == 0 {
		this.last = node
		this.last.next = nil
	}

	this.length++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	node := &Node{
		val:  val,
		next: nil,
	}

	if this.length == 0 {
		this.head = node
		this.last = this.head
	} else {
		this.last.next = node
		this.last = node
	}

	this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	prev := this.head
	for i := 1; i < index; i++ {
		prev = prev.next
	}
	node := &Node{
		val: val,
		next: prev.next,
	}
	prev.next = node
	this.length++
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= this.length {
		return
	}

	if index == 0 {
		this.head = this.head.next
		this.length--
		return
	}

	prev := &Node{0, this.head}
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	current := prev.next
	if current.next == nil {
		prev.next = nil
		this.last = prev
	} else {
		prev.next = current.next
	}
	this.length--
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
