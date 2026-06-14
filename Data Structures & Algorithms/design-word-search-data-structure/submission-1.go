type WordDictionary struct {


	children [26]*WordDictionary

	isEnd bool
    
}

func Constructor() WordDictionary {

	return WordDictionary{}
    
}

func (this *WordDictionary) AddWord(word string)  {


	node := this

	for _, ch := range word {

		ind := ch - 'a'

		if node.children[ind] == nil {
			node.children[ind] = &WordDictionary{}
		}

		node = node.children[ind]

	}

	node.isEnd = true
    
}

func (this *WordDictionary) Search(word string) bool {

	var dfs  func(node *WordDictionary, pos int) bool 


	dfs = func(node *WordDictionary, pos int) bool {

		if node == nil {
			return false
		}


		if pos == len(word) {

			return node.isEnd
		}


		ch := word[pos]


		if ch == '.' {

			for _, child := range node.children {

				if child != nil && dfs(child, pos+1) {

					return true
				}


			}

			return false

		}

		ind := ch - 'a'
		child := node.children[ind] 
		return dfs(child, pos +1)
	}


	return dfs(this, 0)

    
}
