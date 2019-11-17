# Two Pointer
> [Introduction](https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/1211/)

## Problems
* [[19] remove-nth-node-from-end-of-list](./19.remove-nth-node-from-end-of-list.go)
* [[141] linked-list-cycle](./141.linked-list-cycle.go)
* [[142] linked-list-cycle2](./142.linked-list-cycle-ii.go)
* [[160] intersection-of-two-linked-lists](./160.intersection-of-two-linked-lists.go)

## Template
```go
slow, fast := head, head.Next
for slow != nil && fast != nil && fast.Next != nil {
    if slow == fast {
        return true
    }
    slow = slow.Next
    fast = fast.Next.Next
}
return false
```

## Tips
1. Always examine if the node is null before you call next field.
1. Carefully define the end of conditions of your loop.