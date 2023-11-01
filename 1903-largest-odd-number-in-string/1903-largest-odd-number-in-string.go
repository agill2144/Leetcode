func largestOddNumber(num string) string {
    n := 0
    largest := ""
    for i := 0; i < len(num); i++ {
        n  = n*10+int(num[i]-'0')
        if n % 2 != 0 {
            largest = num[0:i+1]
        }
    }
    return largest
}