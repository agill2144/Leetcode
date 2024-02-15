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
    
    * The trick here is to recursively (top down) collect all root nodes and pass them back to parent.Left and parent.Right
    
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    idxMap := map[int]int{}
    for i := 0; i < len(inorder); i++ {
        idxMap[inorder[i]] = i
    }
    pre := 0
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        rootVal := preorder[pre]
        pre++
        root := &TreeNode{Val: rootVal}
        rootIdx := idxMap[rootVal]
        root.Left = dfs(left, rootIdx-1)
        root.Right = dfs(rootIdx+1, right)
        return root
    }
    return dfs(0, len(inorder)-1)
}