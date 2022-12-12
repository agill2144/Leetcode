func findWinners(matches [][]int) [][]int {
    min := math.MaxInt64
    max := math.MinInt64
    resultMap := map[int][]int{} // { $id: [$loses, $countMatches] }

    for i := 0; i < len(matches); i++ {
        winner := matches[i][0]
        loser := matches[i][1]
        if val, _ := resultMap[winner]; val == nil {resultMap[winner] = []int{0,0}}
        if val, _ := resultMap[loser]; val == nil {resultMap[loser] = []int{0,0}}
        
        resultMap[winner][1]++
        resultMap[loser][0]++
        resultMap[loser][1]++

        min = int(math.Min(float64(min), math.Min(float64(winner), float64(loser))))
        max = int(math.Max(float64(max), math.Max(float64(winner), float64(loser))))
    }
    out := make([][]int, 2)
    for i := min; i <= max; i++ {
        data, ok := resultMap[i]
        if !ok {continue}
        if data[1] < 1 {continue}
        countLoses := data[0]
        if countLoses == 0 {
            out[0] = append(out[0], i)
        } else if countLoses == 1 {
            out[1] = append(out[1], i)            
        }
    }
    return out
}