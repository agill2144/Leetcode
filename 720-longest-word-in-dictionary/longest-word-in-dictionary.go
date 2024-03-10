func longestWord(words []string) string {
    root := &trieNode{childrens: [26]*trieNode{}}
    for i := 0; i < len(words); i++ {
        insert(root,words[i])
    }
    // we could either build the word as we are going in recursion "path"
    // or maintain the word in each trieNode and use that later

    // maintaining in each trieNode takes more space then maintaining a path in recursion
    // and backtracking :shrug
    var dfs func(r *trieNode) string
    dfs = func(r *trieNode) string {
        // base
        
        // logic
        largest := ""
        if r.isEnd && len(r.word) > len(largest) {largest = r.word}
        for i := 0; i < len(r.childrens); i++ {
            if r.childrens[i] != nil && r.childrens[i].isEnd {
                val := dfs(r.childrens[i])
                if len(val) > len(largest) {
                    largest = val
                }
            }
        }
        return largest
    }
    return dfs(root)
}


type trieNode struct {
    isEnd bool
    childrens [26]*trieNode
    word string
}


func insert(root *trieNode, word string)  {
    curr := root
    for i := 0; i < len(word); i++ {
        charIdx := int(word[i]-'a')
        if curr.childrens[charIdx] == nil {
            curr.childrens[charIdx] = &trieNode{childrens: [26]*trieNode{}}
        }
        curr = curr.childrens[charIdx]
    }
    curr.isEnd = true
    curr.word = word
}




// func longestWord(words []string) string {
//     set := map[string]struct{}{}
//     for i := 0; i < len(words); i++ {
//         set[words[i]] = struct{}{}
//     }
//     ans := ""
//     // time = o(n) ; n total words
//     // total time = o(n)*(o(k)*o(k)); o(n*k^2)
//     for i := 0; i < len(words); i++ {
//         // avg size of each word is k
//         word := words[i]
//         isValidWord := false
//         // exlcude the last char each time and check ALL prefixes
//         // all prefixes must exist for this word to be a valid word
//         // time = o(k)
//         for len(word) > 0 {
//             // time = o(k)
//             prefix := word[:len(word)-1]
//             if len(prefix) == 0 {isValidWord = true; break}
//             _, ok := set[prefix]
//             if ok {
//                 word = prefix
//                 continue
//             }
//             break
//         }
//         // if the word is valid ( i.e all prefixes exist )
//         if isValidWord {
//             // same len as ans word, then smallest lexiographically wins
//             if len(words[i]) == len(ans) {
//                 if strings.Compare(words[i], ans) == -1 {
//                     ans = words[i]
//                 }
//             // otherwise if this new word size > ans word, save it
//             } else if len(words[i]) > len(ans) {
//                 ans = words[i]
//             }
//         }
//     }
//     return ans
// }

