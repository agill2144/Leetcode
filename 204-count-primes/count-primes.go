func countPrimes(n int) int {
    if n <= 1 {return 0}
    primes := make([]bool, n)
    for i := 0; i < len(primes); i++ {primes[i] = true}
    primes[0] = false
    primes[1] = false
    for i := 0; i < len(primes); i++ {
        if primes[i] {
            for j := i+i; j < len(primes); j+=i {
                primes[j] = false
            }
        }
    }
    count := 0
    for i := 2; i < len(primes); i++ {
        if primes[i] {count++}
    }
    return count
}