/*

*/
func totalFruit(fruits []int) int {
    left := 0
    maxWindow := 0
    idx := map[int]int{}
    for i := 0; i < len(fruits); i++ {
        idx[fruits[i]] = i
        if len(idx) > 2 {
            // invalid window
            // to make it valid, escape the earliest idx to keep window size large
            minIdx := len(fruits)+1
            minIdxFruit := -1
            for k, v := range idx {
                if v < minIdx {
                    minIdx = v
                    minIdxFruit = k
                }
            }
            // make sure left ptr has escaped the minIdx to make window valid
            if left <= minIdx {left = minIdx+1}
            // this fruit has left our window state
            // therefore can be removed
            delete(idx, minIdxFruit)
        }
        // now our window is valid
        maxWindow = max(maxWindow, i-left+1)
    }
    return maxWindow
}

/*
    we have 2 baskets
    and this means we can only have 2 types of fruits
    and the fruits selected must be consecutive

    so at max we can have 2 different types of fruits
    so at max we can have 2 distinct elements
    and we have to find longest subarray with atmost 2 distinct elements

    approach: classic sliding window
    - maintain a window ( freq )
    - if our window is valid, make it valid
        - move left ptr forward 1-by-1 while updating our window state
    - once we have a valid window, save its size as needed

    time = o(2n) = o(n)
    space = o(3) = o(1)
*/
// func totalFruit(fruits []int) int {
//     freq := map[int]int{}    
//     left := 0
//     maxWindow := 0
//     for i := 0; i < len(fruits); i++ {
//         freq[fruits[i]]++

//         // update window if its invalid
//         for len(freq) > 2 {
//             leftEle := fruits[left]
//             freq[leftEle]--
//             if freq[leftEle] == 0 {
//                 delete(freq, leftEle)
//             }
//             left++
//         }
     
//         // now we have a valid window
//         maxWindow = max(maxWindow, i-left+1)
//     }
//     return maxWindow
// }