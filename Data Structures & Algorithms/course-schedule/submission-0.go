func canFinish(numCourses int, prerequisites [][]int) bool {

	graph := make([][]int, numCourses)

	for _, preq := range prerequisites {

		graph[preq[1]] = append(graph[preq[1]], preq[0])


	}

	state := make([]int, numCourses)

	var dfs func(node int) bool 


	dfs = func(node int) bool {

		if state[node] == 1{

			return false
		}

		if state[node] == 2 {
			return true
		}

		state[node] = 1


		for _, neigh := range graph[node] {

			if !dfs(neigh) {

				return false
			}

		}

		state[node] = 2


		return true

		
	}


	for i := 0; i < numCourses; i++ {

		if !dfs(i){

			return false
		}

	}

	return true
    
}
