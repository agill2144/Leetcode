/*
    approach: BFS
    - we have a start point and we have multiple destinations ( there could be n 0's)
    - enqueue start position and look for childs
        - Each item in this question only has 2 childs (you can jump to i + arr[i] or i - arr[i])
        - so ith val + i and ith val - 1
        - enqueue them both ( mark their inidices visited so we do not re-add them, also do boundary checks )
        - while processing an item we will check if this idx val == 0, return true
    - If our bfs exits without returning true, we will simply return false as in there was no way to get to an index with value 0
    
    time: o(n)
    space: o(n)        

*/
func canReach(arr []int, start int) bool {
    // visited := map[int]struct{}{start: struct{}{}}
    q := []int{start}
    n := len(arr)
    for len(q) != 0 {
        idx := q[0]
        q = q[1:]
        if arr[idx] == 0 {return true} 
        numJumps := arr[idx]
        forward := idx + numJumps
        backward := idx-numJumps
        
        if forward <= n-1 && forward >= 0 {
            if arr[forward] == 0 {  // only look for unvisited nodes ( i.e the ones with positive value )
                return true
            }
            if arr[forward] > 0 {
                q = append(q, forward)
                arr[forward] *= -1 // mark visited
            }
            // if _, ok := visited[forward]; !ok {
            //     q = append(q, forward)
            //     visited[forward] = struct{}{}
            // }
        }
        
        if backward >= 0 && backward <= n-1 {
            if arr[backward] == 0 {
                return true
            }
            if arr[backward] > 0 { // only look for unvisited nodes ( i.e the ones with positive value )
                q = append(q, backward)
                arr[backward] *= -1 // mark visited
            }
            // if _, ok := visited[backward]; !ok {
            //     q = append(q, backward)
            //     visited[backward] = struct{}{}
            // }
        }
        
    }
    
    return false
    
}
