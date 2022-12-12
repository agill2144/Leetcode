/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */
/*
    approach: BFS
    - Similar type of approach as Copy LL with random pointers ( approach that used extra space )
    time: o(v+e) where v is the number of verticies and e is number of edges.
    Space: o(v) - we only ever store verticies in our {$srcNode:$clonedCopy} map
*/
func cloneGraph(node *Node) *Node {
    if node == nil {return nil}
    visited := map[*Node]*Node{}
    
    // create clone of initial input node
    clone := &Node{Val: node.Val}
    visited[node] = clone // mark it visited
    q := []*Node{node} // enqueue it
    
    for len(q) != 0 { // process queue
        dq := q[0]
        q = q[1:]
        
        for _, nh := range dq.Neighbors { // go over the src node's childs
            _, ok := visited[nh] // check if they are already visited / in-the-queue 
            if !ok {  // if they are not, this means that the child's clone does not exist 
                visited[nh] = &Node{Val: nh.Val} // so create the child's clone and save it for ease of reference - also mark this child visited
                q = append(q, nh) // enqueue this child, becuase this child may have its own childs to be processed.
            }
            visited[dq].Neighbors = append(visited[dq].Neighbors, visited[nh]) // add this child's clone to current node's clone's list of neighbors
        }
    }
    return clone // finally return the initial clone of input node
}



// func cloneGraph(node *Node) *Node {
//     if node == nil {return nil}
//     visited := map[*Node]*Node{}
//     var dfs func(a *Node)
//     dfs = func(a *Node) {
//         // base
//         if a == nil {return }
//         if _, ok := visited[a]; ok {return} 
        
//         // logic
//         clone := &Node{Val:a.Val}
//         visited[a] = clone
//         for _, nh := range a.Neighbors {
//             dfs(nh)
//             clone.Neighbors = append(clone.Neighbors, visited[nh])
//         }
//     }
//     dfs(node)
//     return visited[node]
// }