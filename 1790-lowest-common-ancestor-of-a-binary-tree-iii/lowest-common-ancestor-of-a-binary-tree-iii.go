/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

/*
    approach: hash hash, brain turning into a potato
    - create a map of nodes from p to root
        - for p != nil { add p to map, p = p.parent }
    - now we can traverse from q to root
        - and check if q has been in the map 
        - if yes, we found the lowest common node thats common between p and q
    time = o(p + q)
    space = o(p)

    approach: no space, using heights
    - 1st find root node ( take the traversal from p or q until we reach the root node )
    - then find heights from root to p and h ( pHeight and qHeight )
    - bring p or q at the same level ( if pHeight > qHeight , bring p to the same height as q )
        - vice-verse 
    - now that our p and q ptrs are at the same level / height
    - walk up both ptrs until they both meet
    - once both ptrs are on the same node, thats our LCA
    
    time = o(p) + o(n) + o(p+q)
    space = o(1) if we dont count recursion space 
*/

// no space solution
func lowestCommonAncestor(p *Node, q *Node) *Node {    
    root := p
    for root.Parent != nil {root = root.Parent}
    pH := -1
    qH := -1
    var getH func(r *Node, h int)
    getH = func(r *Node, h int) {
        // base
        if r == nil {return}
        if pH != -1 && qH != -1 {return }

        // logic
        if r == p {pH = h}
        if r == q {qH = h}
        getH(r.Left, h+1)
        getH(r.Right, h+1)
    }
    getH(root, 0)
    for pH > qH && p != nil {p = p.Parent; pH--}
    for qH > pH && q != nil {q = q.Parent; qH--}
    for p != q { p = p.Parent; q = q.Parent }
    return p
}