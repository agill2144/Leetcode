/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

func copyRandomBinaryTree(root *Node) *NodeCopy {
    if root == nil {return nil}
    m := map[*Node]*NodeCopy{}
    var dfs func(r *Node) *NodeCopy
    dfs = func(r *Node) *NodeCopy{
        // base
        if r == nil {return nil}        
        // logic
        nodeCp, ok := m[r]
        if !ok {
            nodeCp = &NodeCopy{Val: r.Val}
            m[r] = nodeCp
        }
        if r.Random != nil {
            randCp := m[r.Random]
            if randCp == nil {
                randCp = &NodeCopy{Val: r.Random.Val}
                m[r.Random] = randCp
            }
            nodeCp.Random = randCp
        }
        nodeCp.Left = dfs(r.Left)
        nodeCp.Right = dfs(r.Right)
        return nodeCp
    }
    dfs(root)    
    return m[root]
}