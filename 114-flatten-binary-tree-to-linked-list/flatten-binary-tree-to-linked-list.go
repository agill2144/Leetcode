/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    using morris traversal
    - take the right most on left side of curr ( call this tmp )
    - tmp.Right will always connect to curr.Right
        - therefore; tmp.Right = curr.Right
        - but not we lost curr.Right, we dont need it
        - because we will keep moving our curr ptr to right
    - once tmp.Right is connected to curr.Right
    - bring the entire curr.Left to curr.Right
        - i.e curr.Right = curr.Left
        - and curr.Left = nil
    - finally, move curr to curr.Right

    time = o(n)
    space = o(1)
*/
func flatten(root *TreeNode)  {
    if root == nil {return}
    curr := root
    for curr != nil {
        if curr.Left == nil {
            curr = curr.Right
        } else {
            tmp := curr.Left
            for tmp.Right != nil {tmp = tmp.Right}
            tmp.Right = curr.Right
            curr.Right = curr.Left
            curr.Left = nil
            curr = curr.Right
        }
    }
}