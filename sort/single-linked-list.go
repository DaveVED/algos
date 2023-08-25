package main 

import (
	"fmt"
)

type node struct {
	value int
	next *node
}

func addNode(head *node, value int) *node {
	newNode := &node{value: value}
	if head == nil {
	   head = newNode
	 } else {
	   current := head
	   for current.next != nil {
		  current = current.next
	   }
	   current.next = newNode
	}
	return head
 }

func main() {
	var head *node
	head = addNode(head, 10)
	head = addNode(head, 20)
	head = addNode(head, 30)
 
	current := head
	for current != nil {
	   fmt.Println(current.value)
	   current = current.next
	}
}