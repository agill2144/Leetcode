

// nested loops, with each nested loop depending on ith position - maybe stack
// nested dependent loops = maybe stack 
// what are some stack patterns; 1. ngr/ngl 2. nsr/nsl 3. looking back and removing
// in this case, for each ith element we kept expanding our left and right ptrs
// until we ran into a smaller height
// if we had nsr and nsl for each ith element, we do not have to check each time
// nsr/nsr can be done 2 ways; 
// 1. using extra space + stack ; while processing ith element
// 2. no extra space + stack; processing top
// what are we processing for top ?
// am i your(top) nsr ? if yes, nsl is the next top of stack
// now we have nsr and nsl, we can calculate width and area
// nsr/nsl processing of top edge cases;
// 1. we may not find a nsr while processing each ith element
// that is, all elements are in increasing order
// then we would have all elements in stack
// this also means, we do not have a nsr, each bar can be extended as far right as possible
// [1,2,3,4,5]
//. 1 -> 5, 2 -> 5, 3 -> 5, 4 -> 5
// this means nsr of each element is len(heights)
// therefore if there are still elements in stack after 1st pass
// take another pass processing remaining increasing ordered elements in stacl
// whose nsr is n=len(height) and nsl is the next top (if any)
// time = o(n)
// sapce = o(n)
func largestRectangleArea(heights []int) int {
    st := []int{} // indices
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        curr := heights[i]
        for len(st) != 0 && curr < heights[st[len(st)-1]] {
            // process top
            nsr := i
            top := st[len(st)-1]
            st = st[:len(st)-1]
            // if there are no elements inside the stack, it means
            // this bar can be expanded all the way to the left
            // therefore setting initial nsl idx = -1
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            width := nsr-nsl-1
            area := heights[top] * width
            maxArea = max(area, maxArea)
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        nsr := len(heights)
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        width := nsr-nsl-1
        area := heights[top] * width
        maxArea = max(area, maxArea)        
    }
    return maxArea
}


// brute force, try to expand each ith bar as far as possible to right and left
// this will give us the width, and then we can calc area of ith height
// repeat for all elements
// time = o(n^2)
// space = o(1)
// func largestRectangleArea(heights []int) int {
//     maxArea := 0
//     for i := 0; i < len(heights); i++ {
//         curr := heights[i]

//         // for each bar,
//         // we can extend as far right as possible
//         // but only if the next element on right is >= curr height
//         // we cannot expand a bar into next bar if next bar height is smaller than curr
//         right := i+1
//         for right < len(heights) && heights[right] >= curr {
//             right++
//         }

//         // same thing for left hand side
//         left := i-1
//         for left >= 0 && heights[left] >= curr {
//             left--
//         }

//         // now for this ith element, we have the width
//         width := right-left-1
//         area := curr * width
//         maxArea = max(maxArea, area)
//     }
//     return maxArea
// }