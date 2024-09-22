/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    - node in a binary tree is supposed to have 2 nodes
    - total number of nodes in a full complete tree is; (2^h)-1
        - h = height of tree
    - we have to evaluate max width per level
    - therefore perform level order traversal using bfs
    - say we have current level and this level has the left most node and right most node
    - rest of the nodes in between them are missing
    - how do we find the width and count the missing number of nodes?
    - we need to mark each node with a number such that missing nodes in between would not matter

            x
           / \
          x   x
         /     \
        x       x
       /         \
      x           x
    - for example, the last level
    - if we could mark the left most and right most correctly with something that counts the missing node in between
    - then the width calc becomes simple
    - this is where we will use an idx to mark each node
    - idx is just a number that associated with each node
    - start idx with 1 for root node
    - left child idx will be 2*currNodeIdx
    - right child idx will eb 2*currNodeIdx+1
    - now the left most on last level will have an idx of 8
    - and the right most on last level will have an idx of 15
    - This accounts for missing nodes in between them
    - Now width calc is as simple as window size ; lastIdx-startIdx+1
    - we can use bfs q which store a pair ( node, currNodeIdx )
    - then when processing a dq'd node, we can enque left and right child with their corresponding idx
    time = o(n)
    space = o(n)
*/
func widthOfBinaryTree(root *TreeNode) int {
    type qNode struct {
        node *TreeNode
        idx int
    }
    q := []*qNode{&qNode{root,1}}
    maxW := -1
    for len(q) != 0 {
        qSize := len(q)
        start, end := q[0].idx, q[qSize-1].idx
        maxW = max(maxW, end-start+1)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.node.Left != nil {
                q = append(q, &qNode{dq.node.Left, 2*dq.idx})
            }
            if dq.node.Right != nil {
                q = append(q, &qNode{dq.node.Right, 2*dq.idx+1})
            }
            qSize--
        }
    }
    return maxW
}