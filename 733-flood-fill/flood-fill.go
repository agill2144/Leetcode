func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    m := len(image)
    n := len(image[0])
    oldColor := image[sr][sc]
    if oldColor == color {return image}
    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    var dfs func(r, c int)
    dfs = func(r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n {return}
        if image[r][c] != oldColor {return}

        // logic
        image[r][c] = color
        for _, dir := range dirs {
            dfs(r+dir[0], c+dir[1])
        }
    }
    dfs(sr, sc)
    return image
}