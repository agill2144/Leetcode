
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
        idx := word[i] - 'a'
        if curr.childs[idx] == nil {
            curr.childs[idx] = newTrieNode()
        }
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}

func (r *trieNode) findPrefix(word string) (string, bool) {
    curr := r
    tmp := new(strings.Builder)

    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if curr.childs[idx] == nil { return "", false }
        tmp.WriteByte(word[i])
        if curr.childs[idx].isEnd {return tmp.String(), true }
        curr = curr.childs[idx]
    }
    return  tmp.String(), curr.isEnd
}


func replaceWords(dictionary []string, sentence string) string {
    root := newTrieNode()
    for i := 0; i < len(dictionary); i++ {
        root.insert(dictionary[i])
    }
    sSplit := strings.Split(sentence, " ")
    out := new(strings.Builder)
    for idx, word := range sSplit {
        prefix, found := root.findPrefix(word)
        if found {
            out.WriteString(prefix)
        } else {
            out.WriteString(word)
        }
        if idx != len(sSplit)-1 {
            out.WriteString(" ")
        }
    }
    return out.String()
}