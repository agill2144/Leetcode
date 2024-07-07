/*
    approach: compute nsr, nsl while processing ith element
    - the main idea is, we evaluate whether ith element is the NSR for top element in stack
    - instead of processing ith element for nsl and or nsr
    - we will now process the top element
        - am i(ith element) your(top) nsr ?
        - yes, pop top and process top
            - pop top
            - popped elements nsl is now on the top of st
            - compute the width
            - compute the area
            - repeat if ith element is also the nsr for new top
        - push ith element to stack to be processed

    - IMPORTANT EDGE CASE WHEN PROCESSING TOP ELEMENTS USING NSR/NSL
    - what if ith element never becomes NSR for top element ?
    - that is, heights are in increasing order; [1,2,3,4]
    - then at each ith element, no one becomes the NSR of top
        - st = [1]; i = 2 ; 2 is not NSR of top (1)
        - st = [1,2]; i = 3; 3 is not NSR of top (2)
        - st = [1,2,3]; i = 4; 4 is not NSR of top (3)
        - st = [1,2,3,4]
    - this means, each ith element definitely does not have a NSR
    - therefore their right boundary is all the till the end
        - meaning ALL bars on right can be used as part of area calc
        - therefore their NSR is len(heights)
    - so now we have the NSR , we can process the st, like we were processing initially
    - pop the top, process it
        - top is the height
        - its nsr = n
        - its nsl = -1 (if st is empty, otherwise nsl is new top of st)
        - width = nsr-nsl-1
        - area = height*width
        - save the area if its bigger than what we have saved

    time = o(n)
    space = o(n)
*/
func largestRectangleArea(heights []int) int {
    ans := math.MinInt64
    n := len(heights)
    st := []int{}
    for i := 0; i < n; i++ {

        // we are now asking , am i(ith) your(top) nsr ? 
        for len(st) != 0 && heights[i] < heights[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i // curr
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            height := heights[top]
            width := nsr-nsl-1
            area := height * width
            ans = max(ans, area)
        }
        st = append(st, i)
    }

    // if we never got a nsr ( for example, elements were in increasing order )
    // [1,2,3,4]
    // then for each ith element, we never got a true, that yes, ith element is top's nsr
    // therefore we will have all of these elments in our st
    // This means, each st's element nsr is len of the full heights array
    for len(st) != 0 {
        // remeber; we are processing the top element!
        top := st[len(st)-1]
        st = st[:len(st)-1]
        height := heights[top]
        nsr := n
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        width := nsr-nsl-1
        area := height*width
        ans = max(ans, area)
    } 
    return ans
}


/*
    approach: precompute nsl and nsr
    - we have nested loops
    - and our nested loop start from i-1 and i+1
    - that is, its position depends on where i is ( stack hint )
    - also! our nested loops went until they reach a smaller height on left and right
        - this is NSL and NSR
        - if we pre-computed NSL and NSR, then the ith loop will become linear!
    - nsl and nsr is done via stack
    - but we want their idx positions ( so we can compute the widths ), not the nsr/nsl values
    time = o(n)
    space = o(n)
*/
// func largestRectangleArea(heights []int) int {
//     n := len(heights)
//     ans := math.MinInt64
//     nsr := nsr(heights)
//     nsl := nsl(heights)

//     for i := 0; i < n; i++ {
//         height := heights[i]
//         width := nsr[i]-nsl[i]-1
//         area := height * width
//         ans = max(ans, area)
//     }
//     return ans
// }

// naive nsl implementation
// process nsl for each ith element
// func nsl(nums []int) []int{
//     st := []int{}
//     n := len(nums)
//     out := make([]int, n)
//     for i := 0; i < len(nums); i++ {
//         out[i] = -1
//         // remove everyone that is >= to me (ith element); i is looking for its nsl in the stack
//         for len(st) != 0 && nums[st[len(st)-1]] >= nums[i] {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]
//         }
//         st = append(st, i)
//     }
//     return out
// }

// naive nsr implementation
// process nsr for each ith element
// func nsr(nums []int) []int{
//     st := []int{}
//     n := len(nums)
//     out := make([]int, n)
//     for i := n-1; i >= 0; i-- {
//         // why not -1 here?
//         // because we are using the output values to compute width
//         // if there is not a next smaller on right, it means, 
//         // ith bar goes as far right as possible , till the end, hence n
//         out[i] = n
//         // remove everyone that is >= to me (ith element); i is looking for its nsr in the stack
//         for len(st) != 0 && nums[st[len(st)-1]] >= nums[i] {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]
//         }
//         st = append(st, i)
//     }
//     return out
// }

/*
    approach: brute force
    - calculate area for each ith bar
    - find neighoring bars that can be part of this ith rectangle
    - now, we cant skip a bar and go to next one
        - therefore selected bars must be contagious
    - how do we decide whether neighboring bar can be part of this ith rectangle
    - if neighboring bar height < than ith bar, it cannot be selected
    - neighboring bar height must be >= ith bar height.
    - if we have N bars for an ith bar, which height do we use ?
        - smallest amongst all
        - smallest will be common amongst all
        - that is, every Nth bar will be of atleast $smallest height
        - so we know all N bars have the small height, therefore smallest among all
    - what about width ?
        - for each ith bar;
        - we go as far left as possible from i-1 till 0 as long as its height >= ith height
        - same for right side
        - then once our loop stops , we have 2 ptrs, left and right
        - to calculate width; right-left-1
    - area = minHeight * width

    time = o(n^2)
    space = o(1)
*/
// func largestRectangleArea(heights []int) int {
//     ans := math.MinInt64
//     n := len(heights)
//     for i := 0; i < n; i++ {
//         curr := heights[i]
//         minHeight := curr

//         // go as far left as possible
//         left := i-1
//         for left >= 0 && heights[left] >= curr {
//             minHeight = min(minHeight, heights[left])
//             left--
//         }
//         // go as far right as possible
//         right := i+1
//         for right < n && heights[right] >= curr {
//             minHeight = min(minHeight, heights[right])
//             right++
//         }

//         width := right-left-1
//         area := minHeight * width
//         ans = max(ans, area)        
//     }
//     return ans
// }
