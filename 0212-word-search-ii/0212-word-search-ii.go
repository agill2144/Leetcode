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
    var dfs func(t *trieNode, r, c int, path string)
    dfs = func(t *trieNode, r, c int, path string)  {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == '#' {return}

        // logic
        tmp := board[r][c]
        char := string(tmp)
        path += string(board[r][c])
        
        
        newT := search(t, char)
        if newT == nil {return}
        if newT.isEnd { outSet[path] = struct{}{} }

        // visit cell
        board[r][c] = '#'
        for _, dir := range dirs {
            dfs(newT, r+dir[0], c+dir[1], path)
        }
        // backtrack
        board[r][c] = tmp
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(root, i,j, "")
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

func search(root *trieNode, prefix string) *trieNode {
    curr := root
    for _, char := range prefix {
        idx := char-'a'
        if curr.childs[idx] == nil {return nil}
        curr = curr.childs[idx]
    }
    return curr
}
