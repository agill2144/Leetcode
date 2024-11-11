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
    out.pushChilds(root)
    return out
}

func (this *BSTIterator) pushChilds(root *TreeNode) {
    for root != nil {
        this.st = append(this.st, root)
        root = root.Left
    }
}


func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    this.pushChilds(top.Right)
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