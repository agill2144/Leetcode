func isPrefixOfWord(sentence string, searchWord string) int {
    root := &trieNode{childs: [26]*trieNode{}}
    
    wordCount := 1
    tmp := new(strings.Builder)
    for i := 0; i < len(sentence); i++ {
        if sentence[i] >= 'a' && sentence[i] <= 'z' {
            tmp.WriteByte(sentence[i])
        }
        if sentence[i] == ' ' || i == len(sentence)-1 {
            tmpStr := tmp.String()
            if len(tmpStr) > 0 {
                root.insert(tmpStr, wordCount)
            }
            tmp.Reset()
            wordCount++
        }
    }

    return root.searchPrefix(searchWord)

}



type trieNode struct {
	isEnd  bool
	childs [26]*trieNode
    idx int
}

func (r *trieNode) insert(word string, startIdx int) {
	curr := r
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if curr.childs[idx] == nil {
			curr.childs[idx] = &trieNode{childs: [26]*trieNode{}, idx: startIdx}
		}
		curr = curr.childs[idx]
	}
	curr.isEnd = true
}


func (r *trieNode) searchPrefix(prefix string) int {
	curr := r
	for i := 0; i < len(prefix); i++ {
		idx := int(prefix[i] - 'a')
		if curr.childs[idx] == nil {
			return -1
		}
		curr = curr.childs[idx]
	}
	return curr.idx
}
