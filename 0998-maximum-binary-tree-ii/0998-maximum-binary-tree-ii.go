/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
    path := []int{}
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        dfs(r.Left)
        path = append(path, r.Val)
        dfs(r.Right)
    }
    dfs(root)
    path = append(path, val)
    
    var create func(left, right int) *TreeNode
    create = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        
        
        // logic
        // find the max in this window
        maxVal, idx := path[left], left
        i := left+1
        for i <= right {
            if path[i] > maxVal {
                maxVal = path[i]
                idx = i 
            }
            i++
        }
        root := &TreeNode{Val: maxVal}
        root.Left = create(left, idx-1)
        root.Right = create(idx+1, right)
        return root
    }
    return create(0, len(path)-1)

    
}