type trieNode struct{
    isEnd bool
    childrens [26]*trieNode
}

type WordDictionary struct {
    root *trieNode    
}


func Constructor() WordDictionary {
    return WordDictionary{
        root: &trieNode{childrens: [26]*trieNode{}},
    }
}


func (this *WordDictionary) AddWord(word string)  {
    curr := this.root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &trieNode{childrens: [26]*trieNode{}}
        }
        curr = curr.childrens[charIdx]
    }
    curr.isEnd = true
}


func (this *WordDictionary) Search(word string) bool {
    return searchInTrie(this.root, word)
}


func searchInTrie(root *trieNode, word string) bool {
    curr := root
    for i := 0; i < len(word); i++ {
        if word[i] == '.'{
            for j := 0; j < len(curr.childrens); j++ {
                if curr.childrens[j] != nil {
                    if ok := searchInTrie(curr.childrens[j], word[i+1:]); ok {return true}
                }
            }
            return false
        } else {
            charIdx := word[i]-'a'
            if curr.childrens[charIdx] == nil {return false}
            curr = curr.childrens[charIdx]
        }
    }
    return curr.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */