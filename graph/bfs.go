package graph

import "fmt"

// bfs print value using bfs approach
func Bfs(root *Node) {
	q := Queue{}
	visitedNode := map[*Node]bool{}
	q.Enqueue(root)
	visitedNode[root] = true
	for {
		if q.Len() == 0 {
			break
		}

		lenQ := q.Len()
		for i := 0; i < lenQ; i++ {
			node := q.Dequeue()
			if node != nil {
				fmt.Println(node.Val)
				for _, neighbor := range node.Neighbors {
					if neighbor != nil && !visitedNode[neighbor] {
						q.Enqueue(neighbor)
						visitedNode[neighbor] = true
					}
				}
			}
		}
	}
}
