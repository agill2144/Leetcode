/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    Sorted order == valid BST
    How to traverse a tree in a sorted order?
    - Inorder traversal of a BST guarentees a sorted order
    
    If you traverse a tree in inorder fashion, you will see nodes in asc order.
    
    approach 1: Populate a list while traversing inorder
    - inorder traverse
        - time: o(n) - where n is the number of nodes in the tree. We end up seeing all nodes.
        - space: o(h) - where h is the height of the tree. Regardless if we use recursion or iterative stack
    - push each item to a list
        - time: o(1)
        - space: o(n)
    - finally check if the list is sorted
        - time : o(n)
        - space : o(1)
    time: o(n) + o(1) + o(n) = o(n)
    space: o(n) + o(h) + o(1) = o(n)
    
    
    approach 2: Maintain a prev pointer to compare current root node with
    - inorder traversal
    - compare prev with current root
        - prev will always be smaller in a true BST ( prev in a number line thats sorted will always be smalled )
        - if prev >= current { return false }
    time: o(n)
    space: o(h)
    
    approach 3: Iterative implmentation of approach #2
*/


// inorder DFS ( recursive ) -- using prev pointer
// time: o(n) - we visit all nodes in best/worst case
// space: o(h) - max height of the tree will be in recursion stack at worse case and o(n) in a skewed tree

func isValidBST(root *TreeNode) bool {
    
    // why isnt prev part of recursion stack...
    // because when the inorder goes back to a root node from left
    // the prev to this root should be left.
    // if prev is maintained with root in recursion stack, it will never be the left child.
    // therefore global prev
//     var p *TreeNode
    
//     // inorder dfs
//     var inorderDfs func(c *TreeNode) bool
//     inorderDfs = func(c *TreeNode) bool {
        
//         // base
//         if c == nil {
//             return true
//         }
        
//         // logic
//         leftValid := inorderDfs(c.Left)
//         if !leftValid {
//             return false
//         }
//         // process
//         if p != nil {
//             if p.Val >= c.Val {
//                 return false
//             }
//         }
//         p = c
        
//         return inorderDfs(c.Right)
//     }
   
//     return inorderDfs(root)
    return preorderUsingMinMax(root, nil, nil)
}


// we need to check at each node
// it must either be less than parent ( left branch )
// or it must be greater than parent ( right branch )
// so at each level we have a min and max 
// min and max with the parent depending on which side we go on.
// left -> max parent, and min pass it as is
// right -> min parent and max pass it as is
// since we will be checking at each node whether its between min and max, and then only then go left and right, we will do this in preorder fashion

func preorderUsingMinMax(root, min, max *TreeNode) bool {
    // base 
    if root == nil {return true}
    if min != nil && root.Val <= min.Val {return false}
    if max != nil && root.Val >= max.Val {return false}
    
    // logic
    return preorderUsingMinMax(root.Left, min, root) && preorderUsingMinMax(root.Right, root, max)
    
}








