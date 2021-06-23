package leetcode

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	edges := make([][]int, n)
	visited := make([]int, n)
	for i := 0; i < len(graph); i++ {
		edges[graph[i][0]] = append(edges[graph[i][0]], graph[i][1])
	}
	var queue []int
	queue = append(queue, start)
	for len(queue) > 0 {
		tmp := queue[0]
		queue = queue[1:]
		visited[tmp] = 1
		if tmp == target {
			return true
		}
		for i := 0; i < len(edges[tmp]); i++ {
			if visited[edges[tmp][i]] == 0 {
				queue = append(queue, edges[tmp][i])
			}
		}
	}
	return false
}
