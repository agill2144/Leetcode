/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {return node}
    origToCp := map[*Node]*Node{node: &Node{Val:node.Val}}
    q := []*Node{node}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        for _, nei := range dq.Neighbors {
            _, ok := origToCp[nei]
            if !ok {
                origToCp[nei] = &Node{Val: nei.Val}
                q = append(q, nei)
            }
            origToCp[dq].Neighbors = append(origToCp[dq].Neighbors, origToCp[nei])
        }
    }
    return origToCp[node]
}