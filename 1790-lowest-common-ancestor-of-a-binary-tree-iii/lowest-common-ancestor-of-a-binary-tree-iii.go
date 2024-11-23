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
    set := map[*Node]bool{}
    for p != nil {
        set[p] = true
        p = p.Parent
    }
    for !set[q] {
        q = q.Parent
    }
    return q
}