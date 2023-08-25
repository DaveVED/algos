package sort

type node struct {
	value int
	next *node
}

func enqueue(head **node, tail **node, value int) {
    newNode := &node{value: value, next: nil}
    if *tail == nil {
        *head = newNode
        *tail = newNode
    } else {
        (*tail).next = newNode
        *tail = newNode
    }
}


func deque(head **node, tail **node) (int, error) {
    if *head == nil {
        return 0, fmt.Errorf("queue is empty")
    }
    value := (*head).value
    *head = (*head).next
    if *head == nil {
        *tail = nil
    }
    return value, nil
}


func peek(head *node) int {
    if head != nil {
        return head.value
    }
    return -1
}

func length(head *node) int {
    length := 0
    current := head
    for current != nil {
        length++
        current = current.next
    }
    return length
}