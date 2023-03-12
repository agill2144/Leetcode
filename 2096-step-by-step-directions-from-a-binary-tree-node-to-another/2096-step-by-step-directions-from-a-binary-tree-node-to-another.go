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
    var dfs func(r *TreeNode, start, end int, path []string) 
    dfs = func(r *TreeNode, start, end int, path []string) {
        // base
        if r == nil {return}
        if startPath != nil && endPath != nil {return}
        
        // logic
        if r.Val == start || r.Val == end {
            newL := make([]string, len(path))
            copy(newL, path)            
            // newL = append(newL, "T")
            if r.Val == start {
                startPath = newL
            } else {
                endPath = newL
            } 
        }
        if startPath != nil && endPath != nil {return}

        path = append(path,"L")
        dfs(r.Left, start, end, path)
        path = path[:len(path)-1]

        path = append(path, "R")
        dfs(r.Right, start, end, path)
        path = path[:len(path)-1]
        
    }
    dfs(lca, startValue, destValue, nil)
    
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


