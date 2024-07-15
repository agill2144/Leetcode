/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    childs := map[int]struct{}{}
    valToNode := map[int]*TreeNode{}
    for i := 0; i < len(descriptions); i++ {
        parent := descriptions[i][0]
        child := descriptions[i][1]
        isLeft := descriptions[i][2] == 1
        childs[child] = struct{}{}
        parentNode, parentExists := valToNode[parent]
        if !parentExists {
            parentNode = &TreeNode{Val: parent}
            valToNode[parent] = parentNode
        }
        // at this point parent exists in valToNode
        childNode, childExists := valToNode[child]
        if !childExists {
            childNode = &TreeNode{Val: child}
            valToNode[child] = childNode
        }
        // at this point, both parent and child nodes exists
        if isLeft {
            parentNode.Left = childNode
        } else {
            parentNode.Right = childNode
        }
    }
    for i := 0; i < len(descriptions); i++ {
        a := descriptions[i][0]
        _, aIsChild := childs[a]
        if !aIsChild {return valToNode[a]}
    }
    return nil
}