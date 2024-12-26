/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    BST == inorder gives us the sorted form ( as long as the tree is valid bst )
    In this, we do have an invalid BST and we have to fix it by swapping **values**
    
    Everytime we have binaryTree/BST question, evaluate whether you have 3 different types of cases
    1. where 2 could be part of the same subtree 
    2. where 2 nodes could be spread out in different subtrees
    3. where 2 nodes are part of same tree and in the same diagonal path (one is parent, other is child of the same parent)

    Here we are given that there will be EXACTLY 2 nodes who are in bad state;
    which means;
        1. The 2 nodes could be in the same subtree
        2. The 2 nodes could be spread out into 2 different subtrees
        3. The 2 nodes could be in the same subtree and in the same diagonal path

    approach: brute force
    - traverse inorder
    - write to a list
    - sort the list
    - use a ptr on list and traverse the bst in inorder
    - while processing a root node, write the value from sortedList[ptr] to current root node

    tc = o(n) + o(nlogn) + o(n)
    sc = o(n)
    
    
    approach: Without extra list/space
    - We can maintain 2 vars that tell us where BST becomes invalid
    - However, we have to be careful in a sense that the breach could happen in 2 places
        - 1. In the same subtree
            - For this case, we will only ever see 2 total breahces
            - We can easily the two breaches into 2 vars and swap them at the end
        - 2. In different subtrees
            - For this case, we will see total of 4 breaches.
            - When we run into our first group of breach ( prev, current )
            - Lets save them into first, second
            - Then IF we ever run into a breach again ( i.e when first is not nil , i.e we already saw a breach )
            - this means that the first breach first var needs to be swapped with this new current value
    
    - time: o(n)
    - space: o(h) because of recursion stack

*/
func recoverTree(root *TreeNode)  {
    if root == nil {return}
    var (
        fb *TreeNode
        sb *TreeNode
    )
    var prev *TreeNode
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        if prev != nil {
            if prev.Val >= r.Val {
                if fb == nil {fb = prev}
                sb = r
            }
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    fb.Val, sb.Val = sb.Val, fb.Val  
}