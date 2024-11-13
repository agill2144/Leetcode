func isRobotBounded(instructions string) bool {
    dirs := [][]int{{0,1},{1,0},{0,-1},{-1,0}}
    x, y := 0, 0
    n := 4
    dirPtr := 0
    i := 0
    for i < len(instructions) {
        if instructions[i] == 'G' {
            x += dirs[dirPtr][0]
            y += dirs[dirPtr][1]
        } else if instructions[i] == 'R' {
            dirPtr = (dirPtr+1) % n
        } else if instructions[i] == 'L' {
            dirPtr = (dirPtr+3) % n
        }
        i++
    }
    return (x == 0 && y == 0) || dirPtr != 0
}