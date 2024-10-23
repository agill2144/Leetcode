/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    if root == nil {return nil}
    levelSum := map[int]int{} // {level : $sum }
    var levels func(r *TreeNode, level int)
    levels = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        // logic
        levelSum[level] += r.Val
        levels(r.Left, level+1)
        levels(r.Right, level+1)
    }
    levels(root, 0)

    // stand at a parent
    // get child sum ( i.e sibling sum )
    // replace child val with = levelSum - childSum
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}

        // logic
        nextLevelSum, ok := levelSum[level+1]
        if !ok {return}
        siblingSum := 0
        if r.Left != nil {siblingSum += r.Left.Val}
        if r.Right != nil {siblingSum += r.Right.Val}
        newChildVal := nextLevelSum-siblingSum
        if r.Left != nil {
            r.Left.Val = newChildVal
            dfs(r.Left, level+1)
        }
        if r.Right != nil {
            r.Right.Val = newChildVal
            dfs(r.Right, level+1)
        }
    }
    dfs(root, 0)
    root.Val = 0
    return root

}