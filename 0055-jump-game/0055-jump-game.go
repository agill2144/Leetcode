/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: Greedy, explore single path
    - Starting from the start of the list, we end up with multiple choices
        - How? if val at idx 0 is 2, then we can take 1 jump or 2 jumps
    - Instead to explore ONLY 1 path, we will
    - make the last element as our destination
    - and start our i from len(nums)-2 ( 2nd last element )
    - and check whether we can reach destination from ith element
    - if we can, then make the current i the new destination and move i ptr back by 1
    - then for instance, we cannot, move your i ptr back but leave the destination ptr where it is,
    - because it may not be reachable from where current i is , but maybe possible from i-1 element
    - if our destination reaches idx 0, we can gurantee a path with current number of jumps avail from start to end
    
    
    time: o(n)
    space: o(n)
*/
func canJump(nums []int) bool {
    n := len(nums)
    target := n-1
    for i := n-2; i >= 0; i-- {
        if i + nums[i] >= target {
            target = i
        }
    }
    return target == 0
}

/*
    approach: bottom up dp starting from the destination
    - dfs shows repeated subproblems
    - therefore dp is possible
    - for each problem, we want to see whether we can reach a position that is either the destination or a idx thats reachable to destination
    - we will create a dp bool array
    - initially, we will have the last position in dp array ( destination idx ) marked with true
    - then we will start from back and explore all jumps ( by first taking the max len jump and working back to jump size = 1)
    - if after jumping from ith position (i+j) is inbound and we have reached an idx where we can reach the destinaton from (dp[i+j] is true)
    - then dp[i] can be marked as true and we can break out of jump loop
    - then finally the answer of whether we can reach the last idx from 0th idx will be in dp[0]
*/
// func canJump(nums []int) bool {
//     dp := make([]bool, len(nums))
//     dp[len(dp)-1] = true
//     for i := len(dp)-2; i >= 0; i-- {
//         numJumps := nums[i]
//         for j := numJumps; j >= 1; j-- {
//             if i+j < len(nums) && dp[i+j] {
//                 dp[i] = true
//                 break
//             }
//         }
//     }
    
//     return dp[0]
// }

/*
    DFS decision trees shows repeated subproblems
    - top down DP
    - memoize idx whether its possible to reach destination from that idx
    - if an idx we are exploring, is not memoized, explore it, memoize its answer
    - finally return the answer from memoization map for each idx
*/
// func canJump(nums []int) bool {
//     memo := map[int]bool{}
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start >= len(nums)-1 {
//             if start == len(nums)-1{return true}
//             return false
//         }
        
//         // logic
//         if _, ok := memo[start]; !ok {
//             numJumps := nums[start]
//             for j := numJumps; j >= 1; j-- {
//                 if ok := dfs(start+j); ok {memo[start]=true; return true}
//             }
//             memo[start] = false
//         }
//         return memo[start]
//     }
//     return dfs(0)
// }


/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: DFS, explore all paths
    - We can start from idx 0
    - suppose the value at idx 0 is N
    - Then we explore whether jumping 1..N is possible
    - if any one returns true, we will return true and stop exploring
    
    time: exponential ( n^N - i think, there are n elements each have their own distince N value )
    space: exponential - i think
    Result: TLE
    
**/
// func canJump(nums []int) bool {
//     memo := map[int]bool{}
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start >= len(nums)-1 {
//             if start == len(nums)-1{return true}
//             return false
//         }
        
//         // logic
//         if _, ok := memo[start]; !ok {
//             numJumps := nums[start]
//             for j := numJumps; j >= 1; j-- {
//                 if ok := dfs(start+j); ok {memo[start] = true; return true}
//             }
//         }
//         return memo[start]
//     }
//     return dfs(0)
// }