/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    TLDR;
    - Map binary tree nodes into a heap/array-like structure.
    - Start by placing the root node at index 1.
    - For any parent node at index `i`:
        - The left child is located at `2*i`.
        - The right child is located at `2*i + 1`.
    - Apply this rule recursively for each node in the tree.
    - Process nodes in a BFS manner, assigning indices as you go.
    - This works because the structure always assumes 2 children (left and right) for each node.
    - Even if a child is missing, the indexing still accounts for the position, allowing us to track width accurately.


    - Binary Tree Structure: 
        - Each node in a binary tree can have up to 2 children (left and right).
    - Complete Binary Tree Property: 
        - In a full, complete binary tree, the total number of nodes is (2^h) - 1, where h is the height of the tree.
    - Width Calculation: 
        - To compute the maximum width
        - perform a level-order traversal (BFS), focusing on the leftmost and rightmost nodes at each level.
    - Handling Missing Nodes: 
        - Even if some nodes are missing at a given level, 
        - we can account for them using a numbering/indexing system for nodes.
    Indexing System:
        - Assign an index (idx) to each node during traversal.
        - Root node starts at index 1.
        - Left child of a node at index i is at 2*i.
        - Right child of a node at index i is at 2*i + 1.
        - This ensures that missing nodes between leftmost and rightmost nodes are considered when calculating width.
    
    - Width = lastIdx - firstIdx + 1
        - where firstIdx and lastIdx represent the indices of the leftmost and rightmost nodes at that level
    
    - This indexing scheme leverages the mathematical properties of a complete binary tree, which can be represented like an array or heap.
    - Each node has predictable child positions relative to its index:
    - If a node is at index i, its left child is at 2*i (double the parent's index, as the binary tree "splits" into two).
    - The right child is at 2*i + 1, accounting for the sibling of the left child.
    - This allows for a consistent way to track node positions, even when nodes are missing, making it easier to calculate width and handle gaps efficiently.
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



