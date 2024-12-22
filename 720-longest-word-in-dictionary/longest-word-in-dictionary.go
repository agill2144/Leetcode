func longestWord(words []string) string {
    root := newTrieNode()
    for i := 0; i < len(words); i++ {
        root.insert(words[i])
    }
    res := ""
    var dfs func(curr *trieNode, path string)
    dfs = func(curr *trieNode, path string) {
        // base
        if curr == nil {return}

        // logic
        if len(path) > len(res) {res = path}
        for i := 0; i < len(curr.childs); i++ {
            if curr.childs[i] != nil && 
                curr.childs[i].isEnd {
                    path += string(i+'a')
                    dfs(curr.childs[i], path)
                    path = path[:len(path)-1]
            }
        }
    }
    dfs(root, "")
    return res
}


type trieNode struct {
    childs [26]*trieNode
    isEnd bool
}

func newTrieNode() *trieNode{
    return &trieNode{childs:[26]*trieNode{}}
}

func (r *trieNode) insert(word string)  {
    curr := r
    for i := 0; i < len(word); i++ {
        idx := int(word[i] - 'a')
        if curr.childs[idx] == nil {curr.childs[idx] = newTrieNode()}
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}

