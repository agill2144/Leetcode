/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// use height property of a complete full binary tree
// num of nodes in a complete full binary tree = 2^h - 1
// h = height of this tree
// however this 2^h-1 is only true when the tree is full and complete
// i.e left height == right height
// if its unbalanced, recurse and check left and right subtree + add current node
func countNodes(root *TreeNode) int {
    if root == nil {return 0}
    getH := func(r *TreeNode) (int, int) { 
        l := r
        lh := 1
        for l.Left != nil {
            l = l.Left
            lh++
        }
        rh := 1
        for r.Right != nil {
            r = r.Right
            rh++
        }
        return lh, rh
    }
    count := 0
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        lh, rh := getH(r)
        if lh == rh {
            count += (int(math.Pow(2.0, float64(lh))) - 1)
        } else {
            // count current node
            count++
            // + count nodes of left subtree
            dfs(r.Left)
            // + count nodes of right subtree
            dfs(r.Right)
        }
    }
    dfs(root)
    return count
}

// brute force, count using any traversal
// func countNodes(root *TreeNode) int {
//     count := 0
//     var dfs func(r *TreeNode)
//     dfs = func(r *TreeNode) {
//         // base
//         if r == nil {return}
//         // logic
//         count++
//         dfs(r.Left)
//         dfs(r.Right)
//     }
//     dfs(root)
//     return count
// }
