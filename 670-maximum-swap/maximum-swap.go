// bring the largest value in smallest unit to the largest unit
// time = o(n)
// space = o(n) for nums list + o(9) for idxs array = o(n)
func maximumSwap(num int) int {
    numStr := strings.Split(fmt.Sprintf("%v", num), "")
    idxs := make([]int, 10)
    for i := 0; i < len(idxs); i++ {idxs[i] = -1}
    // store the last occurrence of each digit
    for i := 0; i < len(numStr); i++ {
        ii, _ := strconv.Atoi(numStr[i])
        idxs[ii] = i
    }
    n := len(numStr)
    found := false

    // traverse the string to find the first digit that can be swapped with
    for i := 0; i < n && !found; i++ {
        iVal, _ := strconv.Atoi(numStr[i])
        // start with largest possible digit
        // and check if we have seen this digit before
        // because we can only use digits that exist within the num input
        for j := 9; j > iVal; j-- {
            /// if we found a digit larger than ith val
            // its idx must also be on the right of ith idx
            // we want to "bring the largest value in smallest unit to the largest unit"
            if idxs[j] != -1 && idxs[j] > i {
                // swap it, and break, because only 1 swap is allowed
                maxIdx := idxs[j]
                numStr[i], numStr[maxIdx] = numStr[maxIdx], numStr[i]
                found = true
                break
            }
        }
    }

    // reconstruct the num from list back into int
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