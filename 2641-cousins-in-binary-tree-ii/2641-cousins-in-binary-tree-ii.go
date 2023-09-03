/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func replaceValueInTree(root *TreeNode) *TreeNode {
    levelSum := map[int]int{}
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        levelSum[level] += r.Val
        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    var siblingSumDfs func(r *TreeNode, level int)
    siblingSumDfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        if r == root {r.Val = 0}
        nextLevelSum ,ok := levelSum[level+1]
        if ok {
            // if we have next level's sum
            // $this node knows its childs
            // these 2 childs are siblings ( they belong to the same parent )
            // so standing at $this node, we can get sibling sum 
            // and remove sibling sum from nextLevelTotalSum
            // that will be sum of all cousins in the next level
            // then $this node can set its child values to the diff value
            siblingSum := 0
            if r.Left != nil {siblingSum += r.Left.Val}
            if r.Right != nil {siblingSum += r.Right.Val}
            if r.Left != nil {r.Left.Val = nextLevelSum - siblingSum}
            if r.Right != nil {r.Right.Val = nextLevelSum - siblingSum}
        }
        siblingSumDfs(r.Left, level+1)
        siblingSumDfs(r.Right, level+1)
    }
    dfs(root, 0)
    siblingSumDfs(root, 0)
    return root
}


// time = o(n)
// o(n) time for initial bfs +
// o(n) time for dfs

// space = o(n+maxWidth+h)
// o(maxWidth) queue space +
// o(n) space for levels map +   
// o(h) space for dfs recursive stack
// func replaceValueInTree(root *TreeNode) *TreeNode {
//     type level struct {
//         sum int
//         parentChildSum map[*TreeNode]int
//     }
//     levels := map[int]*level{}
//     q := [][]*TreeNode{{root, nil}} // {node, parent}
//     currLevel := 0
//     for len(q) != 0 {
//         qSize := len(q)
//         parentSum := map[*TreeNode]int{}
//         levelSum := 0
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             node := dq[0]
//             parent := dq[1]
            
//             levelSum += node.Val
//             parentSum[parent] += node.Val
            
//             if node.Left != nil {q = append(q, []*TreeNode{node.Left, node})}
//             if node.Right != nil {q = append(q, []*TreeNode{node.Right, node})}
//             qSize--
//         }
//         levels[currLevel] = &level{
//             sum: levelSum,
//             parentChildSum: parentSum,
//         }
//         currLevel++
//     }
//     var dfs func(r, p *TreeNode, level int)
//     dfs = func(r, p *TreeNode, level int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         if p == nil {
//             r.Val = 0
//         } else {
//             totalLevelSum := levels[level].sum
//             // siblings are childs with same value, their sum is saved by parent in each level
//             siblingSum := levels[level].parentChildSum[p]
//             r.Val = totalLevelSum - siblingSum
//         }
//         dfs(r.Left, r, level+1)
//         dfs(r.Right, r, level+1)
//     }
//     dfs(root, nil, 0)
//     return root
// }