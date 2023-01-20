func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    m := len(image)
    n := len(image[0])
    oldColor := image[sr][sc]
    if oldColor == newColor {return image}
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || image[r][c] != oldColor || image[r][c] == newColor {
            return
        }
        
        // logic
        image[r][c] = newColor
        for _, dir := range dirs{
            dfs(r+dir[0],c+dir[1])
        }
        
    }
    dfs(sr, sc)
    return image
    
}
