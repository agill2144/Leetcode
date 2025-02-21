/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    set map[int]bool
}


func Constructor(root *TreeNode) FindElements {
    out := FindElements{set:map[int]bool{}}
    var dfs func(r *TreeNode, idx int)
    dfs = func(r *TreeNode, idx int) {
        // base
        if r == nil {return}
        // logic
        r.Val = idx
        out.set[r.Val] = true
        dfs(r.Left, 2*idx+1)
        dfs(r.Right, 2*idx+2)
    }
    dfs(root,0)
    return out
}


func (this *FindElements) Find(target int) bool {
    return this.set[target]
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */