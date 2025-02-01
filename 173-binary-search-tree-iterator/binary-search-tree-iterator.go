/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    nums []int
    ptr int
}


func Constructor(root *TreeNode) BSTIterator {
    nums := []int{}    
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        nums = append(nums, r.Val)
        dfs(r.Right)
    }
    dfs(root)
    return BSTIterator{nums,0}
}


func (this *BSTIterator) Next() int {
    n := len(this.nums)
    if this.ptr == n {return -1}
    out := this.nums[this.ptr]
    this.ptr++
    return out
}


func (this *BSTIterator) HasNext() bool {
    n := len(this.nums)
    return this.ptr < n
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */