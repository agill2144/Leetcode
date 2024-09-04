func robotSim(commands []int, obstacles [][]int) int {
    obSet := map[[2]int]bool{}
    for i := 0; i < len(obstacles); i++ {
        r,c := obstacles[i][0], obstacles[i][1]
        obSet[[2]int{r,c}] = true
    }
    dirMap := map[[2]int]map[int][2]int{
        [2]int{0,1}: map[int][2]int{-2:[2]int{-1,0}, -1:[2]int{1,0}},
        [2]int{0,-1}: map[int][2]int{-2:[2]int{1,0}, -1:[2]int{-1,0}},
        [2]int{-1,0}: map[int][2]int{-2:[2]int{0,-1}, -1:[2]int{0,1}},
        [2]int{1,0}: map[int][2]int{-2:[2]int{0,1}, -1:[2]int{0,-1}},
    }
    currDir := [2]int{0,1} // looking right 
    maxDist := 0
    r := 0
    c := 0
    for i := 0; i < len(commands); i++ {
        cmd := commands[i]
        if cmd == -1 || cmd == -2 {
            currDir = dirMap[currDir][cmd]
        } else {
            // fmt.Println("currDir: ", currDir)
            // fmt.Println("curr pos: ", r,c)
            // fmt.Println("steps: ", commands[i])
            nr := r
            nc := c
            stepsTaken := 0
            travelSteps := commands[i]
            for stepsTaken < travelSteps {
                nr += currDir[0]
                nc += currDir[1]
                stepsTaken++
                if obSet[[2]int{nr,nc}] { nr -= currDir[0]; nc -= currDir[1]; break }
                dist := int(math.Pow(float64(nr), 2.0)) + int(math.Pow(float64(nc), 2.0))
                maxDist = max(maxDist, dist)
            }
            r = nr
            c = nc
            // fmt.Println("post pos: ", r, c)
        }
    }
    return maxDist
}