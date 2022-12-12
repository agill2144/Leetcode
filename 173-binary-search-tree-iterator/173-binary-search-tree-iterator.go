/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    stack []*TreeNode
    
}


func Constructor(root *TreeNode) BSTIterator {
    b := BSTIterator{stack: []*TreeNode{}}
    b.populateStack(root)
    return b
}

func (this *BSTIterator) populateStack(r *TreeNode) {
    for r != nil {
        this.stack = append(this.stack, r)
        r = r.Left
    }
}


func (this *BSTIterator) Next() int {
    top := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if top.Right != nil {
        this.populateStack(top.Right)
    }
    return top.Val
}


func (this *BSTIterator) HasNext() bool {
    return len(this.stack) > 0
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */