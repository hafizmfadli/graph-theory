package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

type Queue []Node

func (q *Queue) Enqueue(item *Node) {
	*q = append(*q, *item)
}

func (q *Queue) Dequeue() *Node {
	if len(*q) == 0 {
		return nil
	}

	item := (*q)[0]
	*q = (*q)[1:]
	return &item
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
