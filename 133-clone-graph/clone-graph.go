/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

/*
    approach: 2 pass with hashmap
    - traverse the nodes and create copies
    - save copy into a hashmap { orig: copy }
    - then take a 2nd pass and create the edges

*/
func cloneGraph(node *Node) *Node {
    origToCp := map[*Node]*Node{}
    var dfs func(curr *Node)
    dfs = func(curr *Node) {
        // base
        if curr == nil {return}
        if _, ok := origToCp[curr]; ok {return}

        // logic
        cp := &Node{Val: curr.Val}
        origToCp[curr] = cp
        for _, nei := range curr.Neighbors {
            dfs(nei)
            cp.Neighbors = append(cp.Neighbors, origToCp[nei])
        }
    }
    dfs(node)
    return origToCp[node]
}