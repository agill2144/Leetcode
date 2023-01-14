/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type treeNodeWrapper struct {
    node *TreeNode
    alreadyProcessed bool
}
type BSTIterator struct {
    st []treeNodeWrapper
    prevSt []treeNodeWrapper
}


func Constructor(root *TreeNode) BSTIterator {
    b := BSTIterator{st: []treeNodeWrapper{}, prevSt: []treeNodeWrapper{}}
    for root != nil {
        b.st = append(b.st, treeNodeWrapper{node: root, alreadyProcessed: false})
        root = root.Left
    }
    return b
}


func (this *BSTIterator) HasNext() bool {
    return len(this.st) > 0
}


func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    if !top.alreadyProcessed {
        this.inorderLeft(top.node.Right)
        top.alreadyProcessed = true
    }
    this.prevSt = append(this.prevSt, top)
    return top.node.Val
}

func (this *BSTIterator) inorderLeft(root *TreeNode) {
    for root != nil {
        this.st = append(this.st, treeNodeWrapper{node: root})
        root = root.Left
    }
}


func (this *BSTIterator) HasPrev() bool {
    return len(this.prevSt) > 1
}


func (this *BSTIterator) Prev() int {
    top := this.prevSt[len(this.prevSt)-1]; 
    this.prevSt = this.prevSt[:len(this.prevSt)-1]
    this.st = append(this.st, top)
    return this.prevSt[len(this.prevSt)-1].node.Val
}



/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.HasNext();
 * param_2 := obj.Next();
 * param_3 := obj.HasPrev();
 * param_4 := obj.Prev();
 */