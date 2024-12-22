
type trieNode struct {
    childs [26]*trieNode
    isEnd bool
}

func newTrieNode() *trieNode{
    return &trieNode{childs:[26]*trieNode{}}
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
        idx := int(word[i] - 'a')
        if curr.childs[idx] == nil {curr.childs[idx] = newTrieNode()}
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}


func (this *Trie) Search(word string) bool {
    curr := this.root
    for i := 0; i < len(word); i++ {
        idx := int(word[i] - 'a')
        if curr.childs[idx] == nil {return false}
        curr = curr.childs[idx]
    }
    return curr.isEnd    
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := this.root
    for i := 0; i < len(prefix); i++ {
        idx := int(prefix[i] - 'a')
        if curr.childs[idx] == nil {return false}
        curr = curr.childs[idx]
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