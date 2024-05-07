type StreamChecker struct {
    dq []byte
    maxLen int
    root *trieNode
}


func Constructor(words []string) StreamChecker {
    root := &trieNode{
        childs: [26]*trieNode{},
    }
    maxL := 0
    for i := 0; i < len(words); i++ {
        insert(root, words[i])
        maxL = max(maxL, len(words[i]))
    }
    return StreamChecker{
        dq: []byte{},
        maxLen: maxL,
        root: root,
    }
}


func (this *StreamChecker) Query(letter byte) bool {
    this.dq = append(this.dq, letter)
    if len(this.dq) > this.maxLen {
        this.dq = this.dq[1:]
    }
    curr := this.root
    for i := len(this.dq)-1; i >= 0; i-- {
        // search
        char := this.dq[i]
        charIdx := int(char-'a')
        if curr.childs[charIdx] == nil {return false}
        curr = curr.childs[charIdx]
        if curr.isStart {return true}
    }
    return false
}


type trieNode struct {
	isStart  bool
	childs [26]*trieNode
}

func insert(root *trieNode, word string) {
	curr := root
	for i := len(word)-1; i >= 0; i-- {
		idx := int(word[i] - 'a')
		if curr.childs[idx] == nil {
			curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
		}
		curr = curr.childs[idx]
	}
	curr.isStart = true
}
