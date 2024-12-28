func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    deduped := []int{}
    for key, _ := range freq {
        deduped = append(deduped, key)
    }
    
    left := 0
    right := len(deduped)-1
    for left <= right {
        ngf := left
        pivot := right
        for i := left ; i < pivot; i++ {
            if freq[deduped[i]] >= freq[deduped[pivot]] {
                deduped[ngf], deduped[i] = deduped[i], deduped[ngf]
                ngf++
            }
        }
        deduped[ngf], deduped[pivot] = deduped[pivot], deduped[ngf]
        if ngf == k-1 {return deduped[:k]}
        if k > ngf {
            left = ngf+1
        } else {
            right = ngf-1
        }
    }
    return nil
}