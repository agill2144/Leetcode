func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
    /*
        i = nums1 ; j = nums2 ; k = nums3
    
        i < j < k
        nums1 < nums3 < nums2

        nums2 will be the largest
        as soon as we find a number thats larger than nums2,
        promote nums2 to nums3 and update nums2 to new larger value

        st top will act as potential values for nums3
        if and when we find a element thats > top ( i.e ele > nums3 )
        pop the top, and assign top of st to nums3
        this larger element thats popping top of st is nums2
        nums2 idx should be smaller than nums3
        and THATS WHY WE NEED TO START TRAVERSAL FROM THE BACK OF THE ARRAY
        so that the elments in our stack are of greater indicies , so that they can be k
        once we find jth value ; i.e the largest, our stack at this point contains elements pushed from back of the array
        therefore j < k idx condition is already satisfied

        top of stack = potential k values ; potential nums3 values ; nums3 idx should be greater than nums2
        therefore the element that caused the popping of staack is our nums2 value ( greatest value ) and its idx
        will automatically be smaller than top of the stack ; because top of the stack holds nums3 values and nums3
        idx should be greater and it will be greater because we are pushing elements from the right side ( greater idx first )

        nums1 should be smaller than nums3
        so as soon as we run into an element thats smaller than nums3, we have our 132 pattern
        
        nums3 < nums2 is already handled via stack ( st top = nums3 and nums2 will be element that caused nums3 to get set )
        so it nums3 got set, it means we found a large enough value that popped elements from stack and setted the nums3 value
        so nums3 < nums2 is quietly satisfied 
    */
    nums2 := math.MinInt64
    nums3 := math.MinInt64
    st := []int{}
    for i := n-1; i >= 0; i-- {
        if nums[i] < nums3 && nums3 < nums2 {return true}
        for len(st) != 0 && nums[i] > st[len(st)-1] {
            nums2 = nums[i]
            nums3 = st[len(st)-1]
            st = st[:len(st)-1]
        }
        st = append(st, nums[i])
    }
    return false
}