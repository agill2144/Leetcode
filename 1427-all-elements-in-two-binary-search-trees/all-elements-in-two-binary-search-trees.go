/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
    out := []int{}
    st1 := []*TreeNode{}
    st2 := []*TreeNode{}
    for root1 != nil || root2 != nil || len(st1) != 0 || len(st2) != 0 {
        for root1 != nil {
            st1 = append(st1, root1)
            root1 = root1.Left
        }
        
        for root2 != nil {
            st2 = append(st2, root2)
            root2 = root2.Left
        }
        
        s1Peek, s2Peek := math.MaxInt64, math.MaxInt64
        if len(st1) != 0 {s1Peek = st1[len(st1)-1].Val}
        if len(st2) != 0 {s2Peek = st2[len(st2)-1].Val}
        
        // if s1Peek != math.MaxInt64 || s2Peek != math.MaxInt64 {
            if s1Peek < s2Peek {
                out = append(out, s1Peek)
                top := st1[len(st1)-1]
                st1 = st1[:len(st1)-1]
                root1 = top.Right
            } else {
                out = append(out, s2Peek)
                top := st2[len(st2)-1]
                st2 = st2[:len(st2)-1]
                root2 = top.Right
            }
        // }
    }
    return out
}
