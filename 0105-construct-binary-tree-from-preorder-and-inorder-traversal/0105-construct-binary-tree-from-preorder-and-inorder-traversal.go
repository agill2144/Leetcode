/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    approach:
    
    Preorder -> root-left-right
    Inorder -> left-root-right
    
    * The trick here is to recursively collect all root nodes and pass them back to parent.Left and parent.Right
    
    1. Get the initial root value from preorder : it will always be the 0th idx
    2. Find the index of the root value in inorder
        - All elements to the left of rootIdx in inorder is the left subtree
        - All elements to the right to rootIdx in inorder is the right subtree
    3. Using the rootIdx found in inorder, we can locate the left and right subtree in preorder
        - Why do we need to do this?
        - Well we will always look for a starting root in recursive call for a parent.Left and parent.Right
        - Once a recursive call at the lowest level is done, the root obj gets passed back to its parent.Left and or parent.Right
        - So we are setting/finding root at each recursive call and caller parent will set the return value to root.Left and root.Right
    4. How can we find the preorder left and right subtree in preorder using rootIdx from inorder?
        - Well it happens to be that in preorder;
            - preOrder[1:rootIdx+1] - is the left subtree in preorder
                - Why start from 1 here? because 0 is already used - ITS THE INITIAL ROOT!
            - preOrder[rootIdx+1:] is the right subtree in preorder
    5. Now we have preLeft and preRight - i.e we have root for left subtree and right subtree
    6. Which means we can recursively call the buildTree with updated preorder and inorder
        - Inorder left would be everything to the left of rootIdx in inorder
        - Inorder right would be everything to the right of rootIdx in inorder
        - Which means its also important that we send down the updated inorder list to each recursive call so we can correctly find the rootIdx again relative to the new updated preOrder list
        
    
*/

// in this approach we will slice down the preorder and inorder in each recursive call to build tree from root using preorder style.
// slicing is expensive in terms of time and space
// func buildTree(preorder []int, inorder []int) *TreeNode{
    
//     if len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
//         return nil
//     }
    
//     // root will always be 0th idx in preorder
//     // and will be in each recursion call because in each recursion call, we slice down the preorder
//     // so its gauranteed.
    
//     // remeber we are building tree in preorder using preorder
//     // so process root first and then left n then right 
//     root := &TreeNode{Val: preorder[0]}
    
//     // now find the rootIdx in inorder
//     // all elements to the left of rootIdx in inorder list will be the left side of the tree ( from inorders perspective )
//     // all elements to the right of rootIdx in inorder list will be the right side of the tree ( from inorders perspective )
//     // remember we are only going down a direction to set what root should be here in left || right... 
//     rootIdx := -1
//     for i := 0; i < len(inorder); i++ {
//         if inorder[i] == root.Val {
//             rootIdx = i
//             break
//         }
//     }
    
    
//     // we can use the rootIdx from inorder to slice down the preorder
//     // preorder[1:rootIdx+1] is the preLeft and preorder[mid+1:] is preRight
//     // inorder[0:rootIdx] is the inLeft and inorder[rootIdx+1:] is inRight
//     root.Left = buildTree(preorder[1:rootIdx+1], inorder[0:rootIdx])
//     root.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])
//     return root
// }


// the above approach works however time complexity wise, we have o(n^2) because at each recursive call
// we search in o(n) time in inordet to find rootIdx
// and we also take (n) for each preOrder and inorder slicing 
// same thing for space, at each recursive call, we allocated a new smaller inorder and preorder list : o(n^2) 


// for searching over and over again, lets toss inorder into a val:idx map initially
// and for slicing inorder left and right , we will use 2 pointers as boundary idx ( start and end ) 

// time: o(n)
// space: o(n)

    
// instead of slicing we will use 3 ptrs
// 1. idx: to tell us where the next root in preorder is ( so we will gradually increase this ptr by 1 in each recursion call )
// 2. start: to tell us the start position in inorder
// 3. end: to tell us the end position in inorder
func buildTree(preorder []int, inorder []int) *TreeNode {
    ptr := 0
    inorderMap := map[int]int{}
    for idx, ele := range inorder {
        inorderMap[ele] = idx
    }
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        if ptr == len(preorder) {return nil}
        
        // logic
        root := &TreeNode{Val: preorder[ptr]}
        ptr++
        rootIdxInInorder := inorderMap[root.Val]
        root.Left = dfs(left, rootIdxInInorder-1)
        root.Right = dfs(rootIdxInInorder+1, right)
        return root
    }
    return dfs(0, len(preorder)-1)
}