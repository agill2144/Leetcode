/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDirections(root *TreeNode, startValue int, destValue int) string {
    lca := lowestCommonAncestor(root, startValue, destValue)
    var startPath []string
    var endPath []string
    var dfs func(r *TreeNode, target int, path []string) []string 
    dfs = func(r *TreeNode, target int, path []string) []string {
        // base
        if r == nil {return nil}
        
        // logic
        if r.Val == target {
            newL := make([]string, len(path))
            copy(newL, path)
            return newL
        }

        path = append(path,"L")
        if ok := dfs(r.Left, target, path); ok != nil {return ok}
        path = path[:len(path)-1]

        path = append(path, "R")
        if ok := dfs(r.Right, target, path); ok != nil {return ok}
        path = path[:len(path)-1]
        
        return nil
    }
    startPath = dfs(lca, startValue, nil)
    endPath = dfs(lca, destValue, nil)
    // start
    out := ""
    for i := len(startPath)-1; i>=0; i-- {
        if startPath[i] == "T" {continue}
        out += "U"
    }
    for i := 0; i < len(endPath); i++ {
        if endPath[i] == "T" {continue}
        out += endPath[i]
    }
    return out
}




func lowestCommonAncestor(root *TreeNode, p, q int) *TreeNode {
    var dfs func(r *TreeNode) *TreeNode 
    dfs = func(r *TreeNode) *TreeNode {
        // base
        if r == nil {return nil}
        if r.Val == p || r.Val == q {return r}

        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)

        if left != nil && right != nil {
            return r
        }

        if left == nil && right == nil {return nil}
        if left != nil && right == nil {return left}
        if left == nil && right != nil {return right}
        return nil
    }
    return dfs(root)
}


