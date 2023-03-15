func hasValidPath(grid [][]int) bool {
    type streetDirs struct {
        r int
        c int
        validStreets map[int]bool
    }
    
    dirsPerSt := map[int][]streetDirs{
        1: []streetDirs{ {r:0, c:-1,validStreets: map[int]bool{1:true,4:true,6:true} }, {r:0,c:1, validStreets: map[int]bool{1:true,3:true,5:true}}},
        2: []streetDirs{ {r:-1, c:0,validStreets: map[int]bool{2:true,3:true,4:true} }, {r:1,c:0, validStreets: map[int]bool{2:true,5:true,6:true}}},
        3: []streetDirs{ {r:0, c:-1,validStreets: map[int]bool{1:true,4:true,6:true} }, {r:1,c:0, validStreets: map[int]bool{2:true,5:true,6:true}}},
        4: []streetDirs{ {r:0, c:1,validStreets: map[int]bool{1:true,3:true,5:true} }, {r:1,c:0, validStreets: map[int]bool{2:true,5:true,6:true}}},
        5: []streetDirs{ {r:0, c:-1,validStreets: map[int]bool{1:true,3:true,6:true} }, {r:-1,c:0, validStreets: map[int]bool{2:true,3:true,4:true}}},
        6: []streetDirs{ {r:0, c:1,validStreets: map[int]bool{1:true,3:true,5:true} }, {r:-1,c:0, validStreets: map[int]bool{2:true,3:true,4:true}}},

    }
    m := len(grid)
    n := len(grid[0])
    q := [][]int{{0,0}}
    grid[0][0] *= -1
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        cr := dq[0]
        cc := dq[1]
        ct := grid[cr][cc] * -1
        if cr == m-1 && cc == n-1 {return true}
        
        for _, stDir := range dirsPerSt[ct] {
            // fmt.Println(stDir)
            nr := cr + stDir.r
            nc := cc + stDir.c
            if nr < m && nc < n && nr >= 0 && nc >= 0 && grid[nr][nc] > 0 {
                st := grid[nr][nc]
                // fmt.Println("curr street: ", ct, " new street: ", st, " isValid: ", stDir.validStreets[st])
                if _, isValid := stDir.validStreets[st]; isValid {
                    if nr == m-1 && nc == n-1 {return true}
                    q = append(q, []int{nr,nc})
                    grid[nr][nc] *= -1
                }
            }
        }
    }
    return false
}