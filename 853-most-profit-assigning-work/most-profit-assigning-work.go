func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    // greedy + sort + two ptrs

    // o(n)
    combined := [][]int{} // [ [$diffLevel, $profit ]]
    for i := 0; i < len(difficulty); i++ {
        diff := difficulty[i]
        pr := profit[i]
        combined = append(combined, []int{diff, pr})
    }

    // o(nlogn)
    // sort by difficulty
    sort.Slice(combined, func(i, j int)bool{
        iDiff, iPr := combined[i][0], combined[i][1]
        jDiff, jPr := combined[j][0], combined[j][1]
        if iDiff == jDiff {
            return iPr < jPr
        }
        return iDiff < jDiff
    })
    // sort the workers too
    sort.Ints(worker)
    totalProfit := 0
    combinedPtr := 0
    taskProfit := 0
    for i := 0; i < len(worker); i++ {
        wrDiff := worker[i]
        // because a job can be assigned to multiple workers
        // be greedy and keep a running max task profit 
        for combinedPtr < len(combined) {
            currDiff := combined[combinedPtr][0]
            currPr := combined[combinedPtr][1]
            // update mask task profit 
            // if its diff can be handled by worker
            if currDiff <= wrDiff {
                taskProfit = max(taskProfit, currPr)
                combinedPtr++
            } else {
                // if a worker cannot handle this diff
                // its fine, we have saved the max profit seen so far 
                break
            }
        }
        
        totalProfit += taskProfit
    }
    return totalProfit
}

// total time = o(n) + o(nlogn) + o(n^2)
// space = o(n)
// func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
//     // greedy + sort + two ptrs

//     // o(n)
//     combined := [][]int{} // [ [$diffLevel, $profit ]]
//     for i := 0; i < len(difficulty); i++ {
//         diff := difficulty[i]
//         pr := profit[i]
//         combined = append(combined, []int{diff, pr})
//     }

//     // o(nlogn)
//     // sort by difficulty
//     sort.Slice(combined, func(i, j int)bool{
//         iDiff, iPr := combined[i][0], combined[i][1]
//         jDiff, jPr := combined[j][0], combined[j][1]
//         if iDiff == jDiff {
//             return iPr < jPr
//         }
//         return iDiff < jDiff
//     })
//     totalProfit := 0
//     // o(n)
//     for i := 0; i < len(worker); i++ {
//         wr := worker[i]
//         left := 0
//         right := len(combined)-1
//         taskProfit := 0
//         // o(n)
//         for left <= right {
//             leftDiff, leftPr := combined[left][0], combined[left][1]
//             rightDiff, rightPr := combined[right][0], combined[right][1]
//             if leftDiff <= wr { taskProfit = max(taskProfit, leftPr)}
//             if rightDiff <= wr { taskProfit = max(taskProfit, rightPr)}
//             if leftPr < rightPr {
//                 left++
//             } else {
//                 right--
//             }
//         }
//         totalProfit += taskProfit
//     }
//     return totalProfit
// }

func max(x, y int)int {
    if x > y {return x}
    return y
}