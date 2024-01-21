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
    pMap := map[*Node]struct{}{}
    for p != nil {
        pMap[p] = struct{}{}
        p = p.Parent
    }
    for q != nil {
        if _, ok := pMap[q]; ok {return q}
        q = q.Parent
    }
    return nil
}