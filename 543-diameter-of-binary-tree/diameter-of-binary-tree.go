/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // naive , top down dfs
 // get max left and right depth from each root node
 // and max dia from that root node will always be leftDepth + rightDepth
 // but its possible that the maxDepth is not always at the main root
 // consider a child of root has 5 nodes on left and 10 nodes on right
 // but when processing for such root , we will get 10 back as max depth
 // but when processing the child node as the root, we will get back 5+10 = 15
 // hence for each node, we need its max left depth and its max right depth
 // tc = o(n^2)
 // sc = o(n)

 // whenever top down recursion has a n^2 time, consider bottom up
 // can the child tell its parent its max depth
 // the leaf depth will always be 1
 // then a parent right above leaf will get a depth from left and right
 // and then we can calculate the max diameter ( left + right ) and save
 // but now this parent needs to return its max depth, its either left or right
 // we want this nodes parent to get a max possible depth , therefore return max(left,right) + 1
 // the +1 is adding the edge between current root node and its parent node
 // so when parent gets a depth from its child, it has the edge from itself to child also added
 // now when a node from the bottom gets 5 as max left depth and 10 as max right depth
 // we can save diameter as 15 ( 10 + 5 )
 // and return back the max depth path to its parent
 // parent may not result into a longer diameter, but thats okay, because we already saved the max diameter.
 // by doing bottom up recursion, we are always considering the max diameter across each root node ( a path from that goes from left child -> root -> right child )
 // then we return back the longest path back to its parent to do the same
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {return 0}
    maxDia := 0
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        left := dfs(r.Left)
        right := dfs(r.Right)
        maxDia = max(maxDia, left+right)
        return max(left, right)+1
    }
    dfs(root)
    return maxDia
}
