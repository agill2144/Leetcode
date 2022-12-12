/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: brute force DFS
    - maintain a list of paths till current node as an arg
        - path is a list of ints from root to current node (excluding current for comparison )
    - For curr node, compare with all nodes in path
    - as soon as curr node < path val, this node is not considered a good node
    - if curr node goes thru all the nodes in paths and kept returning true : currNode.Val >= nodeValFromBackOfPaths, increment counter by 1
    - Finally, regardless whether we incremented the counter or not, add the current node val to paths list and recursively call dfs on left and right childs.
    
    approach: DFS with max as an arg
    - maintain a "maxSeenSoFar" as an arg ( starting with root node val )
    - For curr node, compare with max -- if currNode >= maxSeenSoFar {count++} -- this is good node
    - Update max -- math.max(max, currNode.Val)
    - Finally, regardless whether we incremented the counter or not, recursively call dfs on left and right childs.
    
    
    
    Time: o(n)
    space: o(h)

*/
func goodNodes(root *TreeNode) int {
    result := 0
    var dfs func(max int, a *TreeNode)
    dfs = func(max int, a *TreeNode) {
        // base
        if a == nil {return}
        
        // logic
        if a.Val >= max {
            max = a.Val
            result++
        }
        dfs(max, a.Left)
        dfs(max, a.Right)
    }
    dfs(root.Val, root)
    return result
}