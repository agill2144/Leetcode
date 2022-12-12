/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach 1: level order using BFS
    - get the sum for each level
    - reset the global after processing a level
    - keep doing this until the end
    time: o(n) - where n is the number of nodes in the tree
    space: o(maxWidthOfTree) 
    
    
    approach 2: level order using DFS
    - maintain a level at each node
    - have a map {level: sum} being populated by the recursion
    - also maintain a maxLevel seen so far
    - finally return from map[maxLevel]
    time: o(n) - where n is the number of nodes
    space:
        1. recursion stack space: o(h)
        2. global map : o(h) - at max this map will hold h level key:val pairs -- since map is storing sum for each level
        o(h) + o(n)
*/


// approach 1: level order using BFS
// func deepestLeavesSum(root *TreeNode) int {
//     if root == nil {
//         return 0
//     }
//     sum := 0
//     queue := []*TreeNode{root}
//     for len(queue) != 0 {
//         qSize := len(queue)
//         lvlSum := 0
//         for qSize != 0 {
//             dq := queue[0]
//             queue = queue[1:]
//             lvlSum += dq.Val
//             if dq.Left != nil {queue = append(queue, dq.Left)}
//             if dq.Right != nil {queue = append(queue, dq.Right)}
//             qSize--
//         }
//         sum = lvlSum
//     }
    
//     return sum
// }


// approach 2: level order using DFS
type dfs struct {
    maxLevel int
    lvlSumMap map[int]int
}
func deepestLeavesSum(root *TreeNode) int {
    d := &dfs{lvlSumMap : map[int]int{}, maxLevel: 0}
    d.preorder(root, 0)
    return d.lvlSumMap[d.maxLevel]
}

func (d *dfs ) preorder(root *TreeNode, level int) {
    
    // base
    if root == nil {
        return
    }    
    
    // logic
    level = level+1
    if level > d.maxLevel {
        d.maxLevel = level
    }
    d.lvlSumMap[level] = d.lvlSumMap[level]+root.Val 
    d.preorder(root.Left, level)
    d.preorder(root.Right, level)
}