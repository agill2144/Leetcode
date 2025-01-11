func findWords(board [][]byte, words []string) []string {
    root := newTrieNode()
    for i := 0; i < len(words); i++ {
        root.insert(words[i])
    }
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    set := map[string]bool{}
    var dfs func(r, c int, curr *trieNode, path string) 
    dfs = func(r, c int, curr *trieNode, path string) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == '#' || curr == nil || curr.childs[int(board[r][c]-'a')] == nil {return}
        
        
        // logic
        idx := int(board[r][c]-'a')
        tmp := board[r][c]
        board[r][c] = '#'
        path += string(tmp)
        curr = curr.childs[idx]
        if curr.isEnd {set[path] = true}
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], curr, path) 
        }
        board[r][c] = tmp
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(i, j, root, "")
        }
    }
    out := []string{}
    for k, _ := range set {out = append(out, k)}
    return out
}

type trieNode struct {
    isEnd bool
    childs [26]*trieNode
}

func newTrieNode() *trieNode {
    return &trieNode{childs: [26]*trieNode{}}
}

func (r *trieNode) insert(word string) {
    curr := r
    for i := 0; i < len(word); i++ {
        idx := int(word[i]-'a')
        if curr.childs[idx] == nil {
            curr.childs[idx] = newTrieNode()
        }
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}