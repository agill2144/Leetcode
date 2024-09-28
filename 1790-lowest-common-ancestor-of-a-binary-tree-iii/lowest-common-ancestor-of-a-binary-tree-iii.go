/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

// p = number of nodes from root to p
// q = number of nodes from root to q
// time = o(p) + o(q)
// space = o(p)
func lowestCommonAncestor(p *Node, q *Node) *Node {
    pPathNodes := map[*Node]bool{}
    for p != nil {
        pPathNodes[p] = true
        p = p.Parent
    }
    for q != nil {
        if pPathNodes[q] {return q}
        q = q.Parent
    }
    return nil
}