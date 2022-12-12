/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: Brute force DFS, explore all paths
    - We can start from idx 0
    - suppose the value at idx 0 is N
    - Then we explore whether jumping 1..N is possible
    - if any one returns true, we will return true and stop exploring
    
    time: exponential ( n^N - i think, there are n elements each have their own distince N value )
    space: exponential - i think
    Result: TLE
    
**/
// func canJump(nums []int) bool {
//     var dfs func(startIdx int) bool
//     dfs = func(startIdx int) bool {
//         // base
//         if startIdx >= len(nums)-1 {
//             return true
//         }
        
//         // logic
//         for i := nums[startIdx]; i >= 1; i-- {
//             nextIdx := i + startIdx
//             if dfs(nextIdx) {
//                 return true
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }



/**
    Each value in the array list represents the jump distance we can jump ( 1 thru N where N is the ith val in array list)
    
    approach: DFS, explore all paths with Memoization
    - We can try caching in a map[idx]bool
        - why not an array ?
        - well we cannot safetly assume array size because a jump value could be 1k but array size is only 5.. and then it we go out of bounds..
        - therefore map
        
    - We can start from idx 0
    - suppose the value at idx 0 is N, then jump (0+N..1)
    - if this new idx position that we will jump to is already explored ( cached ) and if its true, then return true
    - otherwise recursively call the dfs function again and cache its value - if the value is true, return and exit early, otherwise continue exploring...
    
    time: exponential ( n^N - i think, there are n elements each have their own distince N value )
    space: exponential - i think
    Result: actually passes.... but still very slow
    
**/
// func canJump(nums []int) bool {
//     var dfs func(startIdx int) bool
//     memo := map[int]bool{}

//     dfs = func(startIdx int) bool {
//         // base
//         if startIdx >= len(nums)-1 {
//             return true
//         }
        
//         // logic
//         for i := nums[startIdx]; i >= 1; i-- {
//             nextIdx := i + startIdx
//             if res, ok := memo[nextIdx]; ok {
//                 if res {return true}
//             }else {
//                 res := dfs(nextIdx)
//                 memo[nextIdx] = res
//                 if res { return true }
//             }
//         }
//         return false
//     }
//     return dfs(0)
// }


/*
    approach: Brute force BFS 
    - We have a start position
    - We have a destination
    - We have to find whether if its even possible to reach to the destination
    - Enqueue our start idx 0
    - Then 0's childs are all the indicies 0th idx can reach to
        - i.e use val at idx 0 in a greedy manner and enqueue all the childs our jump would land to 
    - if the dq'd idx == destination, return true,
    - otherwise if queue did not find an answer, return false
*/
// func canJump(nums []int) bool {
//     visited := map[int]struct{}{0:struct{}{}}
//     q := []int{0}
//     n := len(nums)
    
//     for len(q) != 0 {
//         dqIdx := q[0]
//         q = q[1:]
//         if dqIdx >= n-1 {
//             return true
//         }
//         for i := nums[dqIdx]; i >= 1; i-- {
//             nextIdx := dqIdx+i
//             if nextIdx >= n-1 {return true}
//             _, alreadyEnqueued := visited[nextIdx]
//             if alreadyEnqueued {continue}
//             visited[nextIdx] = struct{}{}
//             q = append(q, nextIdx)
//         }
//     }
//     return false
// }



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
    space: o(1)
    Result: ofcourse this passes
    
**/
func canJump(nums []int) bool {
    destIdx := len(nums)-1
    for i := len(nums)-2; i >= 0; i-- {
        numJumpsAvail := nums[i]
        if i + numJumpsAvail >= destIdx {
            destIdx = i
        }
    }
    return destIdx == 0
}