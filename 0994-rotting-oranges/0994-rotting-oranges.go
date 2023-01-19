

/*
    "any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten"
    
    Connected components!
    Domino effect if 1 orange is rotten the  4-directionally adjacent to a rotten orange becomes rotten!
    Then the other 4 around the new rotten become rotten!
    We need to count the time its going to take to do this.
    
    Each time we apply the rotten to all 4 neighbors, we increase our time by 1
    
    Seems level by level - once 1 level is done, time++
    
    the fresh nodes depend on whether they have any rotten ones around them.
    I.e each fresh orange depends on a rotten neighbor
    
    However the rotten orange does not get worse than rotten
    its state does not change, which means the rotten oranges are independant!
    Which means our starting nodes will be all positions of rotten nodes

    
    approach 1:
    - level order using BFS
    - Search for all rotten oranges and enqueue them. While searching we will also count the number of fresh ones we have
        - Why? to answer the following: "until no cell has a fresh orange. If this is impossible, return -1."
    - When processing a level, check a rotten orange neighbours in all 4 directions
    - If fresh is found around this node, rotten it and enqueue it
    - Then rinse and repeat
    
    time: o(mn)
    space: o(mn) - if we start everything with rotten, we have enqueued all of the cells right away.
*/
func orangesRotting(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}

    fresh := 0
    q := [][]int{}    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                q = append(q, []int{i,j})
            } else if grid[i][j] == 1 {
                fresh++
            }
        }
    }
    if fresh == 0{return 0}
    if len(q) == 0{return -1}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            cr := dq[0]
            cc := dq[1]
            for _, dir := range dirs {
                r := cr + dir[0]
                c := cc + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
                    q = append(q,[]int{r,c})
                    fresh--
                    grid[r][c] = 2
                }
            }
            qSize--
        }
        level++
    }
    if fresh != 0 {return -1}
    return level-1
}