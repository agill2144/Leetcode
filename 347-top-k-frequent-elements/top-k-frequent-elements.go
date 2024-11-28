/*
    approach: custom sort based on freq
    tc = o(nlogn)
    sc = o(n)

    approach: min heap
    tc = o(n) + o(n*logk) + o(k)
    sc = o(n) + o(k)

    approach: bucket sort
    - whenever using heap / sorting, consider using bucket sort
    - we usually map nums[i] to ith idx in a bucket
    - therefore we would need to know the the largest possible value so that bucket[i] idx exists
        - what happens if we use bucket as freq?

    [2,3,1,2,1]
    - largest val = 3
    - therefore bucket of size 3+1 = 4 ( so that we have largest value avail as an idx )
                0 1 2 3
    - bucket = [0,2 2 1]
        0 repeats 0 times
        1 repeats 2 times
        2 repeats 2 times
        3 repeats 1 time
    - we use value at nums[i] as the idx position and incremented its freq
    - but this does not help find the top k frequent element
    - this traditional bucket sort, where we map value to idx, helps us perform classic sort
    - meaning, we know the val 1 repeats 2 times, and val 2 repeats 2 times
    - then we can created a sorted asc order array [1,1,2,2,3] ( using bucket freq )

    - instead we should use idxs in bucket as the freq!
        - vs idx as value mapping
    - we map freq count to idx ( freq = idx in bucket )
    - if we have a bucket of size n+1 ( at most the freq will of an element will be n )
        - and we want freq count available as an idx position
        - therefore we will create a bucket of size n+1
    
    [2,3,1,2,1]    
    - n = 5
    - therefore bucket of size 6 ( n+1 )
                0 1 2 3 4 5 6
    - bucket = [?,?,?,?,?,?,?]
    - now idxs of bucket act as frequency val ( instead of actual nums[i] value )
    - meaning, at idx 1 in bucket, we need to put in elements that repeat 1 time
    - at idx 6 in bucket, we need to put elements that repeat 6 times
    - therefore the bucket will be a list of list , where each element represents list of numbers that repeats i times
    - convert nums array to freq map
    {2:2, 3:1, 1:2} -> { number: freq }
    - map each freq to bucket idx
       0.  1.    2.   <---- idx acting as freq
    [ [], [3], [2,1]]
    - now we have elements sorted by freq!
    - and now we can start from the back of the array, collecting elements 
    - why from back?
        - because we want top k freq elements
        - meaning, k most frequent elements
        - the higher frequences in our bucket is obv from last idx
        - because idxs map to freq count
        - therefore from the back
    - k = 2
    - we want 2 most freq elements

       0.  1.    2.   <---- idx acting as freq
    [ [], [3], [2,1]]
                 i    
    out = [2,1]
    - now we have k elements in our output array, we can exit
    - if k = 3, we can keep moving back in our bucket array ( idx acting as freq )

       0.  1.    2.   <---- idx acting as freq
    [ [], [3], [2,1]]
           i    
    out = [2,1,3]
    - now we have k most freq elements, we can exit

    time
    - freq map creation = o(n)
        +
    - looping over bucket = o(n)
        *
        - unfurling nested list into flat, o(n)
        - but these are uniq elements in the nested list
        - therefore we are never going to see them again
        - therefore its not n^2 , its 2n = o(n)

    total time = o(n) + o(n) = o(n)
    space = o(n)

*/
func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }

	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
    		return nums[i] < nums[j]
		}
        return freq[nums[i]] < freq[nums[j]]
	})
    out := []int{}
    for i := len(nums)-1; i >= 0 && len(out) != k; i-- {
        for i+1 < len(nums) && i>=0 && nums[i] == nums[i+1] {i--}
        out = append(out, nums[i])
    }
    return out
}