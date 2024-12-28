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
    targetIdx := len(deduped)-k
    for left <= right {
        nsf := left
        pivot := right
        for i := left ; i < pivot; i++ {
            if freq[deduped[i]] <= freq[deduped[pivot]] {
                deduped[nsf], deduped[i] = deduped[i], deduped[nsf]
                nsf++
            }
        }
        deduped[nsf], deduped[pivot] = deduped[pivot], deduped[nsf]
        if nsf == targetIdx {return deduped[nsf:]}
        if targetIdx > nsf {
            left = nsf+1
        } else {
            right = nsf-1
        }
    }
    return nil
}