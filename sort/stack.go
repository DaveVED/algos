package sort

import (
	"fmt"
)

type node struct {
	value int
	prev  *node
}

func push(head **node, value int) {
	newNode := &node{value: value, prev: *head}
	*head = newNode
}

func pop(head **node) (int, error) {
	if *head == nil {
		return 0, fmt.Errorf("stack is empty")
	}
	value := (*head).value
	*head = (*head).prev
	return value, nil
}

func peek(head *node) (int, error) {
	if head != nil {
		return head.value, nil
	}
	return 0, fmt.Errorf("stack is empty")
}