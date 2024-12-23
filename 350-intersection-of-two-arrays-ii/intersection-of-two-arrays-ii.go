/*
    collecting equal values from both arrays
    since the values matter, their positioning does not matter, we can always sort!
    sort opens up 2 options = two ptrs and binary search
    
    appraoch: binary search
    - we can search for the same number in the other array
    - CAREFUL, we cannot run binarySearch on entire array for each number
        - why not ? lets say
                             i
            smaller array = [1,1,2]
            larger array = [1,2,3,4]
        - first we search for 1 in larger array and we find it at idx 0

                               i
            smaller array = [1,1,7]
            larger array = [1,2,3,4]
        - then we search for another 1 in larger array and we again find it at idx 0
        - but 1 is only present in larger array once and smaller array has it twice
        - so we end up using it again and saying [1,1] are common elements in both
        - when the ans should have been just [1]
        
        fix #1
        a naive fix can be to keep track of indices returned from binary search in a hashset
        so that if we see an idx that we have already saved, we wont save it again
        
        fix #2
        we can shrink our binary search window each time we get a idx from it
        - that is, when we get a ans from binarySearch, 
            we will save this idx+1 in a var and use it next time as left position in binary search
        
        time = o(n1logn1) + o(n2logn2) + o(n2)*o(logn1)
*/
func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    // loop over smaller array and binary search on larger array
    // nums1 is always smaller
    out := []int{}
    left := 0
    for i := 0; i < len(nums1); i++ {
        idx := search(nums2, left, nums1[i])
        if idx != -1 {
            out = append(out, nums1[i])
            left = idx+1
        }   
    }
    return out
}

/*
    find the left most idx position of target
    but why left most ?
    - nums1 = [1,1]
    - nums2 = [1,1,2,2]; target 1; left = 0
    - its possible that we will the first idx position of $target at idx 1
    - and we will return idx 1 for nums1[0]
    - then our left ptr for binary search will be set to idx 2 ( left+1 )
    - now for nums1[0], again we need to look for the same element
    - but the starting position of binary search is at idx 2
    - therefore we wont find the 2nd one
    - this is why we need to find the left most possible idx of our target within binary search window!    
*/
func search(nums []int, left, target int) int {
    ans := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            ans = mid
            right = mid-1
            continue
        }
        if target > nums[mid] {left=mid+1} else {right=mid-1}
    }
    return ans
}