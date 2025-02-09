/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    var dfs func(r, c, n int) *Node
    dfs = func(r, c, n int) *Node {
        // base
        
        // logic
        if isAllSame(r,c,n,grid) {
            val := true // 1
            if grid[r][c] == 0 {val = false}
            return &Node{Val: val,IsLeaf:true}
        }
        root := &Node{Val:true}
        root.TopLeft = dfs(r, c, n/2)
        root.TopRight = dfs(r, c+n/2, n/2)
        root.BottomLeft = dfs(r+n/2, c, n/2)
        root.BottomRight = dfs(r+n/2, c+n/2, n/2)
        return root
    }
    return dfs(0, 0, len(grid))    
}

func isAllSame(r, c, n int, grid[][]int) bool {
    desired := grid[r][c]
    endRow := r+n
    endCol := c+n
    for i := r; i < endRow; i++ {
        for j := c; j < endCol; j++ {
            if grid[i][j] != desired {return false}
        }
    }
    return true
}