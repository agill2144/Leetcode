type Trie struct {
    children map[byte]*Trie
    word string
}

func (t *Trie) Add(word string) {
    c := t
    for i := 0; i < len(word); i++ {
        v := word[i]
        if n, ok := c.children[v]; ok {
            c = n
        } else {
            c.children[v] = &Trie{children: make(map[byte]*Trie)}
            c = c.children[v]
        }
    }
    c.word = word
}

func findWords(board [][]byte, words []string) []string {
    // Construct trie from input words
    trie := &Trie{children: make(map[byte]*Trie)}
    for _, word := range words {
        trie.Add(word)
    }

    moves := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    adj := func(row, col int) [][2]int {
        var ans [][2]int
        for _, m := range moves {
            r, c := row+m[0], col+m[1]
            if r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) {
                ans = append(ans, [2]int{r, c})
            }
        }
        return ans
    }

    ans := make(map[string]bool)

    var dfs func(int, int, *Trie)
    dfs = func(row, col int, t *Trie) {
        v := board[row][col]
        board[row][col] = '_'
        // Backtracking with defer!
        defer func() {
            board[row][col] = v
        }()
        if tn, ok := t.children[v]; ok {
            t = tn
        } else {
            return
        }
        if len(t.word) > 0 {
            ans[t.word] = true
        }

        for _, n := range adj(row, col) {
            v := board[n[0]][n[1]]
            if v == '_' {
                continue
            }
            dfs(n[0], n[1], t)
        }
    }

    for row := 0; row < len(board); row++ {
        for col := 0; col < len(board[0]); col++ {
            dfs(row, col, trie)
        }
    }
    var ans2 []string
    for w, _ := range ans {
        ans2 = append(ans2, w)
    }
    return ans2
}
