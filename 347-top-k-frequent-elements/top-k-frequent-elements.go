func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    deduped := []int{}
    for i := 0; i < len(nums); i++ {
        if freq[nums[i]] == 0 {
            deduped = append(deduped, nums[i])
        }
        freq[nums[i]]++
    }
    left := 0
    right := len(deduped)-1
    targetIdx := len(deduped)-k
    for left <= right {
        pivot := right
        nsf := left
        for i := left; i < pivot; i++ {
            if freq[deduped[i]] <= freq[deduped[pivot]] {
                deduped[nsf], deduped[i] = deduped[i], deduped[nsf]
                nsf++ 
            }
        }
        deduped[pivot], deduped[nsf] = deduped[nsf], deduped[pivot]
        if nsf == targetIdx {return deduped[nsf:]}
        if targetIdx < nsf {
            right = nsf-1
        } else {
            left = nsf+1
        }
    }
    return nil
}
