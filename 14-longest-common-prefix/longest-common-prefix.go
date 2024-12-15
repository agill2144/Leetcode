/*
    approach: strings + prefix = trie
    - anytime we see strings and something related to prefix/suffix - think TRIE!
    - toss all words into a trie
        - with an additional field at each node
        - count, everytime a char is inserted, the count of that char goes up by 1
    - then perform a dfs to create and find the longest common prefix
        - loop over root childs ( 26 childs )
        - if a child exists and its COUNT == len(strs)
        - meaning at this level, we have seen len(strs) occurences of this char
        - take this char into our path, and recurse on this node ( this node becomes root for next recursion )
        - repeat ... 
        - keep updating our answer with our path
            - if the path we have formed so far is greater in size then the path in our answer, update our answer
    
    - wouldn't we run into cases where; [fl, flox, xflow ]
    - char f is repeated 3 times, so f is valid prefix
    - no, this wont happen
    - because at level 0 of the trie, we wont have f repeating 3 times, only 2 times
    - the f from the last word is at a level below 1st level
    - This means, count at a node represents how many times we have seen this char at this level
    - and it must be len(strs) - because we want to find common prefix that exists in ALL strs
*/


type trieNode struct {
    isEnd bool
    count int
    childs [26]*trieNode
}

func newTrieNode() *trieNode {
    return &trieNode{childs:[26]*trieNode{}}
}

func (r *trieNode) insert(word string) {
    curr := r
    for i := 0; i < len(word); i++ {
        idx := int(word[i]-'a')
        if curr.childs[idx] == nil {
            curr.childs[idx] = newTrieNode()
        }
        curr = curr.childs[idx]
        curr.count++
    }
    curr.isEnd = true
}


func longestCommonPrefix(strs []string) string {
    n := len(strs)
    if n <= 1 {if n == 1 {return strs[0]}; return ""}
    root := newTrieNode()
    for i := 0; i < len(strs); i++ {
        root.insert(strs[i])
    }
    ans := ""
    var dfs func(curr *trieNode, path string)
    dfs = func(curr *trieNode, path string) {
        // base
        if len(path) > len(ans) {ans = path}

        // logic
        for i := 0; i < len(curr.childs); i++ {
            if curr.childs[i] != nil && curr.childs[i].count == n {
                path += string(i+'a')
                dfs(curr.childs[i], path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(root, "")
    return ans
}
/*
    assume strs were sorted
    approach: 2 ptrs
    - if the list was sorted in asc order
    - the common prefixes will be grouped automatically
    - [flow, fl, flag, pet, peter, petrol]
    - but we want all prefixes to match
    - meaning all strs should have the same prefix
    - stack the sorted words on top of each other
    - and take 2 ptrs to compare char at strs[0] and strs[n-1] ; idx-by-idx
    - if an chars do not match in this sorted list
        - it means whatever we have matched so far is at strs[0][:ptr1]
    n = len(words)
    k = avg len of each word
    tc = o(nk logn) + o(k)
    sc = o(1) if not assuming sorting space
*/
// func longestCommonPrefix(strs []string) string {
//     n := len(strs)
//     if n <= 1 {if n == 0 {return ""}; return strs[0]}
//     sort.Strings(strs)
//     p1, p2 := 0, 0
//     for p1 < len(strs[0]) && p2 < len(strs[n-1]) {
//         if strs[0][p1] != strs[n-1][p2] {return strs[0][:p1]}
//         p1++
//         p2++
//     }
//     return strs[0]
// }

/*
    approach: stack words on top and perform char scan idx-by-idx
    - instead of creating $evalPrefix
    - and then comparing with rest of words,
    - we can just compare whether ith idx char in reserved word 
        matches ith idx in rest of the strings
    - so this still means, reserver a word first
        - strs[0]
    - loop over this reserved word, char by char - ith idx
    - then loop over remaining strs from index 1
    - and check if ith idx char == reserved char ith idx
    - if it does not equal, ith idx was tracking the common prefix so far
        - so return reservedWord[:i]
        - substr up-to but not including i
    S = total number of character in all strs
    tc = o(S)
    sc = o(1)
*/
// func longestCommonPrefix(strs []string) string {
//     n := len(strs)
//     if n <= 1 {
//         if n == 1 {return strs[0]}
//         return ""
//     }
//     reserved := strs[0]
//     for i := 0; i < len(reserved); i++ {
//         for j := 1; j < len(strs); j++ {
//             word := strs[j]
//             if i == len(word) || word[i] != reserved[i] {return reserved[:i]}
//         }
//     }
//     return reserved
// }
/*
    approach: brute force
    - create all possible substr from strs[0]; starting from size 1, then 2, then 3
        - $evalPrefix
    - and check if this substr as a prefix in rest of the strings
        - that is, start from idx 1
        - check if strs[j] has a prefix of $evalPrefix
        - if it does, increment a counter and if count reaches n-1
            - then is a good answer, save it
            - and try a longer $evalPrefix
            - because we want longest common
    n = len(strs)
    k = avg size of each strs[i]
    tc = o(k * n*k) = o(n*k^2)
    sc = o(1)
*/
// func longestCommonPrefix(strs []string) string {
//     n := len(strs)
//     if n <= 1 {
//         if n == 0 {return ""}
//         return strs[0]
//     }
//     word := strs[0]
//     // loop over reserved word
//     // to form substrs
//     ans := ""
//     for i := 0; i < len(word); i++ {
//         prefix := word[:i+1]
//         count := 0
//         for j := 1; j < len(strs); j++ {
//             if strings.HasPrefix(strs[j], prefix) {
//                 count++
//                 if count == n-1 {
//                     ans = prefix
//                     break
//                 }
//             }
//         }
//     }
//     return ans
// }