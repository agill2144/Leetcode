func suggestedProducts(products []string, searchWord string) [][]string {
    root := &trie{childrens: [26]*trie{}}
    for _, word := range products {
        insert(root, word)
    }
    out := [][]string{}
    prefix := new(strings.Builder)
    for i := 0; i < len(searchWord); i++ {
        prefix.WriteByte(searchWord[i])
        out = append(out, getKWordsWithPrefix(root, 3, prefix.String()))
    } 
    return out
}

type trie struct {
    isEnd bool
    childrens [26]*trie
    words []string
}


func insert(root *trie, word string)  {
    curr := root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &trie{childrens: [26]*trie{}, words: []string{}}
        }
        if len(curr.childrens[charIdx].words) < 3 {
            curr.childrens[charIdx].words = append(curr.childrens[charIdx].words, word)           
        }
        curr = curr.childrens[charIdx]
        
    }
    curr.isEnd = true
}


func getKWordsWithPrefix(root *trie, k int, prefix string) []string {
    curr := root
    for i := 0; i < len(prefix); i++ {
        charIdx := prefix[i]-'a'
        if curr.childrens[charIdx] == nil {return nil}
        curr = curr.childrens[charIdx]
    }
    words := []string{}
    var dfs func(r *trie, path string)
    dfs = func(r *trie, path string) {
        // base
        if len(words) == k {return}

        // logic
        if r.isEnd {
            words = append(words, path)
        }
        for i := 0; i < 26; i++ {
            if r.childrens[i] != nil {
                dfs(r.childrens[i], path + string(i+'a'))
            }
        }
    }
    dfs(curr, prefix)
    return words
}

