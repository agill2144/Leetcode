func countPrimes(n int) int {
    if n <= 1 {return 0}
    primes := make([]bool, n)
    for i := 0; i < len(primes); i++ {
        primes[i] = true
    }
    // 1 is not a prime number
    primes[1] = false
    count := 0
    for i := 1; i < len(primes); i++ {
        ithVal := i
        if primes[i] {
            count++
            // ith num is a prime number
            // therefore its multiples are not primes
            for j := ithVal*ithVal; j < len(primes); j+=ithVal {
                primes[j] = false
            }
        }
    }
    return count
}