package graph

import "fmt"

// dfs print value using dfs approach
func Dfs(visitTracking map[*Node]bool, node *Node) {
	if visitTracking[node] {
		return
	}

	fmt.Println(node.Val)
	visitTracking[node] = true
	for _, neighbor := range node.Neighbors {
		Dfs(visitTracking, neighbor)
	}
}
