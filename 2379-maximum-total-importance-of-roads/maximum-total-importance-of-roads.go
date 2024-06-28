// greedy; highest degree node should get the highest val assigned to maximize sum
func maximumImportance(n int, roads [][]int) int64 {
    degrees := make([]int, n)
    for i := 0; i < len(roads); i++ {
        a,b := roads[i][0], roads[i][1]
        degrees[a]++
        degrees[b]++
    }
    // if a node has x degrees, its value will b y
    // meaning if y will be added x times
    // i.e y*x
    // HOWEVER, we want to give highest value to node with highest number of edges
    // therefore sort 
    // ( we do not care if the mapping of node -> degrees is messed up )

    var res int64
    sort.Ints(degrees)
    for i := len(degrees)-1; i >= 0; i-- {
        res += int64(degrees[i]) * int64(n)
        n--
    }
    return res
}