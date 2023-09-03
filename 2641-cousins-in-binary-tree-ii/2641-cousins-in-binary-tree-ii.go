/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
    type level struct {
        sum int
        parentChildSum map[*TreeNode]int
    }
    levels := map[int]*level{}
    q := [][]*TreeNode{{root, nil}} // {node, parent}
    currLevel := 0
    for len(q) != 0 {
        qSize := len(q)
        parentSum := map[*TreeNode]int{}
        levelSum := 0
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq[0]
            parent := dq[1]
            
            levelSum += node.Val
            parentSum[parent] += node.Val
            
            if node.Left != nil {q = append(q, []*TreeNode{node.Left, node})}
            if node.Right != nil {q = append(q, []*TreeNode{node.Right, node})}
            qSize--
        }
        levels[currLevel] = &level{
            sum: levelSum,
            parentChildSum: parentSum,
        }
        currLevel++
    }
    var dfs func(r, p *TreeNode, level int)
    dfs = func(r, p *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        if p == nil {
            r.Val = 0
        } else {
            totalLevelSum := levels[level].sum
            // siblings are childs with same value, their sum is saved by parent in each level
            siblingSum := levels[level].parentChildSum[p]
            r.Val = totalLevelSum - siblingSum
        }
        dfs(r.Left, r, level+1)
        dfs(r.Right, r, level+1)
    }
    dfs(root, nil, 0)
    return root
}