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
    var dfs func(r, c, currDir int)
    dfs = func(r, c, currDir int) {
        // base
        if visited[[2]int{r,c}] {return}

        // logic
        robot.Clean()
        visited[[2]int{r,c}] = true
        for i := 0; i < 4; i++ {
            if robot.Move() {
                nr := r + dirs[currDir][0]
                nc := c + dirs[currDir][1]
                dfs(nr,nc,currDir)
                robot.TurnRight()
                robot.TurnRight()
                robot.Move()
                robot.TurnRight()
                robot.TurnRight()
            }
            robot.TurnRight()
            currDir = (currDir+1)%4
        }
    }

    dfs(0,0,0)
}