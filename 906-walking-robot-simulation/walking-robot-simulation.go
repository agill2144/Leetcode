func robotSim(commands []int, obstacles [][]int) int {
    obSet := map[[2]int]bool{}
    for i := 0; i < len(obstacles); i++ {
        r,c := obstacles[i][0], obstacles[i][1]
        obSet[[2]int{r,c}] = true
    }
    // clockwise starting from east
    dirs := [][]int{{0,1},{1,0},{0,-1},{-1,0}} // east -> south -> west -> north
    ptr := 0 // facing east
    maxDist := 0
    r := 0
    c := 0
    for i := 0; i < len(commands); i++ {
        cmd := commands[i]
        if cmd == -1 || cmd == -2 {
            if cmd == -1 {ptr++} else { ptr-- }
            if ptr == -1 {ptr = len(dirs)-1} else if ptr == len(dirs) {ptr = 0}
        } else {
            currDir := dirs[ptr]
            nr := r
            nc := c
            stepsTaken := 0
            travelSteps := commands[i]
            for stepsTaken < travelSteps {
                nr += currDir[0]
                nc += currDir[1]
                stepsTaken++
                if obSet[[2]int{nr,nc}] { 
                    // ran into an obstacle
                    // cannot be in a cell that has obstacle
                    // therefore take 1 step back in the curr dir we are in
                    nr -= currDir[0]; 
                    nc -= currDir[1]; 
                    break
                }
                dist := int(math.Pow(float64(nr), 2.0)) + int(math.Pow(float64(nc), 2.0))
                maxDist = max(maxDist, dist)
            }
            r = nr
            c = nc
        }
    }
    return maxDist
}