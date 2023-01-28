
type trieNode struct {
    isEnd bool
    childs [26]*trieNode
}

func insert(root *trieNode, word string) {
    curr := root
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if curr.childs[idx] == nil {
            curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
        }
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}


func longestWord(words []string) string {
    root := &trieNode{childs: [26]*trieNode{}}
    for _, word := range words {
        insert(root, word)
    }
    maxLenWord := ""
    var dfs func(r *trieNode, path string)
    dfs = func(r *trieNode, path string) {
        // base
        if len(path) > len(maxLenWord) {
            maxLenWord = path
        }
        if r == nil {return}
        
        // logic
        for i := 0; i < len(r.childs); i++ {
            child := r.childs[i]
            if child != nil && child.isEnd {
                path += string(i+'a')
                dfs(child, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(root, "")
    return maxLenWord
}