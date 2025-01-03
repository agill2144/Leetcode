/**
 * // This is the robot's control interface.
 * // You should not implement it, or speculate about its implementation
 * type Robot struct {
 * }
 * 
 * // Returns true if the cell in front is open and robot moves into the cell.
 * // Returns false if the cell in front is blocked and robot stays in the current cell.
 * func (robot *Robot) Move() bool {}
 *
 * // Robot will stay in the same cell after calling TurnLeft/TurnRight.
 * // Each turn will be 90 degrees.
 * func (robot *Robot) TurnLeft() {}
 * func (robot *Robot) TurnRight() {}
 *
 * // Clean the current cell.
 * func (robot *Robot) Clean() {}
 */

func cleanRoom(robot *Robot) {
    //   0    1     2      3
    // [ up, right, down, left]
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    visited := map[[2]int]bool{}
    var dfs func(r, c, dir int)
    dfs = func(r, c, dir int) {
        // base

        if visited[[2]int{r,c}] {return}
        // logic
        visited[[2]int{r,c}] = true
        robot.Clean()
        for i := 0; i < 4; i++ {
            if robot.Move() {
                newDir := (dir+i) % 4
                nr, nc := r+dirs[newDir][0], c+dirs[newDir][1]
                dfs(nr,nc,newDir)
                // turn back around
                robot.TurnRight()
                robot.TurnRight()
                // go back to last cell
                robot.Move()
                // go back to the dir it was facing in
                robot.TurnRight()
                robot.TurnRight()
            }
            robot.TurnRight()
        }
    }

    dfs(0,0,0)
}