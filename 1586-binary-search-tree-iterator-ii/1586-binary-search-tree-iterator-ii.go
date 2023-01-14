/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st *stackNode
    prev *stackNode
    visited map[*TreeNode]struct{}
}


func Constructor(root *TreeNode) BSTIterator {
    st := &stackNode{treeNode: &TreeNode{Val: math.MinInt64}}
    prev := &stackNode{treeNode: &TreeNode{Val: math.MinInt64}}
    b := BSTIterator{visited: map[*TreeNode]struct{}{}, st: st, prev:prev}
    b.inorderLeft(root)
    return b
}


func (this *BSTIterator) HasNext() bool {
    dummyHead := this.st
    actualHead := dummyHead.Next
    return actualHead != nil
}


func (this *BSTIterator) Next() int {
    top, newHead := pop(this.st)
    this.st = newHead
    treeNode := top.treeNode
    _, ok := this.visited[treeNode]
    if !ok {
        this.inorderLeft(treeNode.Right)
        this.visited[treeNode] = struct{}{}
    }
    this.prev = add(this.prev, treeNode)
    return treeNode.Val
}

func (this *BSTIterator) inorderLeft(root *TreeNode) {
    for root != nil {
        newHead := add(this.st, root)
        this.st = newHead
        root = root.Left
    }
}
func (this *BSTIterator) HasPrev() bool {
    dummyHead := this.prev
    actualHead := dummyHead.Next
    return actualHead != nil && actualHead.Next != nil
}


func (this *BSTIterator) Prev() int {
    top, newPrevHead := pop(this.prev)
    this.prev = newPrevHead
    this.st = add(this.st, top.treeNode)
    return this.prev.treeNode.Val
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.HasNext();
 * param_2 := obj.Next();
 * param_3 := obj.HasPrev();
 * param_4 := obj.Prev();
 */

type stackNode struct {
    treeNode *TreeNode
    Next *stackNode
}

func add(head *stackNode, x *TreeNode) *stackNode {
    if head == nil {
        return &stackNode{
            treeNode: x,
        }
    }
    newHead := &stackNode{treeNode: x}
    newHead.Next = head
    head = newHead
    return head
} 

func pop(head *stackNode) (*stackNode, *stackNode) {
    if head == nil || head.Next == nil {
        return head, nil
    }
    out := head
    newHead := head.Next
    head.Next = nil
    head = newHead
    return out, head
} 