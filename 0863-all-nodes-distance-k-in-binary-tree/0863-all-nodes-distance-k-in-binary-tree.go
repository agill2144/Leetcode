/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    
//     // turn the tree into a graph
//     // A TREE IS A UNDIRECTED GRAPH
//     // when you see tree, also consider undirected graph
    
//     // other appraoch is create parent ptrs so then we can traverse from target
//     // but we cannot modify input TreeNode struct, as it does not have "parent" field
//     // ok then use a hashmap to keep track of each parent per node :shrug ( implemented as another approach below )
//     // then ( whichever way we have parent ptrs )
//     // start from dfs(target, 0)
//     // each recursion will dfs(r.Left, currDist+1), dfs(r.Right, currDist+1), dfs(r.parent, currDist+1)
//     // save when currDist becomes k, and return back since going further when currDist is already k, will not help
    
//     // build the adjList
//     adjList := map[*TreeNode][]*TreeNode{}
//     var buildGraph func(r *TreeNode, parent *TreeNode)
//     buildGraph = func(r *TreeNode, parent *TreeNode) {
//         // base
//         if r == nil {return}
        
//         // logic
//         if parent != nil {adjList[r] = append(adjList[r], parent)}
//         if r.Left != nil {adjList[r] = append(adjList[r], r.Left)}
//         if r.Right != nil {adjList[r] = append(adjList[r], r.Right)}
//         buildGraph(r.Left, r)
//         buildGraph(r.Right, r)
//     }
//     buildGraph(root, nil)
    
//     // dfs graph traversal on a undirected graph ( undirected graphs when represented as a tree do not have cycles )
//     out := []int{}
//     // so we dont visit same node again
//     visited := map[*TreeNode]struct{}{} 
//     var dfs func(node, prev *TreeNode, currDist int)
//     dfs = func(node, prev *TreeNode, currDist int) {
//         // base
//         if node == nil {return}
//         if _, ok := visited[node]; ok {return}
        
//         // logic
//         visited[node] = struct{}{}
//         if currDist == k {
//             out = append(out, node.Val)
//             // going beyond this currDist, will exceed k
//             // therefore return and exit early
//             return
//         }
//         for _, nei := range adjList[node] {
//             if nei == prev {continue}
//             dfs(nei, node, currDist+1)
//         }
//     }
//     dfs(target, nil, 0)
//     return out
// }


func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    parents := map[*TreeNode]*TreeNode{}
    var dfs func(r *TreeNode, parent *TreeNode)
    dfs = func(r, parent *TreeNode) {
        // base
        if r == nil {return}
        
        // logic
        if parent != nil {
            parents[r] = parent
        }
        dfs(r.Left, r)
        dfs(r.Right, r)
    }
    dfs(root, nil)
    
    out := []int{}
    var distDfs func(node *TreeNode, currDist int)
    visited := map[*TreeNode]struct{}{}
    distDfs = func(node *TreeNode, currDist int)  {
        // base
        if node == nil {return}
        if _, ok := visited[node]; ok {return}
        
        // logic
        visited[node] = struct{}{}
        if currDist == k {
            out = append(out, node.Val)
            return
        }
        distDfs(node.Left, currDist+1)
        distDfs(node.Right, currDist+1)
        if _, ok := parents[node]; ok {
            distDfs(parents[node], currDist+1)            
        }
    }
    
    distDfs(target, 0)
    return out
}