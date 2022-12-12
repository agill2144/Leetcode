type TrieNode struct {
    // to denote that this word ends here ( full valid word )
    isEnd bool
    // optional if we want to search for words that begin with this trie node as prefix
    words []string
    
    // why 26 ? because
    // 1. trie nodes are used to store characters of a string ( 1 char == 1 trie node )
    // 2. And we have 26 total characters in english ( 26 lower case )
    // if we are supposed to handle lower and upper case, than the size would 26(lower-case)+26(upper-case)
    // if we are supposed to handle lower, upper and all other ascii values, then 256 would be children array length
    childrens [26]*TrieNode
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
        root: &TrieNode{childrens: [26]*TrieNode{}},
    }
}

// time: o(k) where k is the length of word
// space: o(k)
func (this *Trie) Insert(word string)  {
    curr := this.root
    for _, char := range word {
        if curr.childrens[char-'a'] == nil {
            curr.childrens[char-'a'] = &TrieNode{isEnd: false, childrens: [26]*TrieNode{}}
        }
        // optionally if we are told to find all words with $somePrefix
        curr.words = append(curr.words, word)
        curr = curr.childrens[char-'a']
    }
    // once we have inserted all characters of word, our curr ptr will be pointing to the last TrieNode of this word
    // we need to denote that this is the end of this word by setting isEnd=true at the last TrieNode of $word ( which is where curr will be pointing to)
    curr.isEnd = true
}

// time: o(k) where k is the length of word
// space: o(1)
func (this *Trie) Search(word string) bool {
    // we are looking for FULL VALID WORD ( not prefix )
    curr := this.root
    for _, char := range word {
        // get the idx of this char and if that idx is nil - this means the char does not exist which also means the word does not exist
        if curr.childrens[char-'a'] == nil {return false}
        // or else move current ptr to that node, and continue searching for the next character in $word
        curr = curr.childrens[char-'a']
    }
    
    // once we get here, it means we were will able to find ALL characters of $word
    // but that does not mean that this is a valid word ( it could be a prefix to some other longer word )
    // if this is a valid word that exists in our trie, then curr ( which points to the last trieNode char of $word ), its .isEnd property should be true 
    return curr.isEnd
}

// time: o(k) where k is the length of word
// space: o(1)
func (this *Trie) StartsWith(prefix string) bool {
    // here we are searching for whether we have a prefix 
    // existance of a prefix and not all words that began with this $prefix
    // so move curr ptr thru all chars of $prefix - if successful, it means we do have a prefix
    curr := this.root
    for _, char := range prefix {
        if curr.childrens[char-'a'] == nil {return false}
        curr = curr.childrens[char-'a']
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