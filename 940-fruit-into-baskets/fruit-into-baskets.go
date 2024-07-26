func totalFruit(fruits []int) int {
    left := 0
    freq := map[int]int{}
    maxWindow := 0
    for i := 0; i < len(fruits); i++ {
        freq[fruits[i]]++
        if len(freq) > 2 {
            // invalid window

            // use non-shrinking window
            // reduce left fruit count from our window
            leftFruit := fruits[left]
            freq[leftFruit]--
            // if left fruit count is now 0, it means we have 
            // fully removed this left fruit from our window
            // therefore we can remove it from our window state.
            // thereby making the next iteration window valid
            if freq[leftFruit] == 0 {
                delete(freq, leftFruit)
            }
            left++
        } else {
            // valid window
            maxWindow = max(maxWindow, i-left+1)
        }
    }
    return maxWindow
}

/*
    nested sliding window optimization
    - can we make left ptr take bigger jumps instead of going 1-by-1 ?
    - yes
    - once we have a invalid window, it means we have more than 2 fruit types
    - therefore we need to escape/remove a fruit
    - we want to remove the earliest idx fruit so that we can keep a bigger window
    - therefore use a idx map to keep track of idx position of a fruit
        - if a fruit is already tracked, it gets overriden
        - and thats fine because if we have 3 fruits of type x
        - and we must escape x, then we have to all 3 x's
        - therefore the overriden value of x idx will be its last seen idx
    - move left to earliest idx + 1
    - remove this fruit from our window
    - now we have a valid window
    - save its size as needed

    time = o(n*3) 
    - the *3 because when window becomes invalid, we loop over idx map trying to find the earliest idx to remove from window
    - this map will never exceed 3
    - therefore o(n*3)
    - which is just o(n)
    space = o(3) = o(1)
*/
// func totalFruit(fruits []int) int {
//     left := 0
//     maxWindow := 0
//     idx := map[int]int{}
//     for i := 0; i < len(fruits); i++ {
//         idx[fruits[i]] = i
//         if len(idx) > 2 {
//             // invalid window
//             // to make it valid, escape the earliest idx to keep window size large
//             minIdx := len(fruits)+1
//             minIdxFruit := -1
//             for k, v := range idx {
//                 if v < minIdx {
//                     minIdx = v
//                     minIdxFruit = k
//                 }
//             }
//             // make sure left ptr has escaped the minIdx to make window valid
//             if left <= minIdx {left = minIdx+1}
//             // this fruit has left our window state
//             // therefore can be removed
//             delete(idx, minIdxFruit)
//         }
//         // now our window is valid
//         maxWindow = max(maxWindow, i-left+1)
//     }
//     return maxWindow
// }

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