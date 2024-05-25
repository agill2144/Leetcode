type WordDictionary struct {
    root *trieNode
}


func Constructor() WordDictionary {
    return WordDictionary{
        root: &trieNode{childs: [26]*trieNode{}},
    }    
}


func (this *WordDictionary) AddWord(word string)  {
    this.root.insert(word)    
}


func (this *WordDictionary) Search(word string) bool {
    var dfs func(ptr int, node *trieNode) bool
    dfs = func(ptr int, node *trieNode) bool {
        // base
        if ptr == len(word) {
            return node.isEnd
        }
        
        // logic
        char := word[ptr]
        // char == '.'
        // char = 'a'
        for i := 0; i < len(node.childs); i++ {
            child := node.childs[i]
            if child != nil {
                childChar := byte(i+'a')
                if char == '.' || char == childChar {
                    if dfs(ptr+1,child) {return true}
                }
            }
        }
        return false
    }
    return dfs(0, this.root)
}


type trieNode struct {
	isEnd  bool
	childs [26]*trieNode
}

func (r *trieNode) insert(word string) {
	curr := r
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if curr.childs[idx] == nil {
			curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
		}
		curr = curr.childs[idx]
	}
	curr.isEnd = true
}



/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */