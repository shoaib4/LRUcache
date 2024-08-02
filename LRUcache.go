package main

import "fmt"

//2. make it extensable
// lru cache
// dll  NODE == ( val, left node* , right node* )
// hash map -> (key, node*)

type Node struct {
	val   string
	left  *Node
	right *Node
}

type DLL struct {
	head    *Node
	tail    *Node
	size    int
	maxSize int
}

type LRUCache struct {
	keyMap map[string]*Node
	list   DLL
}

func (c *LRUCache) PrintMap() {
	for k, v := range c.keyMap {
		fmt.Println("( ", k, " ", v, ") ")
	}
}
func (c *LRUCache) PrintList() {
	fmt.Println(c.list.head, "  ", c.list.tail)
	curr := c.list.head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.left
	}
}

func (c *LRUCache) Get(key string) string {

	valNode, ok := c.keyMap[key]
	if !ok {
		fmt.Println("No Key found!!!")
		return ""
	}
	val := valNode.val
	leftOfValNode := valNode.left
	rightOfValNode := valNode.right

	if leftOfValNode != nil {
		leftOfValNode.right = rightOfValNode
	}
	if rightOfValNode != nil {
		rightOfValNode.left = leftOfValNode
	}
	currHead := c.list.head
	valNode.left = currHead
	if currHead != nil {
		currHead.right = valNode
	}
	if c.list.tail == valNode {
		c.list.tail = valNode.right
	}
	valNode.right = nil
	c.list.head = valNode
	return val
}

func (c *LRUCache) Set(key string, val string) {
	fmt.Println(c.list.maxSize)
	valNode, ok := c.keyMap[key]
	if !ok {
		fmt.Println("adding ", key)
		newNode := &Node{
			val: val,
		}
		c.keyMap[key] = newNode
		innitHead := c.list.head
		newNode.left = innitHead
		if innitHead != nil {
			innitHead.right = newNode
		}
		c.list.head = newNode
		if c.list.tail == nil {
			c.list.tail = newNode
		}
		c.list.size++
		if c.list.size == c.list.maxSize {
			fmt.Println("Exxxited")
			tailNode := c.list.tail
			rightTail := tailNode.right
			c.list.tail = rightTail
			rightTail.left = nil
			c.list.size--
			//	todo delete form map
		}
		//fmt.Println(c)
		return
	}
	valNode.val = val
	_ = c.Get(key)

	return
}
