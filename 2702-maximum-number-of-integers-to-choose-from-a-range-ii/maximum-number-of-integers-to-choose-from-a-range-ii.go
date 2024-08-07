import (
    "sort"
)

func maxCount(banned []int, n int, maxSum int64) int {
    sort.Ints(banned)
    
    count := 0
    currentSum := int64(0)
    bannedIndex := 0
    
    for i := 1; i <= n; i++ {
        // Skip if i is in the banned list
        if bannedIndex < len(banned) && banned[bannedIndex] == i {
            bannedIndex++
            continue
        }
        
        // Check if adding i would exceed maxSum
        if currentSum+int64(i) > maxSum {
            break
        }
        
        currentSum += int64(i)
        count++
    }
    
    return count
}
