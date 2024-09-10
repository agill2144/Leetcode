/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 /*
    why we cant use dfs while maintain nodes in a col ( i.e width ) ?
    - draw it out with 4 levels of tree and see the problem
    - the order matters, top nodes in a col should be placed first
    - with pre-order dfs, top nodes will not always be placed first
    - therefore inorder to respect top nodes to be placed first, and then bottom nodes
    - this is why we need bfs
    - level order, top nodes are processed first and then bottom nodes
    - to group together nodes per col, we can save them in a map { $colNum: [nodes..] }
     
 */
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	type qNode struct {
		node  *TreeNode
		width int
	}
	q := []*qNode{{root, 0}}
	minW := math.MaxInt64
	maxW := math.MinInt64
	widthToNodes := map[int][]int{}
	for len(q) != 0 {
		dq := q[0]
		q = q[1:]
		cn := dq.node
		cw := dq.width
		widthToNodes[cw] = append(widthToNodes[cw], cn.Val)
		minW = min(minW, cw)
		maxW = max(maxW, cw)
		if cn.Left != nil {
			q = append(q, &qNode{cn.Left, cw - 1})
		}
		if cn.Right != nil {
			q = append(q, &qNode{cn.Right, cw + 1})
		}
	}
	out := [][]int{}
	for i := minW; i <= maxW; i++ {
		out = append(out, widthToNodes[i])
	}
	return out
}