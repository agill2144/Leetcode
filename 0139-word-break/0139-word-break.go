/*
    we have to split s into substrs such that each splitted substr exists in wordDict
    approach: bottom up dp
    - we will create a 1d dp array of size s+1 ( bool array )
    - we will look over the dp array from left to right ( i loop )
    - and at each position, we will answer;
        - whether the substr from 0:i-1 is a valid segment / split or not
        - whether the substr before the ith char is a valid substr or not
    - what about all other substrings we explored when doing recursively ?
    - we will be doing the same thing except if at a ith position in dp array is false
        - this means the previous substr cannot be split into valid word
        - if previous cannot be split, then there is no point in further exploring next substr after previous split
        - this is where we save extra recursive calls
        - but it maybe possible that previous substr is not possible but maybe some other substr in the middle is possible.
        - this is where we will have a j pointer starting from idx 0
        - j loop runs from 0 to i ( excluding i )
        - if value at dp[j] is true
            - it means the previous string is a valid split
            - it maybe worth further exploring down this path
            - i.e substr[j:i]
            - if the above at point is true at any point, it means we have a valid split.
            - set dp[i] = true, break out of the j loop
    
    n = len(s)
    time: o(len(wordDict)) + o(n^3)  how? o(n) for the s string * o(n) for the j loop, * o(n) for substr creation
    space: o(len(wordDict)) + o(n+1)
*/

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
                    // follow up, return the exact subStr splits
                    // append the subStr we have formed here into an output array
                    // fmt.Println(subStr)
                    break
                }
            }
        }
    }
    return dp[len(dp)-1]
}

/*    
    
    approach: bottom up DP ( splitting substrs from back )
    - back partitioning only if j-1 is a valid split
    - start splitting from the back
    - continue following the model of bottom up
    - start from the back and make decisions; 
    - and that is subStr[j:i] can be split only if dp[j-1] is a valid split
*/

// func wordBreak(s string, wordDict []string) bool {
//     set := map[string]struct{}{}
//     for i := 0; i < len(wordDict); i++ { set[wordDict[i]] = struct{}{} }
//     dp := make([]bool, len(s))
//     for i := 0; i < len(s); i++ {
//         for j := i; j >= 0; j-- {
//             if j == 0 || dp[j-1] {
//                 subStr := string(s[j:i+1])
//                 if _, ok := set[subStr]; ok {
//                     dp[i] = true
//                     break
//                 }
//             }
//         }
//     }
//     return dp[len(dp)-1]
// }


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