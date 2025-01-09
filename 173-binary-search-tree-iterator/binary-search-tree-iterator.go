/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    out := BSTIterator{st: []*TreeNode{}}
    out.pushNodes(root)
    return out
}

func (this *BSTIterator) pushNodes(r *TreeNode) {
    for r != nil {
        this.st = append(this.st, r)
        r = r.Left
    }
}


func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    this.pushNodes(top.Right)
    return top.Val
}


func (this *BSTIterator) HasNext() bool {
   return len(this.st) > 0 
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */