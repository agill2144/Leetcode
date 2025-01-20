func isRobotBounded(ins string) bool {
    // clockwise dirs
    //                up  right  down   left   
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    currDir := 0
    r, c := 0,0
    for i := 0; i < len(ins); i++ {
        if ins[i] == 'G' {
            r += dirs[currDir][0]
            c += dirs[currDir][1]
        } else if ins[i] == 'R' {            
            // wherever we are facing,
            // turning right will always be +1
            // because our dirs are written in clockwise order
            currDir++
            if currDir == 4 {currDir = 0}
        } else if ins[i] == 'L' {
            // wherever we are facing,
            // turning left will awlays be -1
            // because our dirs are written in clockwise order
            currDir--
            if currDir < 0 {currDir = 3}
        }
    }
    if r == 0 && c == 0 {return true}
    return currDir != 0
}