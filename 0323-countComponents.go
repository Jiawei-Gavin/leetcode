package leetcode

func countComponents(n int, edges [][]int) int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	count := 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(adj, i, visited)
			count++
		}
	}
	return count
}

func dfs(adj [][]int, u int, visited []bool) {
	visited[u] = true
	successors := adj[u]
	for _, successor := range successors {
		if !visited[successor] {
			dfs(adj, successor, visited)
		}
	}
}
