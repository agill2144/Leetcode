func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    root := &Trie{childrens: [26]*Trie{}}
    for _, word := range products {
        insert(root, word)
    }
    out := [][]string{}
    wordSoFar := new(strings.Builder)
    for i := 0; i < len(searchWord); i++ {
        wordSoFar.WriteByte(searchWord[i])
        tmp := searchInTrie(root, wordSoFar.String())
       
        out = append(out, tmp)
    } 
    return out
}

type Trie struct {
    isEnd bool
    childrens [26]*Trie
    words []string
}


func insert(root *Trie, word string)  {
    curr := root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &Trie{childrens: [26]*Trie{}, words: []string{}}
        }
        if len(curr.childrens[charIdx].words) < 3 {
            curr.childrens[charIdx].words = append(curr.childrens[charIdx].words, word)           
        }
        curr = curr.childrens[charIdx]
        
    }
    curr.isEnd = true
}


func searchInTrie(root *Trie, word string) []string {
    curr := root
    for i := 0; i < len(word); i++ {
        char := word[i]
        charIdx := char-'a'
        if curr.childrens[charIdx] == nil {return nil}
        curr = curr.childrens[charIdx]
    }
    return curr.words
}

