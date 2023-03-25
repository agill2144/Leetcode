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
    origToCp := map[*Node]*NodeCopy{}
    var dfs func(r *Node) *NodeCopy
    dfs = func(r *Node) *NodeCopy {
        // base
        if r == nil {return nil}
        
        // logic
        
        
        rootCopy, exists := origToCp[r]
        if !exists {
            rootCopy = &NodeCopy{Val: r.Val}
            origToCp[r] = rootCopy
        }
        if r.Random != nil {
            randomCp, ok := origToCp[r.Random]
            if !ok {
                randomCp = &NodeCopy{Val: r.Random.Val}
                origToCp[r.Random] = randomCp
            }
            rootCopy.Random = randomCp
            
        }            
        

        
        rootCopy.Left = dfs(r.Left)
        rootCopy.Right = dfs(r.Right)
        return rootCopy
    }
    cp := dfs(root)
    
    return cp
}