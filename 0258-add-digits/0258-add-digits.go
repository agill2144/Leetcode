func addDigits(num int) int {
    if num == 0 {return 0}
    total := 0
    for true {
        for num != 0 {
            last := num % 10
            total += last
            num /= 10
        }
        if total < 10 {break}
        num = total
        total = 0
    }
    return total
}