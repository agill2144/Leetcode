func findWords(board [][]byte, words []string) []string {
    trie := NewTrie()
    for _, w := range words {
        trie.InsertWord(w)
    }

    wordsFound := make(map[string]struct{}, 0)
    for i := range board {
        for j := range board[0] {
            dfs(i, j, board, make(map[pair]struct{}), "", wordsFound, trie.root)
        }
    }

    ans := make([]string, 0, len(wordsFound))
    for w := range wordsFound {
        ans = append(ans, w)
    }

    return ans
}

type pair struct {
    r, c int
}

func dfs(r, c int, board [][]byte, visited map[pair]struct{}, curString string, wordsFound map[string]struct{}, curNode *TrieNode) {
    if curNode.isEnd {
        wordsFound[curString] = struct{}{}
    }

    inBounds := r >= 0 && r < len(board) && c >= 0 && c < len(board[0])
    _, isVisited := visited[pair{r, c}]
    if !inBounds || isVisited {
        return
    }
    if _, ok := curNode.data[board[r][c]]; !ok {
        return
    }

    visited[pair{r, c}] = struct{}{}

    newString := curString+string(board[r][c])
    dfs(r+1, c, board, visited, newString, wordsFound, curNode.data[board[r][c]])
    dfs(r-1, c, board, visited, newString, wordsFound, curNode.data[board[r][c]])
    dfs(r, c+1, board, visited, newString, wordsFound, curNode.data[board[r][c]])
    dfs(r, c-1, board, visited, newString, wordsFound, curNode.data[board[r][c]])
    
    delete(visited, pair{r, c})
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: NewTrieNode()}
}

type TrieNode struct {
    data map[byte]*TrieNode
    isEnd bool
}

func NewTrieNode() *TrieNode {
    return &TrieNode{data: make(map[byte]*TrieNode)}
}

func (t *Trie) InsertWord(word string) {
    cur := t.root

    for i := range word {
        if _, ok := cur.data[word[i]]; !ok {
            cur.data[word[i]] = NewTrieNode()
        }
        cur = cur.data[word[i]]
    }

    cur.isEnd = true
}
