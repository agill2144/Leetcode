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
    visited := map[int]map[int]bool{} // {$row: {$col: true|false}}
    dirs := [][]int{{-1,0},{0,1},{1,0},{0,-1}}
    currDir := 0
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if visited[r] != nil && visited[r][c] == true {return}

        // logic
        if visited[r] == nil {visited[r] = map[int]bool{}}
        visited[r][c] = true
        robot.Clean()
        for i := 0; i < 4; i++ {
            nr, nc := r+dirs[currDir][0], c+dirs[currDir][1]
            if (visited[nr] == nil || !visited[nr][nc]) && robot.Move() {
                dfs(nr, nc)
                robot.TurnRight()
                robot.TurnRight()
                robot.Move()
                robot.TurnRight()
                robot.TurnRight()
            }
            currDir++
            if currDir >= 4 {currDir = 0}
            robot.TurnRight()
        }
    }
    dfs(0,0)
}