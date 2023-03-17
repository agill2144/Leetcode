type Trie struct {
    isEnd bool
    childrens [26]*Trie
}


func Constructor() Trie {
    return Trie{
        childrens: [26]*Trie{},
    }
}


func (this *Trie) Insert(word string)  {
    curr := this
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &Trie{childrens: [26]*Trie{}}
        }
        curr = curr.childrens[charIdx]
    }
    curr.isEnd = true
}


func (this *Trie) Search(word string) bool {
    curr := this
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childrens[charIdx] == nil {return false}
        curr = curr.childrens[charIdx]
    }
    return curr.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := this
    for i := 0; i < len(prefix); i++ {
        charIdx := prefix[i]-'a'
        if curr.childrens[charIdx] == nil {return false}
        curr = curr.childrens[charIdx]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */