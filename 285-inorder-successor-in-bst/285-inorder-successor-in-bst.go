/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    The successor is basically the immediate next element in a sorted order.
    So the next greatest element after p
    
    approach: Brute force
    - remember inorder lets us run thru a bst in a order
    - we could just traverse the entire tree in inorder and create a sorted array
    - then using binary search, we could locate the next greater element
    time : o(n) + o(logn)
    space: o(h) + o(n)

    approach: optimized
    - since we are looking for the next greater element that comes after p
    - we are given the reference to p
    - which means we technically have access to p's right subtree
    - which means all we need to do is if p.right has a subtree
    - then we need to find the smallest elemet in the right subtree
        - finding min in a bst can be done by going left all the way
        - why are we looking for min in the right subtree of p?
        - because this is a bst, we need a successor
        - a sucessor is a element that is greater than p
        - elements that are greater than p is always going to be on the right side
        - because bst properties enforce that
        - so then why search for min in right subtree?
        - because the right subtree contains multiple elements
        - that are all greater than p itself
        - but we need the smallest one ( i.e the next immediate element )
        - which means we need the smallest element we can find in the right subtree ( which only contains elements > p )
        - therefore find min in the right subtree
    - what happens when p.right is nil
    - what happens when p does not have a right subtree ?
    - We could have 2 cases
        1. the immediate parent is the successor becasue this child is the left child its parent ( left(p) -> root(parent) -> right )
        2. p could be a right child of a parent which makes parent smaller than p
            - we could have a chain of right childs skewed to p
            n
             \
              n
               \
                n
                 \
                  p
            - all the above parent are less than p
            - therefore they are not successors.
            - so we need the deepest/lowest node that has a left child that is either directly p or a child that leads to p
            n
             \
              n  r(ancestor of p)
               \/
                n
                 \
                  n
                   \
                    p
                    
    time: o(n)
    space: o(1)
*/
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    if root == nil || p == nil {return nil}
    if p.Right != nil {
        curr := p.Right
        for curr.Left != nil {
            curr = curr.Left
        }
        return curr
    } 
    var ans *TreeNode
    for root != p {
        if p.Val > root.Val {
            root = root.Right
        } else {
            ans = root
            root = root.Left
        }
    }
    return ans
}