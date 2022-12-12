func maxEnvelopes(envelopes [][]int) int {
    
    // sort by width
    sort.SliceStable(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] { // if widhts are same
            // sort in descending order of heights
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    
    // then just use same technique as longest increasing subsequence 
    // DP was o(n^2) -- TLE's on this 
    // then use binary search technique to build the effective sorted array
    
    // input: [[5,4],[6,4],[6,7],[2,3]]
    // after sorting by width: [[2,3],[5,4],[6,7],[6,4]]
    // now we will binary search on the sorted heights : idx 1
    
    effective := [][]int{envelopes[0]}
    
    for i := 1; i < len(envelopes); i++ {
        height := envelopes[i][1]
        lastIdxOfEffective := len(effective)-1
        if height > effective[lastIdxOfEffective][1] {
            effective = append(effective, envelopes[i])
        } else {
            nextSmallestHeightIdx := nextSmallestBinarySearch(height, effective)
            effective[nextSmallestHeightIdx] = envelopes[i]
        }
    }
    return len(effective)
}

// time : o(logn)
// space: o(1)
func nextSmallestBinarySearch(target int, nums [][]int) int {
    ans := -1
    left := 0
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid][1] >= target {
            if nums[mid][1] == target {return mid}
            ans = mid
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return ans
}