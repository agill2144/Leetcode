// time = o(n) + o(n) + o(nlogn) + o(n)
// space = o(n) + o(n)

func findLeastNumOfUniqueInts(arr []int, k int) int {
    freq := map[int]int{}
    for i := 0; i < len(arr); i++ {
        freq[arr[i]]++
    }
    freqArr := []int{}
    for _, v := range freq { freqArr = append(freqArr, v) }
    // sort 
    sort.Ints(freqArr)

    i := 0
    for i < len(arr) {
        count := freqArr[i]
        if count <= k {
            // we can remove all
            k -= count
            i++
            continue
        }
        break
    }
    return len(freqArr)-i
}

// func findLeastNumOfUniqueInts(arr []int, k int) int {
//     freq := map[int]int{}
//     for i := 0; i < len(arr); i++ {
//         freq[arr[i]]++
//     }
//     freqArr := [][]int{} // {val, freq}
//     for k, v := range freq {
//         freqArr = append(freqArr, []int{k,v})
//     }
//     // sort by freq
//     sort.Slice(freqArr, func(i, j int) bool{
//         iFreq, jFreq := freqArr[i][1], freqArr[j][1]
//         return iFreq < jFreq
//     })

//     i := 0
//     for i < len(arr) {
//         count := freqArr[i][1]
//         if count <= k {
//             // we can remove all
//             k -= count
//             i++
//             continue
//         }

//         break
//     }
//     return len(freqArr)-i
// }

/*
    [2,1,1,3,3,3]

    [[2,1] [1,2] [3,3]]
k=3.   2.    0.    i
*/