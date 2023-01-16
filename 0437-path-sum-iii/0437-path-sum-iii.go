/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: running/prefix sum + backtracking
    - this is like subarray sum that equals target with numbers that contain negatives
    - where we used runningSum pattern.
    - The runningSum pattern was we searched for a complement
    - x + y = z
    - if have x and z, do get y, we can do z-x
    - [1,2,3,4,5,6,7,8]
               x     z
                  y find sum subarray ( y ) - which is subarray between x and z
        - x is the running sum up until x
        - z is the running sum up until z
        - if we have z and a x, then we can remove x from z and get the y value
            - therefore y = z-x
    - But we have z and y and we are looking for an x.
    - if we find an x than that means we have a perfect solution for x+y = z
    - how do we search for x ?
        - search! hash map!
        - are we looking for positions or count?
        - count, so freq map
    - What are searching for when we are searching for x
    - x+y=z
    - we have y and z
    - to find x = z-y
    
    That same subarray sum idea can be applied to trees.
    - We dont have a subarray so we need something to keep track of a path going down ( i.e our own subarray )
    - but since numbers in a path going down can repeat, we will have to use a hashmap ( freqMap )
    - once our recursion comes back from a node ( i.e done with a node, ) than this node is no longer being considered as part of our "subarray" - therefore remove it from freqMap (decrement first, and then remove if needed )
        - i.e backtrack our map/path/subarray
    
    time: o(n)
    space: o(n) + o(h) ; or worse case its just o(n)
*/
func pathSum(root *TreeNode, targetSum int) int {
    count := 0
    var dfs func(r *TreeNode, rSum int, path map[int]int)
    dfs = func(r *TreeNode, rSum int, path map[int]int) {
        // base
        if r == nil {return}
        
        // logic
        rSum += r.Val
        complement := rSum-targetSum
        val, ok := path[complement]
        if ok {count += val}
        path[rSum]++
        
        dfs(r.Left, rSum, path)
        dfs(r.Right, rSum, path)
        
        path[rSum]--
        if val, _ := path[rSum]; val == 0 {delete(path, rSum)}
        
    }
    dfs(root,0,map[int]int{0:1})
    return count
}