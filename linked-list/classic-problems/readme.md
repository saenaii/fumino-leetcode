# Classic Problems

## Problems
* [[203 remove-linked-list-elements](./203.remove-linked-list-elements.go)]
* [[206 reverse-linked-list](./206.reverse-linked-list.go)]
* [[234 palindrome-linked-list](./234.palindrome-linked-list.go)]

## Tips
### Going through some test cases will save you time.
It is not easy to debug when using a linked list. Therefor, it is always useful to try several different examples on you own to validate your algorithm before writing code.

### Fell free to use several pointers at the same time.
Sometimes when you design an algorithm for a linked list problem, there might be several nodes you want to track at the same time. You should keep in mind which nodes you need to track and fell free to use several different pointers to track these nodes at the same time.

### In many cases, you need to track the previous node of the current node.
You are not able to track back the previous node in a singly linked list. So you have to store not only the current node but also the previous node.