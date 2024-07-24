func plusOne(digits []int) []int {
    out := []int{}
    carry := 0
    n := len(digits)
    for i := n-1; i >= 0; i-- {
        val := digits[i]
        if i == n-1 { val++ }
        total := val + carry
        lastDig := total % 10
        carry = total / 10
        out = append(out, lastDig)
    }
    if carry != 0 {
        out = append(out, carry)
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
    }
    return out
}