func findWords(board [][]byte, words []string) []string {
    if words == nil || len(words) == 0 {return nil}
    m := len(board)
    n := len(board[0])
    
    root := &trieNode{childs: [26]*trieNode{}}
    for i := 0; i < len(words); i++ {
        insert(root, words[i])
    }
    outSet := map[string]struct{}{}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int, path string)
    dfs = func(r, c int, path string)  {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == '#' {return}

        // logic
        path += string(board[r][c])
        tmp := board[r][c]
        
        found, isEnd := search(root, path)
        if !found {return}
        if isEnd { outSet[path] = struct{}{} }

        // visit cell
        board[r][c] = '#'
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], path)
        }
        // backtrack
        board[r][c] = tmp
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(i,j, "")
        }
    }
    out := []string{}
    for k, _ := range outSet {
        out = append(out, k)
    }
    return out
}



type trieNode struct {
    isEnd bool
    childs [26]*trieNode
    
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

func search(root *trieNode, prefix string) (bool, bool) {
    curr := root
    for _, char := range prefix {
        idx := char-'a'
        if curr.childs[idx] == nil {return false, false}
        curr = curr.childs[idx]
    }
    return true, curr.isEnd
}
