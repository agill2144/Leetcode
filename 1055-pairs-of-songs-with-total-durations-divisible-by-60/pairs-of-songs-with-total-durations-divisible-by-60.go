func numPairsDivisibleBy60(time []int) int {
    rem := map[int]int{}
    count := 0
    /*
        (a+b) % 60
        is same as
        a%60 + b%60

    */
    for i := 0; i < len(time); i++ {
        r := time[i] % 60
        
        count += rem[(60-r)%60]
        rem[r]++
    }
    return count
}