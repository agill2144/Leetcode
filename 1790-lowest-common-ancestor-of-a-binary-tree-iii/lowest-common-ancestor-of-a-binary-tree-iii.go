/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
    if p == nil || q == nil {return nil}
    pPaths := []*Node{}
    qPaths := []*Node{}
    for p != nil {pPaths = append(pPaths, p); p = p.Parent}
    for q != nil {qPaths = append(qPaths, q); q = q.Parent}
    p1, p2 := len(pPaths)-1, len(qPaths)-1
    var out *Node
    for p1 >= 0 && p2 >= 0 {
        if pPaths[p1] != qPaths[p2] {
            break
        }
        out = pPaths[p1]
        p1--
        p2--
    }
    return out
}