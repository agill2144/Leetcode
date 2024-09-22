/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    - number of nodes in a full complete binary tree
        - (2^h) - 1
        - h = height ( number of nodes including itself )
        - However, the left height and right height must be equal
        - therefore it must be a full complete binary tree
    - above ^ applies to every subtree too!
    - so if a subtree is not full / complete, recurse down until it is
    - once we have leftH == rightH, than the number of nodes in that subtree = (2^h)-1
    - then we can pass this answer back to the parent

    - but what happens when leftH != rightH
    - we recurse down and add current node to the recursion result
        - ie; 1 + dfs(r.Left) + dfs(r.Right)
        - the +1 is from this current node's uneven leftH and rightH
        - but we dont NOT want to count it, 
        - therefore we will recurse down to find even heights but also add this node to the main count
*/
func countNodes(root *TreeNode) int {
    getH := func(r *TreeNode) (int, int) {
        // optional check; 
        // dfs wont get call getH on a nil node
        // because dfs has a base condition to check for nil node
        if r == nil {return 0,0}

        l := r
        leftH := 1
        for l.Left != nil {
            l = l.Left
            leftH++
        }
        rightH := 1
        for r.Right != nil {
            r = r.Right
            rightH++
        }
        return leftH, rightH
    }
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        // base
        if r == nil {return 0}

        // logic
        lh, rh := getH(r)
        if lh == rh {
            return int(math.Pow(2.0, float64(lh)))-1
        } else {
            return 1 + dfs(r.Left) + dfs(r.Right)
        }
        return 0
    }
    
    return dfs(root)
}