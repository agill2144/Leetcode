
type trieNode struct{
    isEnd bool
    childs [26]*trieNode
    prefixCount int
    count int
}
func newTrieNode() *trieNode {
    return &trieNode{childs: [26]*trieNode{}}
}


type Trie struct {
    root *trieNode    
}


func Constructor() Trie {
    return Trie{root: newTrieNode()}   
}


func (this *Trie) Insert(word string)  {
        curr := this.root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childs[charIdx] == nil {
            curr.childs[charIdx] = newTrieNode()
        }
        curr = curr.childs[charIdx]
        curr.prefixCount++
    }
    curr.isEnd = true
    curr.count++

}


func (this *Trie) CountWordsEqualTo(word string) int {
    curr := this.root
    for i := 0; i < len(word); i++ {
        charIdx := word[i]-'a'
        if curr.childs[charIdx] == nil { return 0 }
        curr = curr.childs[charIdx]
    }
    if curr.isEnd {return curr.count}
    return 0
}


func (this *Trie) CountWordsStartingWith(prefix string) int {
    curr := this.root
    for i := 0; i < len(prefix); i++ {
        charIdx := prefix[i]-'a'
        if curr.childs[charIdx] == nil { return 0 }
        curr = curr.childs[charIdx]
    }
    return curr.prefixCount
}


func (this *Trie) Erase(word string)  {
    curr := this.root
    wordPtr := 0
    var dfs func(r *trieNode) bool
    dfs = func(r *trieNode) bool {
        // base
        if wordPtr == len(word) {
            if !r.isEnd {
                return false
            }
            r.prefixCount--
            r.count--
            
            return true
        }
        
        // logic
        charIdx := word[wordPtr]-'a'
        wordPtr++
        if dfs(r.childs[charIdx]) {
            r.prefixCount--
            // parent will delete child node if both count and prefixCount of a child 
            // has reached 0
            if r.childs[charIdx].count == 0 && r.childs[charIdx].prefixCount == 0 {
                r.childs[charIdx] = nil
            }
            return true
        }
        return false
    }
    dfs(curr)
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */