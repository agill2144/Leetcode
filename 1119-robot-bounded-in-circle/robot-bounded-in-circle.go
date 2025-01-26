func isRobotBounded(ins string) bool {
    //                 up, right, down, left
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    currDir := 0
    x, y := 0, 0
    n := len(ins)
    for i := 0; i < 2 * n; i++ {
        if ins[i%n] == 'G' {
            x += dirs[currDir][0]
            y += dirs[currDir][1]
        } else if ins[i%n] == 'R' {
            currDir++
            if currDir == len(dirs) {
                currDir = 0
            }
        } else if ins[i%n] == 'L' {
            currDir--
            if currDir < 0 {
                currDir = len(dirs)-1
            }
        }
    }
    if x == 0 && y == 0 {return true}
    return currDir != 0
}