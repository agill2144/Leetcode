func findWords(board [][]byte, words []string) []string {
    root := &trieNode{childs: [26]*trieNode{}}
    for _, word := range words {
        insert(root, word)
    }
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    // outSet := map[string]struct{}{}
    out := []string{}

    var dfs func(r *trieNode, i, j int, path string)
    dfs = func(r *trieNode, i, j int, path string) {
        // base
        if i < 0 || i == m || j < 0 || j == n || board[i][j] == '#' {return}
        
        // logic
        path += string(board[i][j])
        newRoot := search(root, path)
        // not found
        if newRoot==nil {return}
        if newRoot.isEnd && !newRoot.used {
            newRoot.used = true
            out = append(out, path)
        }
        
        
        tmp := board[i][j]
        board[i][j] = '#'
        for _, dir := range dirs {
            dfs(newRoot, i+dir[0], j+dir[1], path)
        }
        board[i][j]= tmp

    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(root, i,j, "")
        }
    }
    
    return out
}




type trieNode struct {
    isEnd bool
    childs [26]*trieNode
    used bool // acting as my set :D 
}



func insert(root *trieNode, word string)  {
    curr := root
    for _, char := range word {
        if curr.childs[char-'a'] == nil {
            curr.childs[char-'a'] = &trieNode{isEnd: false, childs: [26]*trieNode{}}
        }
        curr = curr.childs[char-'a']
    }
    curr.isEnd = true
}

func search(root *trieNode, prefix string) *trieNode {
    curr := root
    for _, char := range prefix {
        idx := char-'a'
        if curr.childs[idx] == nil {return nil}
        curr = curr.childs[idx]
    }
    return curr
}
