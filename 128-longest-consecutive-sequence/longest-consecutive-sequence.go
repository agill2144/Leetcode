func longestConsecutive(nums []int) int {
    num_set := make(map[int]bool)
    for _, num := range nums {
        num_set[num] = true
    }
    longestStreak := 0
    for num := range num_set {
        if _, exists := num_set[num-1]; !exists {
            currentNum := num
            currentStreak := 1
            for exists := num_set[currentNum+1]; exists; exists = num_set[currentNum+1] {
                currentNum++
                currentStreak++
            }
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        }
    }
    return longestStreak
}