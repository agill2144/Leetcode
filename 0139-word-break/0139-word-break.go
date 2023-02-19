
/*
    approach: bottom up dp
    - at any char, find out whether the string before this char can be split into valid splits or not
*/
func wordBreak(s string, wordDict []string) bool {
    dp := make([]bool, len(s)+1)
    dp[0] = true
    root := &trieNode{childs: [26]*trieNode{}}
    for i := 0; i < len(wordDict); i++ { insertWord(root, wordDict[i])}
    
    for i := 1; i < len(dp); i++ {
        for j := 0; j < i; j++ {
            if dp[j] {
                subStr := string(s[j:i])
                if search(root, subStr) {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[len(dp)-1]
}



/*
    approach: brute force
    - find a partition path that works
    - using for loop recursion and pivot, form each substr
    - if a substr exists in wordDict, recurse with the remaining words
    - keep going until our pivot goes out of bounds
    - when pivot goes out of bounds, it means we were able to find substrs that do exists in the wordDict
        - return true
    - otherwise at the end return false
*/
// func wordBreak(s string, wordDict []string) bool {
//     root := &trieNode{childs: [26]*trieNode{}}
//     for i := 0; i < len(wordDict); i++ { insertWord(root, wordDict[i])}
    
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start == len(s) {return true}
        
//         // logic
//         for i := start; i < len(s); i++ {
//             subStr := string(s[start:i+1])
//             if search(root, subStr) {
//                 if yes := dfs(i+1); yes {return true}
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }


type trieNode struct {
    childs [26]*trieNode
    isEnd bool
}

func insertWord(root *trieNode, word string)  {
    curr := root
    for i := 0; i < len(word); i++ {
        idx := word[i]-'a'
        if curr.childs[idx] == nil {
            curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
        }
        curr = curr.childs[idx]
    }
    curr.isEnd = true
}

func search(r *trieNode, word string) bool {
    curr := r
    for i := 0; i < len(word); i++ {
        idx := word[i]-'a'
        if curr.childs[idx] == nil { return false }
        curr = curr.childs[idx]
    }
    return curr.isEnd
}