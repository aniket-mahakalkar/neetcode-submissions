/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {


	cloned := make(map[*Node]*Node)


	var dfs func(node *Node) *Node 

	dfs = func(node *Node) *Node {

		if node == nil {
			return nil
		}

		if copiedNode, exist :=  cloned[node]; exist {
			return copiedNode
		}

	

		copied := &Node{Val : node.Val}

		cloned[node] = copied

		for _, neigh := range node.Neighbors {

			copied.Neighbors = append(copied.Neighbors, dfs(neigh))
		}

		return copied

	}

	return dfs(node)
    
}
