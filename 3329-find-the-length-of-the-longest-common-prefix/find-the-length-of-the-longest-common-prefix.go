func longestCommonPrefix(arr1 []int, arr2 []int) int {
    if len(arr1) == 0 || len(arr2) == 0 {return 0}
    root := newTrie()
    // a1 = len(arr1); k1 = avg num of digits in arr1
    // a2 = len(arr2); k2 = avg num of digits in arr2

    // o(a2 *  k2)
    for i := 0; i < len(arr2); i++ {
        root.insert(arr2[i])
    }

    largest := 0
    // o(a1 * k1)
    for i := 0; i < len(arr1); i++ {
        arr1Str := fmt.Sprintf("%v", arr1[i])
        if len(arr1Str) < largest {continue}
        // use trie to find the longest prefix possible for this num
        val := root.longestPossiblePrefix(arr1Str)
        largest = max(largest, val)
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

func (r *trieNode) longestPossiblePrefix(word string) int {
    curr := r
    size := 0
    for i := 0; i < len(word); i++ {
        childIdx := int(word[i]-'0')
        if curr.childs[childIdx] == nil {break}
        curr = curr.childs[childIdx]
        size++
    }
    return size
}