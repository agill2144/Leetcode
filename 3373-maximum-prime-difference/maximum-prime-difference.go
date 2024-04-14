/*
    a number is prime if;
    #1 MUST be > 1
    #2 MUST only have 2 factors ( i.e only EVER divisble by 1 and itself )

    - 1 is not a prime number 
        - condition 1 is not satisfied
    - 2 is a prime number ( > 1 and only divisible by 1 and itself )
    - if 2 is a prime number , can any other number divisble by 2 be a prime number?
        - 4,6,8,10
        - no, because 4 is divisible by 1,4 AND 2
        - therefore all numbers that are divisible by 2 are automatically NOT prime numbers
        - therefore we can run a nested loop that hops over 2 positions each time
            and mark other numbers divisible by 2 not prime number
    - 3 is a prime number (> 1 and is only divisible by 1 and itself )
    - if 3 is a prime number, can any other number divisble by 3 be a prime number ?
        - 6,9,12,15...
        - no, therefore start another loop that jumps 3 places to mark all numbers divisible by 3 not a prime number
    - until what number do we need to check ?
    - we checked 2,3, and marked other numbers that were divisible by 2 and 3 ; not prime
    - when do we stop ?
    - we dont, we exhaust all possible numbers
    - we will automatically skip some numbers in the middle
    - for example, next number is 4 ( which is already marked false, not a prime number )
        - therefore multiples of 4 are also marked false
    - then we will check 5,6,7, ... n

    This prime number pre-computation is called "Sieve of Eratosthenes"
    YT Ref: https://www.youtube.com/watch?v=V08g_lkKj6Q&t=179s

    approach: Sieve of Eratosthenes
    - precompute a list of prime numbers
    - our biggest possible number is given in constraints; 100
    - therefore generate a list that defines whether a num is prime or not
        - from 0 to 101
        - using "Sieve of Eratosthenes" approach
    - we need to maximize the difference between 2 indices
    - to maximize the difference between 2 indicies, we need
        - first idx that is a prime number
        - last idx that is a prime number
    - loop over nums array and check if each element is a prime number using 
        precomputes primes array
    - if num is prime and first idx is not yet set, set i as first idx
    - same for second idx
    - but we want last possible second idx, therefore keep override second idx if 
        an ith number is prime
    

    time = o()
    
*/
func maximumPrimeDifference(nums []int) int {
    // 1 <= nums[i] <= 100
    
    primes := make([]bool, 101)
    // assume all numbers are prime to start with
    for i := 2; i < len(primes); i++ {primes[i] = true}
    for i := 2; i*i < len(primes); i++ {
        if primes[i] {
            for j := i*i; j < len(primes); j+=i {
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
