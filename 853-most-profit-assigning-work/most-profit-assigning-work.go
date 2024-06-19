func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    n := len(difficulty)
    merged := [][]int{}
    for i := 0; i < n; i++ {
        merged = append(merged, []int{difficulty[i], profit[i]})
    }
    sort.Slice(merged, func(i, j int)bool{
        if merged[i][0] == merged[j][0] {
            return merged[i][1] < merged[j][1]
        }
        return merged[i][0] < merged[j][0]        
    })

    sort.Ints(worker)
    totalProfit := 0
    workerProfit := 0
    ptr := 0
    for i := 0; i < len(worker); i++ {
        workerSkillLevel := worker[i]
        for ptr < len(merged) {
            if merged[ptr][0] <= workerSkillLevel {
                workerProfit = max(workerProfit, merged[ptr][1])
                ptr++
            } else {
                break
            }
        }
        totalProfit += workerProfit
    }
    return totalProfit
}