
type trieNode struct {
    isEnd bool
    childs [26]*trieNode
}

func newTrieNode() *trieNode {
    return &trieNode{childs: [26]*trieNode{}}
}

// time: o(n) where n is the len of word
// space: o(26*n) where n is the len of word ( for each char we allocate 26 sized array )
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

// time: o(n) - we found the entire word as is, no early exits 
// space: o(n) - we found the entire word as is, and we allocated entire word space in str builder
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
    // o(d) * o(k) where d is number of words in dict, k is the avg len of each word
    // therefore o(dk)
    for i := 0; i < len(dictionary); i++ {
        root.insert(dictionary[i])
    }
    
    // let w == len of words in sentence
    // o(w) space
    sSplit := strings.Split(sentence, " ")
    
    // o(w) space 
    out := new(strings.Builder) 

    // o(w) * o(k) where w is the number of words we have in sentence, k is the avg len of each word
    // for each w word, we do a prefixSearch where prefixSearch takes o(k) time worse case
    // therefore o(wk)
    for idx, word := range sSplit {
        prefix, found := root.findPrefix(word) // time: o(k), space: o(k)
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