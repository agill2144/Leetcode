// if asked to start from left -> right
// we keep track of building heights in a stack ( store idxs )
// when an ith building, if this building blocks prev building (top of stack )
// then remove prev building (top of stack) - and keep removing it until ith building keeps blocking ( while stack is !empty )
// finally, add this building, it might get removed later on
func findBuildings(heights []int) []int {
    st := []int{} // idx
    n := len(heights)
    for i := 0; i < n; i++ {
        for len(st) != 0 && heights[i] >= heights[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return st
}
// func findBuildings(heights []int) []int {
//     n := len(heights)
//     out := []int{n-1}
//     highest := heights[n-1]
//     for i := n-2; i >= 0; i-- {
//         if heights[i] > highest {
//             out = append(out, i)
//             highest = heights[i]
//         }
//     }
//     for i := 0; i < len(out)/2; i++ {
//         out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
//     }
//     return out
// }