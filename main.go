package main

import (
	. "github.com/hafizmfadli/graph-theory/graph"
)

func main() {
	n1 := Node{
		Val: 1,
	}

	n2 := Node{
		Val: 2,
	}

	n3 := Node{
		Val: 3,
	}

	n4 := Node{
		Val: 4,
	}

	n5 := Node{
		Val: 5,
	}

	n1.Neighbors = append(n1.Neighbors, &n4, &n5, &n2)
	n2.Neighbors = append(n2.Neighbors, &n1, &n3)
	n3.Neighbors = append(n3.Neighbors, &n2, &n4)
	n4.Neighbors = append(n4.Neighbors, &n3, &n1)

	// Bfs(&n1)
	tracking := map[*Node]bool{}
	Dfs(tracking, &n1)
}
