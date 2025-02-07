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
    prev []*TreeNode
    inSt map[*TreeNode]bool
}


func Constructor(root *TreeNode) BSTIterator {
    out := BSTIterator{
        st: []*TreeNode{},
        prev: []*TreeNode{},
        inSt: map[*TreeNode]bool{},
    }
    out.addChilds(root)
    return out
}

func (this *BSTIterator) addChilds(r *TreeNode) {
    for r != nil {
        if !this.inSt[r] {
            this.inSt[r] = true
            this.st = append(this.st, r)
        }
        r = r.Left
    }
}

func (this *BSTIterator) HasNext() bool {
    return len(this.st) > 0
}


func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    this.inSt[top] = false
    this.prev = append(this.prev, top)
    this.addChilds(top.Right)
    return top.Val
}


func (this *BSTIterator) HasPrev() bool {
    return len(this.prev) > 1
}


func (this *BSTIterator) Prev() int {
    this.st = append(this.st, this.prev[len(this.prev)-1])
    this.inSt[this.prev[len(this.prev)-1]] = true
    this.prev = this.prev[:len(this.prev)-1]
    this.st = append(this.st, this.prev[len(this.prev)-1])
    this.inSt[this.prev[len(this.prev)-1]] = true
    this.prev = this.prev[:len(this.prev)-1]
    return this.Next()
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.HasNext();
 * param_2 := obj.Next();
 * param_3 := obj.HasPrev();
 * param_4 := obj.Prev();
 */