type PrefixTree struct {
	children [26]*PrefixTree
	isEnd   bool
}

func Constructor() PrefixTree {

	return PrefixTree{}
    
}

func (this *PrefixTree) Insert(word string) {

	node := this

	for _, ch := range word {

		ind := ch - 'a'
		if node.children[ind] == nil {
			node.children[ind] = &PrefixTree{}
		}

		node = node.children[ind]
	}

	node.isEnd = true

}

func (this *PrefixTree) Search(word string) bool {

	node := this

	for _, ch := range word {

		ind := ch - 'a'
		

		if node.children[ind] == nil {
			return false
		}

		node =  node.children[ind]

	}

	return node.isEnd

}

func (this *PrefixTree) StartsWith(prefix string) bool {

	node := this


	for _, ch := range prefix {

		ind := ch - 'a'

		if node.children[ind] == nil {
			return false
		}

		node = node.children[ind]
	}

	return true

}
