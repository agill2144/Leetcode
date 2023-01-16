/*
    approach #1: ( wrong )
    - save the entire bst inorder in a static list
    - take a snapshot of the entire tree in an inorder fashion and save the values in a list
    - then use ptr and the list to solve for next() and hasNext()
    
    BUT THE ABOVE IS INCORRECT!
    ITERATORS ARE SUPPOSED TO HANDLE DYNAMIC CHANGES AS MUCH AS POSSIBLE!
    
    What if the right subtree change values? 
    - couldnt we have a snapshot of the nodes in an inorder fashion?
    - we could, but
    
    What if the right subtree structly changes?
    - then we cannot use a list of treeNodes.
    
    So then, how do we handle this dynamic nature of changing BST, like a true iterator?
    - note that we wont be able to handle ALL dynamic changes!
    - we have to try and support as much dynamic changes as possible.
    
    We do not have access to the root of the BST outside of constructor,
    so somehow we have to have nodes in our data structure that lets us go a diff subtree
    
    We also must do this in an inorder fashion
    - which means left -> root -> right
    - but inorder dfs will not stop, its not controllable ( recurssion is not controllable )
    - however we can do inorder iteratively, then we can have control over the "fake over the hood recursion" using stack ( because recursion uses stacks )
    
    We need to split the inorder iterative function 
    inorder iterative psuedocode:
    stack = []
    for root != nil || stack is not empty {
        // left
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        // process root
        top = stack.pop()
        
        // go right
        root = top.Right 
    }
    
    approach #2: ( correct )
    - Whatever we want the answer to be for the next() call, IT MUST BE ON TOP OF THE STACK
    - In the constructor, we will save all the left nodes in a stack.
        - using inorder iterative way
    - next() - will pop the top of the stack, save the value and save subtree under top.Right back into the stack
    - hasNext() - uses len of stack to return the correct answer
    
*/
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    b := BSTIterator{st: []*TreeNode{}}
    b.inorderLeft(root)
    return b
}

func (this *BSTIterator) inorderLeft(root *TreeNode) {
    // dfs(r.Left) piece from recursion
    for root != nil {
        this.st = append(this.st, root)
        root = root.Left
    }
}


func (this *BSTIterator) Next() int {
    top := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    // dfs(r.Right) call which then starts a new recursion call and gets hit with dfs(r.Left) first
    this.inorderLeft(top.Right)
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
