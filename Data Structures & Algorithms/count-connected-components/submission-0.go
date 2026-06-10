func countComponents(n int, edges [][]int) int {

	graph := make([][]int, n)


	for _, edge := range edges {

		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool,n)

	cnt := 0

	var dfs func(node int)

	dfs = func(node int) {

		if visited[node] {
			return
		}

		visited[node] = true


		for _, neigh := range graph[node]{

			dfs(neigh)
		}

	}

	for i := 0; i < n; i++ {

		if !visited[i] {

			cnt++
			dfs(i)
		}


	}

	return cnt
    
}
