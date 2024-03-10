func longestWord(words []string) string {
    root := &trieNode{childs: [26]*trieNode{}}
    for i := 0; i < len(words); i++ {
        word := words[i]
        root.insert(word)
    }
    out := ""
    var dfs func(curr *trieNode, path []string)
    dfs = func(curr *trieNode, path []string) {
        // base
        pathStr := strings.Join(path, "")
        if path != nil && len(pathStr) > len(out) {
            out = pathStr
        }

        // logic
        for i := 0; i < len(curr.childs); i++ {
            if curr.childs[i] != nil && curr.childs[i].isEnd {
                path = append(path, string(i+'a'))
                dfs(curr.childs[i], path)
                path = path[:len(path)-1]
            }
        }
    }

    dfs(root, nil)
    return out

}


type trieNode struct {
    isEnd bool
    childs [26]*trieNode
}

func (t *trieNode) insert(word string) {
    curr := t
    for i := 0; i < len(word); i++ {
        char := word[i]
        idx := int(char-'a')
        if curr.childs[idx] == nil {
            curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
        }
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}

