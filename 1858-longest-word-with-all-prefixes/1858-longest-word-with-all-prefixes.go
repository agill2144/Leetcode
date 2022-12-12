func longestWord(words []string) string {
    root := new(TrieNode)
    root.childrens = [26]*TrieNode{}
    root.isEnd = false
    // o(n)
    for _, word := range words {
        // o(k)
        insertWord(word, root)
    }
    // time for above loop : o(nk)

    
    maxWord := ""
    var backtrack func(r *TrieNode, pathBldr string)
    backtrack = func(r *TrieNode, pathBldr string) {
        // base
        if len(pathBldr) > len(maxWord) { 
            maxWord = pathBldr
        }

        
        // logic
        for i := 0; i <= 25; i++ {
            if r.childrens[i] != nil && r.childrens[i].isEnd {
                // TODO: find out how to backtrack in strings.Builder in golang ( still have no solution to this .. sigh )
                pathBldr += string('a'+i)
                backtrack(r.childrens[i], pathBldr)
                pathBldr = pathBldr[:len(pathBldr)-1]
            }
        }
    }
    backtrack(root, "")
    return maxWord
}

type TrieNode struct {
    isEnd bool
    childrens [26]*TrieNode
}

// time: o(k) where k is the len of word string
// space: o(k)
func insertWord(word string, root *TrieNode) {
    curr := root
    for i := 0; i < len(word); i++ {
        char := word[i]
        idx := char-'a'
        if curr.childrens[idx] == nil {
            curr.childrens[idx] = &TrieNode{
                childrens: [26]*TrieNode{},
            }
        }
        curr = curr.childrens[idx]
    }
    curr.isEnd = true

}
