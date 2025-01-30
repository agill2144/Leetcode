func isRobotBounded(ins string) bool {
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    currDir := 0
    r, c := 0, 0
    for i := 0; i < len(ins); i++ {
        if ins[i] == 'G' {
            r += dirs[currDir][0]
            c += dirs[currDir][1]
        } else if ins[i] == 'R' {
            currDir++
            if currDir == 4 {
                currDir = 0
            }
        } else if ins[i] == 'L' {
            currDir--
            if currDir < 0 {
                currDir = 3
            }
        }
    }
    if currDir == 0 && r == 0 && c == 0 {return true}
    return currDir != 0 
}