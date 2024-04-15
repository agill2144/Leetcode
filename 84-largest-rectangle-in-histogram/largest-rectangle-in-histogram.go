func largestRectangleArea(heights []int) int {
    n := len(heights)
    if n == 0 {return 0}
    if n == 1 {return heights[0]}

    maxArea := math.MinInt64
    st := []int{} // index
    for i := 0; i <= n; i++ {
        for len(st) != 0 && (i == n || heights[i] < heights[st[len(st)-1]]) {
            // curr height is the nsr for the ele that is on top of stack
            // what about nsl ? its at the bottom of top of st
            height := heights[st[len(st)-1]]
            st = st[:len(st)-1]
            nsr := i
            width := nsr
            if len(st) != 0 {
                width = nsr-st[len(st)-1]-1
            }
            maxArea = max(maxArea, height*width)
        }
        st = append(st, i)
    }
    return maxArea
}
// func largestRectangleArea(heights []int) int {
//     nsl := nsl(heights)
//     nsr := nsr(heights)
//     ans := math.MinInt64
//     for i := 0; i < len(heights); i++ {
//         winSize := nsr[i] - nsl[i] + 1
//         area := heights[i] * winSize
//         ans = max(area, ans)
//     }
//     return ans
// }

// func nsr(arr []int) []int {
//     out := make([]int, len(arr))
//     st := []int{}
//     for i := len(arr)-1; i >= 0; i-- {
//         h := arr[i]
//         for len(st) != 0 && arr[st[len(st)-1]] >= h {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]-1
//         } else {
//             out[i] = len(arr)-1
//         }
//         st = append(st, i)
//     }
//     return out
// }

// func nsl(arr []int) []int {
//     out := make([]int, len(arr))
//     st := []int{}
//     for i := 0; i < len(arr); i++ {
//         h := arr[i]
//         for len(st) != 0 && arr[st[len(st)-1]] >= h {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]+1
//         }
//         st = append(st, i)
//     }
//     return out
// }