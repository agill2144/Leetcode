/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func diameter(root *Node) int {
    ans := 0
    var dfs func(r *Node) int
    dfs = func(r *Node) int {
        // base
        if r == nil {return 0}

        // logic
        var first int
        var second int
        for _, child := range r.Children {
            depth := dfs(child)
            if depth > first {
                second = first
                first = depth
            } else if depth > second {
                second = depth
            }
        }
        ans = max(ans, first+second)
        return max(first,second)+1
    }
    dfs(root)
    return ans
}