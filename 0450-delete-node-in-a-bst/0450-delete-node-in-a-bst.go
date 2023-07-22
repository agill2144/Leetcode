/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Successor = "after node" ( next node but smallest)
// Predecessor = "before node" (prev node but biggest)

func successor(r *TreeNode) *TreeNode {
    if r == nil {return nil}
    r = r.Right
    for r.Left != nil {r = r.Left}
    return r
}

func predecessor(r *TreeNode) *TreeNode {
    if r == nil {return nil}
    r = r.Left
    for r.Right != nil {r = r.Right}
    return r
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {return root}
        
    var dfs func(r, prev *TreeNode)
    dfs = func(r, prev *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if key > r.Val {
            dfs(r.Right, r)
        } else if key < r.Val {
            dfs(r.Left, r)
        } else if key == r.Val {
            
            newLeft := r.Left
            newRight := r.Right            

            if prev == nil {
                if newLeft != nil {
                    root = newLeft
                    tmp := newLeft
                    for tmp.Right != nil {
                        tmp = tmp.Right
                    }
                    tmp.Right = newRight
                } else if newRight != nil {
                    root = newRight
                    tmp := newRight
                    for tmp.Left != nil {
                        tmp = tmp.Left
                    }
                    tmp.Left = newLeft
                } else {root = nil}
                return
            }
            
            isLeftChild := prev.Left == r
            isRightChild := prev.Right == r
            
            // if leaf
            if r.Left == nil && r.Right == nil {
                if isLeftChild {prev.Left = nil}
                if isRightChild {prev.Right = nil}
                return
            }
            r.Left = nil
            r.Right = nil
            
            if newLeft != nil {
                // bring up left
                if isLeftChild {
                    prev.Left = newLeft
                    
                } else if isRightChild {
                    prev.Right = newLeft
                }
                tmp := newLeft
                for tmp.Right != nil {
                    tmp = tmp.Right
                }
                tmp.Right = newRight
            } else if newRight != nil {
                // bring up right
                if prev.Left == r {
                    prev.Left = newRight
                } else if prev.Right == r {
                    prev.Right = newRight
                }
                tmp := newRight
                for tmp.Left != nil {
                    tmp = tmp.Left
                }
                tmp.Left = newLeft
            }
            return
        }
    }
    dfs(root, nil)
    return root
}