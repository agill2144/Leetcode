func wordBreak(s string, wordDict []string) bool {
    set := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ { set[wordDict[i]] = struct{}{} }
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 0; i < len(dp); i++ {
        for j := 0; j < i; j++ {
            if dp[j] {
                subStr := string(s[j:i])
                if _, ok := set[subStr]; ok {
                    dp[i] = true
                    break
                }
            }
        }
    }
    return dp[len(dp)-1]
}

/*
    we have to split s into substrs such that each splitted substr exists in wordDict
    
    approach: brute force 
    - for loop based dfs
    - explore all possible splits
    - using a start / pivot ptr, make a substr
    - once a substr is found, then we know for sure this is 1 valid split
    - now go explore the rest of the string ( start = i + 1 ); recursively
    - once we put pivot out of bounds, it means that we were able to find substrings that exists in wordDict.
        - return true
    - otherwise return false
    s = len(s)
    
    to make wordDict searchable; we have 2 choices
    1. hashset 
    2. trie ( saves more memory compared to hashset because it uses same characters for repeating characters in a path )


    why recursion? why cant we just use fast and slow ptr
    - make each substr [slow:fast+1]
    - if substr exists, move slow ptr to fast+1
    - if slow ptr eventually reaches the end and goes out of bounds, it means we have all substrs that exists in wordDict.
    - However this does not for all test cases
    - s = "aaaaaaa", wordDict = ["aaaa","aaa"]
        - we will find "aaa" first
        - then we will another "aaa"
        - then we wont be able to find the last "a"
        - because the above didnt work, we have to be exhaustive and try other substrs
        - "aaaa" can be formed at the end
        - therefore inorder to be exhaustive, we have to use a for loop based recursion and try all possible sub-strings
*/
// func wordBreak(s string, wordDict []string) bool {
//     set := map[string]struct{}{}
//     for i := 0; i < len(wordDict); i++ { set[wordDict[i]] = struct{}{} }
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start == len(s) {return true}
        
//         // logic
//         for i := start; i < len(s); i++ {
//             subStr := string(s[start:i+1])
//             if _, ok := set[subStr]; ok {
//                 if yes := dfs(i+1); yes {
//                     return true
//                 }
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }