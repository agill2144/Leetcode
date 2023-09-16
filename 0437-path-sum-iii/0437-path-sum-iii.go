/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: runningSum pattern
    - if we have a runningSum up to certain node
    - how do we have this runningSum == target ?
        - diff = runningSum - target
    - meaning, we need to remove $diff value from current runningSum to == target
    - we then need to check if we have this $diff as a runningSum before in this path ?
    - inorder to keep track of all old runningSum at each old node, we will track them in a map
    - Then we can check if $diff was seen before as a runningSum upto some old node?
        - if yes, it means
        - that from root to some old node the runningSum is $diff
        - if we remove this $diff from current runningSum, we will be left exactly with $target
        - [............]
                     x
                y
        - x is current running sum
        - y if some old running sum
        - if we remove y from x we exactly get $target
        - how do we know what y to search for ?
        - by asking the question; "How do make current $runningSum == target"
            - target = 10
            - if runningSum is 16, we need to remove 6 ; 16-10 = 6 ; then 16-6 = 10 ; so we need to search for a 6
            - if runningSum is 6 we need to remove -4 ; 6-10 = -4 ; then 6-(-4) = 10 ; so we need to search for a -4
            - search; hashmap!
            - why not a hashset ?
            - well turns out we need count number of paths
            - not just whether a path is possible or not
            - 6 or -4 could have been in our path several times,
            - which means there are several paths
            - therefore we need to count them all
            - therefore hashmap {$runningSum : $count}
    
    time = o(n)
    space = o(n)
*/

func pathSum(root *TreeNode, targetSum int) int {
    count := 0
    var dfs func(r *TreeNode, rSum int, path map[int]int)
    dfs = func(r *TreeNode, rSum int, path map[int]int) {
        // base
        if r == nil {return}
        
        // logic
        rSum += r.Val
        comp := rSum - targetSum
        val, ok := path[comp]
        if ok {
            count+= val
        }
        path[rSum]++
        
        dfs(r.Left, rSum, path)
        dfs(r.Right, rSum, path)
        path[rSum]--
    }
    dfs(root, 0, map[int]int{0:1})
    return count
}

/*
    approach: brute force
    - top down, make all possible paths from each node
    time = o(n^2)
    space = o(n^2)
*/
// func pathSum(root *TreeNode, targetSum int) int {
//     if root == nil {return 0}
//     count := 0
//     var dfs func(r *TreeNode, sum int)
//     dfs = func(r *TreeNode, sum int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         sum += r.Val
//         if sum == targetSum {
//             count++
//         }
//         dfs(r.Left, sum)
//         dfs(r.Right, sum)
//     }
    
//     var dfs2 func(r *TreeNode)
//     dfs2 = func(r *TreeNode) {
//         // base
//         if r == nil {return}
        
//         // logic
//         dfs(r, 0)
//         dfs2(r.Left)
//         dfs2(r.Right)
//     }
    
//     dfs2(root)
//     return count
// }

