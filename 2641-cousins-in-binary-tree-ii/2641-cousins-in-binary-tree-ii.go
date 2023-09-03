/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    We have to replace a node value with sum of its cousins value
    Cousins are all nodes on the same level with different parents
    
    approach: two pass dfs
    - 1st pass: get level sum for each level and store it by level ( hashmap )
    - 2nd pass;
        - have each node look down and replace values of its childs
        - Each node has 2 childs, i.e this node is a parent of left and right child
        - so this node can look down and sum its childs sum ( siblings )
        - then this same node can get the next level sum ( currLevel + 1 ) from hashmap
        - and remove its childs sum from the levelSum
        - this becomes the new value for both of its childs
    
    time = o(n)
    space = o(n)
*/
func replaceValueInTree(root *TreeNode) *TreeNode {
    if root == nil {return root}
    levelSum := map[int]int{}
    var dfs1 func(r *TreeNode, level int)
    dfs1 = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        levelSum[level] += r.Val
        dfs1(r.Left, level+1)
        dfs1(r.Right, level+1)
    }
    // populate level sum
    dfs1(root, 0)
    
    var dfs2 func(r *TreeNode, level int)
    dfs2 = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        nextLevelSum, ok := levelSum[level+1]
        if ok {
            childSum := 0
            if r.Left != nil {childSum += r.Left.Val}
            if r.Right != nil {childSum += r.Right.Val}
            cousinSum := nextLevelSum - childSum
            if r.Left != nil {r.Left.Val = cousinSum}
            if r.Right != nil {r.Right.Val = cousinSum}
        }
        dfs2(r.Left, level+1)
        dfs2(r.Right, level+1)        
    }
    dfs2(root, 0)
    root.Val = 0
    return root
}