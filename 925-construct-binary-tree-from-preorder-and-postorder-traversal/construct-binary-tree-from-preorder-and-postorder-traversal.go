/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
//     if len(preorder) == 0 {return nil}
//     if len(preorder) != len(postorder) {return nil}
//     pos := map[int]int{}
//     for i := 0; i < len(preorder); i++ {
//         pos[postorder[i]] = i
//     }
//     ptr := 0
//     n := len(preorder)
//     var dfs func(left, right int) *TreeNode
//     dfs = func(left, right int) *TreeNode {
//         // base
//         if left > right {return nil}

//         // logic
//         root := &TreeNode{Val: preorder[ptr]}
//         ptr++
//         if left == right {return root}
//         postIdx := pos[preorder[ptr]]
//         root.Left = dfs(left, postIdx)
//         root.Right = dfs(postIdx+1, right-1)
//         return root
//     }
//     return dfs(0, n-1)
// }


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    if len(preorder) == 1 {
        return &TreeNode{Val: preorder[0]}
    }

    rootNode := &TreeNode{Val: preorder[0]}

    leftChildNode := preorder[1]

    leftEndNode := 0
    for ; leftEndNode < len(postorder); leftEndNode++ {
        if postorder[leftEndNode] == leftChildNode {
            break
        }
    }
    leftEndNode++
    leftSubTreeNodeCount := leftEndNode

    rootNode.Left = constructFromPrePost(preorder[1:1+leftSubTreeNodeCount], postorder[:leftSubTreeNodeCount])
    rootNode.Right = constructFromPrePost(preorder[1+leftSubTreeNodeCount:], postorder[leftSubTreeNodeCount:len(postorder) - 1])
    return rootNode
}