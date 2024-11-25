/*
    x^n can be re-written in two ways
    1. (x^n/2) * (x^n/2) 
    - this is basically (x^2)^n/2
    - like #2 ... 
    - where exponent computation is done FIRST
    - because PEMDAS
    - repeatedly squaring x and halving n

    2. (x^2)^n/2
    - repeatedly squaring x and halving n

    In both cases, we have to perform n/2
    which only works when exp is divisible by 2
    
    To handle odd exponents, we simply add another term of current base
    1. (x^n/2) * (x^n/2) * x
    - where exponent computation is done FIRST
    - because PEMDAS
    - the extra * x happens after the left handle is solved

    2. ((x^2)^n/2) * x
    - where the  brackets are solved first
    - because PEMDAS
    - the extra * x happens after the left handle is solved


    approach: brute force
    - run a loop 0 -> n times
    - and repeatedly multiply base
    time = o(n)
    space = o(1)

    approach: recursive binary exponentiation
    - we can use either #1 or #2 form of binary exponentiation
    - as long as we handle multiplying extra term when exp becomes odd
    - and making sure that the extra base term
        multiplication happens after exponent is computed,
        because PEMDAS
    time = o(logn)
    space = o(logn) ; recursive stack

    approach: iterative binary exponentiation
    time = o(logn)
    space = o(1)
*/
func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n *= -1
    }
    var res float64 = 1.0
    for n != 0 {
        if n % 2 != 0 {
            n--
            res *= x
        } else {
            x *= x
            n /= 2
        }
    }
    return res
}