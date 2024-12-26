/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    approach: store snapshot in a list
    - traverse the BST in constructor
    - and store values in a list
    - this is wrong because it does not handle changes to bst nodes as much as possible
    - iterators are supposed to handle dynamic changes :) 

    approach: controlled recursion
    - put all left nodes in a stack
    - such that top of the stack *ALWAYS* represents the next node in inorder traversal
    - when a node is popped from stack, go to its right and store all the left childs again
    - this way we have only stored left nodes
    - allowing right nodes to change overtime as needed
    - it does not handle FULL dynamic ability, but its tries to as much as possible
*/
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    out := BSTIterator{st: []*TreeNode{}}
    out.pushLeftChilds(root)
    return out    
}

func (this *BSTIterator) pushLeftChilds(r *TreeNode) {
    for r != nil {
        this.st = append(this.st, r)
        r = r.Left
    }
}

func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    this.pushLeftChilds(top.Right)
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