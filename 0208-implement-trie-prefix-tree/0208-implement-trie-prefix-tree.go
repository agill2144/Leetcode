type trieNode struct{
    isEnd bool
    childs [26]*trieNode
}
func newTrieNode() *trieNode {
    return &trieNode{childs: [26]*trieNode{}}
}


type Trie struct {
    root *trieNode
}


func Constructor() Trie {
    return Trie{
        root: newTrieNode(),
    }
}


func (this *Trie) Insert(word string)  {
    curr := this.root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childs[charIdx] == nil {
            curr.childs[charIdx] = newTrieNode()
        }
        curr = curr.childs[charIdx]
    }
    curr.isEnd = true
}


func (this *Trie) Search(word string) bool {    
    curr := this.root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childs[charIdx] == nil { return false }
        curr = curr.childs[charIdx]
    }
    return curr.isEnd
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := this.root
    for i := 0; i < len(prefix); i++ {
        charIdx := prefix[i]-'a'
        if curr.childs[charIdx] == nil { return false }
        curr = curr.childs[charIdx]
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