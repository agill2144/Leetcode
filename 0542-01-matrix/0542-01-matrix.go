/*
    connected components!
    A dist in a cell (that already reached 0) can tell us a dist for a cell that cannot reach 0 directly
    
    approach: BFS
    - We need starting INDEPENDANT nodes.
    - 0s do not depend on anything, we are not looking for distance of each 0 to another 0
    - So 0s are independant
    - Enqueue all 0 positions
    - While we are all looking for all 0's , we will mark all the 1's = -1
        - Why? Well this is to distinguish between 1 as a distance vs visited/not-visited.
    - So when we are processing all 0's from our queue
    - We will look for all -1's ( unvisited nodes around this 0 )
    - When we find it, we will enqueue all neighbouring -1's ( but change the value of them to currentNode val + 1)
    - Then we will have cells with 1's in our queue
    - Keep looking for all -1's around each item from our queue, keep adding current val + 1 to new all -1's we find.
    - Instead of keeping the distance maintained at each level, the cell being processed itself will tell us how far away this cell is from a 0.
    - We can take that info and add + 1 for its neighbors that are -1.
    - And then repeat until full matrix is processed
    
    Time: o(mn)
    space: o(mn)
*/

// func updateMatrix(mat [][]int) [][]int {
    
    
//     if mat == nil {
//         return nil
//     }
//     dirs := [][]int{{1,0}, {-1,0}, {0,-1}, {0,1}}
//     m := len(mat)
//     n := len(mat[0])
//     q := [][]int{}
   
//      // enqueue all 0s and mark all 1s -1 so we can distinguish between visited and non-visited
//     // -1s are the non-visited ones, so when processing initial 0s, we will always search for non-visited nodes (i.e : -1)
//     for i := 0; i < m; i++ {
//         for j := 0; j < n; j++ {
//             if mat[i][j] == 0 {
//                 q = append(q, []int{i,j})
//             } else {
//                 mat[i][j] = -1
//             }
//         }
//     }
    
    
//     // each node we add while processing 0s, ( the -1s ), they are 1 distance away from 0
//     // since our intial queue will be all 0s, any -1s found around them will be 1 distance away ( or other words, currentNodeVal + 1)
//     // once we add the 1's to our queue, their neighbors will also be currentNodeVal + 1 distance away from nearest 0.
//     // if a parent (1) is 1 distance away from a 0, and if its childs which are 1 distance away from parent, then those childs are automatically parentDist+1 distance away from nearest 0. 
//     for len(q) != 0 {
//         dq := q[0]
//         q = q[1:]
        
//         for _, dir := range dirs {
//             r := dq[0] + dir[0]
//             c := dq[1] + dir[1]
//             if r >= 0 && r < m && c >= 0 && c < n && mat[r][c] == -1 {
//                 mat[r][c] = mat[dq[0]][dq[1]] + 1
//                 q = append(q, []int{r,c})
//             }
//         }
        
//     }
//     return mat
// }

/*
    approach: bottom up DP
    time: o(mn)
    space: o(1)
*/
func updateMatrix(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    topLeftDirs := [][]int{{0,-1},{-1,0}}
    bottomRightDirs := [][]int{{1,0},{0,1}}
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 0 {continue}
            minVal := math.MaxInt64-10
            for _, dir := range topLeftDirs {
                r := i+dir[0]
                c := j+dir[1]
                if r >= 0 && c >= 0 && mat[r][c] < minVal {
                    minVal = mat[r][c]
                }
            }
            mat[i][j] = minVal+1
        }
    }
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if mat[i][j] == 0 {continue}
            minVal := math.MaxInt64-10
            for _, dir := range bottomRightDirs {
                r := i+dir[0]
                c := j+dir[1]
                if r < m && c < n && mat[r][c] < minVal {
                    minVal = mat[r][c]
                }
            }
            mat[i][j] = min(mat[i][j], minVal+1)
        }
    }
    return mat 
}

func min(x, y int) int {
    if x < y {return x}
    return y
}