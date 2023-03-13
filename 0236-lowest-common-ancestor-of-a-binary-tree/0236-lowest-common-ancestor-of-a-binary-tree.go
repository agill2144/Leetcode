/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
        
     var dfs func(r, target *TreeNode, path []*TreeNode) []*TreeNode
     dfs = func(r, target *TreeNode, path []*TreeNode) []*TreeNode {
         // base
         if r == nil {return nil}
         
         // logic
         path = append(path, r)

         if r == target {
             newL := make([]*TreeNode, len(path))
             copy(newL, path)
             return newL
         }
         if ok := dfs(r.Left, target, path); ok != nil {return ok}
         if ok := dfs(r.Right, target, path); ok != nil {return ok}
         path = path[:len(path)-1]
         return nil
     }
     
     pathToP := dfs(root, p, nil)
     pathToQ := dfs(root, q, nil)
     var lastCommon *TreeNode
     pPtr := 0
     qPtr := 0
     for pPtr < len(pathToP) && qPtr < len(pathToQ) {
         if pathToP[pPtr] == pathToQ[qPtr] {
             lastCommon = pathToP[pPtr]
             pPtr++
             qPtr++
         } else {
             break
         }
     }
     return lastCommon
}