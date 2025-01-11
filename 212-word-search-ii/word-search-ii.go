func findWords(board [][]byte, words []string) []string {
    root := newTrieNode()
    for i := 0; i < len(words); i++ {
        root.insert(words[i])
    }
    m := len(board)
    n := len(board[0])
    dirs := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
    out := []string{}
    var visited byte = '#'
    var dfs func(r, c int, curr *trieNode) 
    dfs = func(r, c int, curr *trieNode) {
        // base
        if r < 0 || r == m || c < 0 || c == n || board[r][c] == visited || curr == nil  {return}
        
        
        // logic
        idx := int(board[r][c]-'a')
        curr = curr.childs[idx]
        if curr == nil {return}
        if curr.word != "" {
            out = append(out, curr.word)
            // remove the word from being used again
            // this is acting like our hashset
            curr.word = ""
        }
        tmp := board[r][c]
        board[r][c] = visited
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1], curr) 
        }
        board[r][c] = tmp
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if root.childs[int(board[i][j]-'a')] == nil {continue}
            dfs(i, j, root)
        }
    }
    return out
}

type trieNode struct {
    isEnd bool
    word string
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
    curr.word = word
}