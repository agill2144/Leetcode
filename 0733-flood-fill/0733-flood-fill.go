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

/*
    Connected components  == BFS/DFS
    - We need starting point(s) -- all independant nodes 
    There is a domino effect of changing 1 cell color to new color 
    """
        plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, 
        plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
    """
    
    approach : using BFS
    - We are given a starting point ( sr, sc )
    - Enqueue that starting
    - Save the color we are changing from ( oldColor, so we can compare with its neighbors )
    - Change the sr,sc to new color
    - Start processing the queue
    - As soon as we add an item to our queue, we will change the color right away to newColor
    - So that we do not visit this node again
    - if we are not allowed to modify the matrix ( save the visited poins in a set )
    time: o(mn)
    space: o(mn) 
    

*/

// func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
//     oldColor := image[sr][sc]
//     if oldColor == newColor {return image}
//     dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
//     m := len(image)
//     n := len(image[0])
    
//     q := new(queue)
//     image[sr][sc] = newColor
//     q.enqueue(sr,sc)

//     for !q.isEmpty() {
//         r,c := q.dequeue()
//         for _, dir := range dirs {
//             nr := r+dir[0]
//             nc := c+dir[1]
//             if nr >= 0 && nr < m && nc >=0 && nc < n && image[nr][nc] == oldColor {
//                 image[nr][nc] = newColor
//                 q.enqueue(nr,nc)
//             }
//         }
//     }
    
//     return image
// }