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
    
    Everytime we have binaryTree/BST question, evaluate whether you have 2 different types of cases
    where 2 could be part of the same subtree or 2 nodes could be spread out in different subtrees
    
    Here we are given that there will be EXACTLY 2 nodes who are in bad state;
    which means;
        1. The 2 nodes could be in the same subtree
        2. The 2 nodes could be spread out into 2 different subtrees
    
    approach: brute force
    - use inorder to generate a list of node values
    - sort them
    - use the sorted list to write back into the tree using inorder
    
    time:
    - o(n) to create initial list
    - + o(nlogn) to sort
    - + o(n) to write the sorted list back into the tree
    - o(2n) + o(nlogn) = o(nlogn)
    
    space:
    - o(h) + o(h) = 2 recursive funcions ( 1 to read and other to write )
    - o(n) = list to sort
    - total space = o(n) -- o(2h) will be smaller compared to o(n) ( not true in skewed tree tho )
*/
// func recoverTree(root *TreeNode)  {
//     nodes := []int{}
//     var listGenerator func(r *TreeNode)
//     listGenerator = func(r *TreeNode) {
//         if r == nil {return}
//         listGenerator(r.Left)
//         nodes = append(nodes, r.Val)
//         listGenerator(r.Right)
//     }
//     listGenerator(root)
    
//     sort.Ints(nodes)
//     ptr := 0
//     var inorderWriter func(r *TreeNode)
//     inorderWriter = func(r *TreeNode) {
//         if r == nil {
//             return
//         }
//         inorderWriter(r.Left)
//         r.Val = nodes[ptr]
//         ptr++
//         inorderWriter(r.Right)
//     }
//     inorderWriter(root)
// }

/*
    BST == inorder gives us the sorted form ( as long as the tree is valid bst )
    In this, we do have an invalid BST and we have to fix it by swapping **values**
    
    Everytime we have binaryTree/BST question, evaluate whether you have 2 different types of cases
    where 2 could be part of the same subtree or 2 nodes could be spread out in different subtrees
    
    Here we are given that there will be EXACTLY 2 nodes who are in bad state;
    which means;
        1. The 2 nodes could be in the same subtree
        2. The 2 nodes could be spread out into 2 different subtrees
    
    
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
    var first *TreeNode
    var second *TreeNode
    var prev *TreeNode
    var inorder func(r *TreeNode)
    inorder = func(r *TreeNode) {
        // base
        if r == nil {return}
        // logic
        inorder(r.Left)
        if prev != nil && prev.Val >= r.Val {
            if first == nil {
                first = prev
            }
            second = r
        }
        prev = r
        inorder(r.Right)
    }
    inorder(root)
    first.Val,second.Val = second.Val,first.Val
    
}