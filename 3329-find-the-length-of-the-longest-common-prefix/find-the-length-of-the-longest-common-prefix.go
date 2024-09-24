func longestCommonPrefix(arr1 []int, arr2 []int) int {
    if len(arr1) == 0 || len(arr2) == 0 {return 0}
    root := newTrie()
    for i := 0; i < len(arr2); i++ {
        root.insert(arr2[i])
    }

    // start with the longest one first!
    sort.Ints(arr1)
    largest := 0
    for i := len(arr1)-1; i >= 0; i-- {
        arr1Str := fmt.Sprintf("%v", arr1[i])
        if len(arr1Str) == largest {break}
        for j := len(arr1Str)-1; j >= 0; j-- {
            subStr := arr1Str[:j+1]
            if len(subStr) < largest {break}
            if root.prefixExists(subStr) {
                largest = max(largest, len(subStr))
            }
        }
        
    }
    return largest
}

type trieNode struct {
    isEnd bool
    childs []*trieNode
}

func newTrie() *trieNode {
    root := &trieNode{childs: make([]*trieNode, 10)}
    return root
}

func (r *trieNode) insert(n int) {
    word := fmt.Sprintf("%v", n)
    curr := r
    for i := 0; i < len(word); i++ {
        charIdx := int(word[i]-'0')
        if curr.childs[charIdx] == nil {
            curr.childs[charIdx] = &trieNode{childs: make([]*trieNode, 10)}
        }
        curr = curr.childs[charIdx]
    }
    curr.isEnd = true
}

func (r *trieNode) prefixExists (word string) bool {
    curr := r
    for i := 0; i < len(word); i++ {
        charIdx := int(word[i]-'0')
        if curr.childs[charIdx] == nil {return false}
        curr = curr.childs[charIdx]
    }
    return true
}