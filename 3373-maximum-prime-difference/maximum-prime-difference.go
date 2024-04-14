func maximumPrimeDifference(nums []int) int {
    end := math.MinInt64
    for i := 0; i < len(nums); i++ {
        end = max(end, nums[i])
    }
    primes := make([]bool, end+1)
    for i := 0; i < len(primes); i++ {primes[i] = true}
    primes[0], primes[1] = false, false
    for i := 2; i*i < len(primes); i++ {
        if primes[i] {
            for j := i+i; j < len(primes); j+=i {
                primes[j] =  false
            }
        }
    }
    first := -1
    second := -1
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        if primes[num] {
            if first == -1 {first = i}
            second = i
        }
    }
    return second-first
}
