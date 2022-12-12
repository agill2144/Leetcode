func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    m := len(image) // num of rows
    n := len(image[0]) // num of cols
    oldColor := image[sr][sc]
    if oldColor == color {return image}  
    dirs := [][]int{
        // { r, c }
        {0,1}, // go right
        {0,-1}, // go left
        {1, 0}, // go down
        {-1, 0}, // go up
    }
    
    
    q := [][]int{
        {sr,sc},
    }
    image[sr][sc] = color
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            
            // dq a node
            dq := q[0]
            q = q[1:]
            currRow := dq[0]
            currCol := dq[1]
            
            
            // process a node --- nothing to do
            
            
            // find childs n enqueue as needed
            for _, dir := range dirs {
                r := currRow + dir[0]
                c := currCol + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && image[r][c] == oldColor {
                    image[r][c] = color
                    q = append(q, []int{r,c})
                }
                
            }
            
            qSize--
        }
    }
    return image
}