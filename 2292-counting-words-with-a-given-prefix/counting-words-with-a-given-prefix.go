func prefixCount(words []string, pref string) int {
    root := newTrieNode()
    for i := 0; i < len(words); i++ {
        root.insert(words[i])
    }
    return root.prefixCount(pref)
}

type trieNode struct {
    childs [26]*trieNode
    isEnd bool
    count int
}

func newTrieNode() *trieNode {
    return &trieNode{childs:[26]*trieNode{}}
}

func (r *trieNode) insert(word string) {
    curr := r
    for i := 0; i < len(word); i++ {
        if curr.childs[int(word[i]-'a')] == nil {
            curr.childs[int(word[i]-'a')] = newTrieNode()
        }
        curr = curr.childs[int(word[i]-'a')]
        curr.count++
    }
    curr.isEnd = true
}

func (r *trieNode) prefixCount(pref string) int {
    curr := r
    for i := 0; i < len(pref); i++ {
        idx := int(pref[i]-'a')
        if curr.childs[idx] == nil {return 0}
        curr = curr.childs[idx]
    }
    return curr.count
}