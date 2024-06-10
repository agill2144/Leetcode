/*
    approach: nsr + nsl ( process top ) pattern
*/
func sumSubarrayMins(arr []int) int {
    mod := 1000000007
    n := len(arr)
    st := []int{} // indices; processing top
    ans := 0
    for i := 0; i < n; i++ {
        for len(st) != 0 && arr[i] < arr[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i; nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            leftCount := top-nsl
            rightCount := nsr-top
            totalCount := leftCount * rightCount
            // arr[top] is min in $totalCount number of subarrays
            // and we have to sum of all minimums; i.e totalCount * arr[top]
            sum := arr[top] * totalCount
            ans = (ans+sum) % mod
        }
        st = append(st, i)
    }

    // monotonic stack edge cases;
    // do we need to do something with elements that are still in stack?
    // yes, because we were processing nsr, and if there are elements in stack
    // for example; [1,2,3,4]
    // it means we did not find a ith element that was a nsr of a previous element
    // therefore these elements range/window can be the whole array
    // therefore nsr for each of the remaining element is n (size of array)
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n; nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        leftCount := top-nsl
        rightCount := nsr-top
        totalCount := leftCount*rightCount
        sum := arr[top] *totalCount
        ans = (ans+sum) % mod
    }
    return ans
}

/*
    approach: brute force
    - used nested for loop
    - form all possible subarrays
    - keep track of curr min in each subarray
    - add curr min to total sum each time
    time = o(n^2)
    space = o(1)
*/
// func sumSubarrayMins(arr []int) int {
//     total := 0
//     n := len(arr)
//     mod := 1000000007
//     for i := 0; i < n; i++ {
//         currMin := arr[i]
//         for j := i; j < n; j++ {
//             currMin = min(currMin, arr[j])
//             total = (total + currMin) % mod
//         }
//     }
//     return total
// }