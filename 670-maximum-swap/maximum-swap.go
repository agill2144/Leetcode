func maximumSwap(num int) int {
    numStr := strings.Split(fmt.Sprintf("%v", num), "")
    idxs := make([]int, 10)
    for i := 0; i < len(idxs); i++ {idxs[i] = -1}
    for i := 0; i < len(numStr); i++ {
        ii, _ := strconv.Atoi(numStr[i])
        idxs[ii] = i
    }
    n := len(numStr)
    found := false
    for i := 0; i < n && !found; i++ {
        iVal, _ := strconv.Atoi(numStr[i])
        for j := 9; j > iVal; j-- {
            if idxs[j] != -1 && idxs[j] > i {
                maxIdx := idxs[j]
                numStr[i], numStr[maxIdx] = numStr[maxIdx], numStr[i]
                found = true
                break
            }
        }
    }
    out := 0
    for i := 0; i < len(numStr); i++ {
        val, _ := strconv.Atoi(numStr[i])
        out = out*10+val
    }
    return out

}
// func maximumSwap(num int) int {
//     numStr := strings.Split(fmt.Sprintf("%v", num),"")
//     n := len(numStr)
//     maxIdxs := make([]int, n)
//     maxIdxs[n-1] = n-1
//     for i := n-2; i >= 0; i-- {
//         currVal, _ := strconv.Atoi(numStr[i])
//         maxSoFar, _ := strconv.Atoi(numStr[maxIdxs[i+1]])
//         if currVal > maxSoFar {
//             maxIdxs[i] = i
//         } else {
//             maxIdxs[i] = maxIdxs[i+1]
//         }
//     }

//     for i := 0; i < n; i++ {
//         maxIdx := maxIdxs[i]
//         maxOnRight, _:= strconv.Atoi(numStr[maxIdx])
//         currVal, _ := strconv.Atoi(numStr[i])
//         if maxOnRight > currVal {
//             numStr[maxIdx], numStr[i] = numStr[i], numStr[maxIdx]
//             break
//         }
//     }
//     out := 0
//     for i := 0; i < len(numStr); i++ {
//         val, _ := strconv.Atoi(numStr[i])
//         out = out*10+val
//     }
//     return out
// }

/*
    126546
    555555

    626541
*/