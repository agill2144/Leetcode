/*
    approach: use graphs indegrees/outdegrees arrays 
    
    indegrees array to track number of incoming edge at a node,
    then whichever node has exactly n-1 value, thats the node with the most amount
    of incoming egdes 
    
    outdegrees array to track number of outgoing edges from a node,
    
    if we do not find a n-1 value, then that means there is no trusted person, return -1
    
    remember indegrees array counts the number of incoming edges to $this node ( i.e this translates to idx which is a node )
    
    in this case, we will track number of trusts given to a node 
    then look for a node with exactly n-1 value
    and only return that node if this node does not have any outgoing edges
    in other words, this person does not trusts 
    
    time: o(n)
    space: o(2n) or simply o(n) -- for indegrees array
*/

// func findJudge(n int, trust [][]int) int {
//     indegrees := make([]int, n+1) // count number of incoming edges to a node
//     outdegrees := make([]int, n+1) // count number of outgoing edges from a node
    
//     for _, ele := range trust {
//         person := ele[0]
//         trusts := ele[1]
        
//         indegrees[trusts]++
//         outdegrees[person]++
//     }
    
//     for i := 1; i < len(indegrees); i++ {
//         if indegrees[i] == n-1 && outdegrees[i] == 0 {
//             return i
//         }
//     }
//     return -1
// }

/*
    approach: use graphs indegrees ONLY array 
    
    indegrees array to track number of incoming edge at a node,
    then whichever node has exactly n-1 value, thats the node with the most amount
    of incoming egdes 
    
    But instead of maintaining a outdegrees array to keep track of count number of edges going
    from this node, if this node does have a outgoing edge(s), then we will simply deduct its indeegree value 
    by 1 for each outgoing edge

    Then we will look for n-1 val in indegree, return idx if we find such value, 
    if we do not find a n-1 value, then that means there is no trusted person, return -1
    
    remember indegrees array counts the number of incoming edges to $this node ( i.e this translates to idx which is a node )
    
    in this case,
    we will track number of trusts given to a node 
    and we will reduce indegrees of person value
    then look for a node with exactly n-1 value
    and only return that node if this node does not have any outgoing edges
    in other words, this person does not trusts 
    
    time: o(e) + o(v) OR o(v+e)
    space: o(v)
*/


func findJudge(n int, trust [][]int) int {
    indegrees := make([]int, n+1) // count number of incoming edges to a node
    
    for _, ele := range trust {
        a := ele[0]
        b := ele[1]
        
        indegrees[b]++
        indegrees[a]--
    }
    
    for i := 1; i < len(indegrees); i++ {
        if indegrees[i] == n-1 {
            return i
        }
    }
    return -1
}
