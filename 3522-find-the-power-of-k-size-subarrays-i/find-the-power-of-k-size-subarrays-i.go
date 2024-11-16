/*
    approach: sliding window with dq maintaining next greater
    - monotonically increasing dq storing elements in asc order
    - top of dq will always be max 
    - how do we check if our current window is sorted and elements are consec?
    - thats what dq is literally telling us
    - if we pushed elements into dq , they would be increasing and and in consec order
    - therefore when we have a window of size k
    - and number of elements in our dq == k, then top of dq is our answer
    - when sliding win from left, if left idx == dq[0], drop dq[0]
    - what happens when current element is not sorted and or not in consec order?
    - we would traditionally not push such element into our mono increasing dq
    - however, it also no longer makes sense to keep any element in the dq anymore
    - because elements in dq must be consec and contagious
    - if we ran into a element that breaks the consec and sorted nature, we have to drop the entire dq
    - and form a new one!
    - so when current element is not consec / sorted when compared with top of dq
    - instead of just not pushing to dq, we will drop all elements from current dq
    - and start forming a new consec chain starting from curr element
    - therefore reset dq with just current element

    time = o(n)
    space = o(n)
*/
func resultsArray(nums []int, k int) []int {
    dq := []int{0} // idx
    n := len(nums)
    left := 0
    out := []int{}
    for i := 0; i < n; i++ {
        if len(dq) == 0 {
            dq = append(dq, i)
        } else if nums[i] == nums[dq[len(dq)-1]]+1 {
            dq = append(dq, i)
        } else {
            dq = []int{i}
        }

        if i-left+1 == k {
            if len(dq) == k {
                out = append(out, nums[dq[len(dq)-1]])
            } else {
                out = append(out, -1)
            }
            for len(dq) > 0 && dq[0] <= left {dq = dq[1:]}
            left++
        }
    }
    return out
}