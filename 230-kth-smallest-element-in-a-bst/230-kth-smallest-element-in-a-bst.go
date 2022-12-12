/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    BST == running inorder gives us a sorted output  ( as long as its a valid bst )
    
    approach: brute force
    - run inorder
    - at each node, store the node val in an array
        - this array will be global and cannot be passed as arg because reference data structures combined with pass-by-value slice abstraction in golang will output incorrect answers
    - finally, return the kth smallest item arry from inorder ( idx: k-1 )
    
    time: o(n) -- we visit all nodes
    space: o(h) + o(n) --- h (height) in recursion stack, and o(n) for the array
*/

// func kthSmallest(root *TreeNode, k int) int {
//     result := []int{}
//     var inorder func(r *TreeNode)
//     inorder = func(r *TreeNode) {
//         // base
//         if r == nil {
//             return
//         }
        
//         // logic
//         inorder(r.Left)
//         result = append(result, r.Val)
//         inorder(r.Right)
//     }
//     inorder(root)
//     return result[k-1]
// }

/*
    BST == running inorder gives us a sorted output  ( as long as its a valid bst )
    
    approach: inorder dfs + count till k
    - run inorder recursively and maintain a count
    - everytime we come back to a node from left side, count++
    - if count == k - return that node val and exit early
    - otherwise recurse right
    
    time: o(h+k) -- worse case we visit max height ( all left nodes ) and then +k more nodes to get the kth smallest
    space: o(h)
*/
func kthSmallest(root *TreeNode, k int) int {
    var inorder func(r *TreeNode) int
    inorder = func(r *TreeNode) int {
        // base
        if r == nil {
            return -1
        }    
        
        // logic
        left := inorder(r.Left)
        if left != -1 {
            return left
        }
        
        k--
        if k == 0 {
            return r.Val 
        }
        
        
        right := inorder(r.Right)
        return right
    }
    
    return inorder(root)
    
}



/*
    BST == running inorder gives us a sorted output  ( as long as its a valid bst )
    
    approach: inorder (iterative) + decrement k instead of a aux var like count 
    - inorder iteratively
    - when we "fake pop" from stack, decrement k
    - if k == 0 , then we have found our kth smallest
        - return that val and exit early
    - otherwise go right 
    
    time: o(h+k) -- worse case we visit max height ( all left nodes ) and then +k more nodes to get the kth smallest
    space: o(h)
*/
// func kthSmallest(root *TreeNode, k int) int {
//     s := []*TreeNode{}
//     for root != nil || len(s) != 0 {
//         for root != nil {
//             s = append(s, root)
//             root = root.Left
//         }
//         root = s[len(s)-1]
//         s = s[:len(s)-1]
//         k--
//         if k == 0 {
//             return root.Val
//         }
//         root = root.Right
//     }
//     return root.Val
// }