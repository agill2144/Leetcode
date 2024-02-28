/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: validate bst min/max range pattern
    - we know our start root ; preorder[0]
    - left of root in a bst must be less than parent
        - i.e at max left child value MUST be < parent
        - i.e max = parent value
        - and min = no restriction
    - right of root in a base must be greater than parent
        - i.e at min right child value MUST be > parent
        - i.e min = parent value
        - and max = no restriction
    - use the min and max to identify whether to place this child or not
    - recursively build left and right subtrees

    time = o(n)
    space = o(h) max height ; worst case its o(n) for a one-sided tree
*/
func bstFromPreorder(preorder []int) *TreeNode {
    ptr := 0
    var dfs func(minVal, maxVal int) *TreeNode
    dfs = func(minVal, maxVal int) *TreeNode {
        // base
        if ptr == len(preorder) {return nil}
        if preorder[ptr] <= minVal || preorder[ptr] >= maxVal {return nil}
        
        // logic
        root := &TreeNode{Val: preorder[ptr]}
        ptr++
        root.Left = dfs(minVal, root.Val)
        root.Right = dfs(root.Val, maxVal)
        return root
    }
    return dfs(math.MinInt64, math.MaxInt64)
}
 /*
    approach: brute force
    - assuming the preorder is valid for a BST
    - we can copy that into a new array, sort it
    - this new sorted array is our inorder 
    - now we have 1. preorder and 2. inorder arrays
    - This question now becomes a traditional "construct tree using preorder and inorder"

    time = o(n) + o(nlogn) + o(n) = o(nlogn)
    space = o(n)
 */
// func bstFromPreorder(preorder []int) *TreeNode {
//     inorder := make([]int, len(preorder))
//     copy(inorder, preorder)
//     sort.Ints(inorder)
//     inorderIdx := map[int]int{}
//     for i := 0; i < len(inorder); i++ {
//         inorderIdx[inorder[i]] = i
//     }
    
//     prePtr := 0
//     var dfs func(left, right int) *TreeNode
//     dfs = func(left, right int) *TreeNode {
//         // base
//         if left > right {return nil}

//         // logic
//         rootVal := preorder[prePtr]
//         prePtr++
//         root := &TreeNode{Val: rootVal}
//         rootIdx := inorderIdx[rootVal]
//         root.Left = dfs(left, rootIdx-1)
//         root.Right = dfs(rootIdx+1, right)
//         return root
//     }
//     return dfs(0, len(inorder)-1)
    
// }