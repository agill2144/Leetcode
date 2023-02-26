func addToArrayForm(num []int, k int) []int {
    ptr := len(num)-1
    ans := []int{}
    carry := 0
    for ptr >= 0 {
        lastK := 0
        if k != 0 {
            lastK = k % 10
            k /= 10
        }
        sum := num[ptr] + lastK + carry
        carry = 0
        if sum >= 10 {
            ans = append(ans, sum%10)
            carry = sum / 10
        } else {
            ans = append(ans, sum)
        }
        ptr--
    }
    for carry != 0 {
        lastK := 0
        if k != 0 {
            lastK = k % 10
            k /= 10
        }
        sum := carry + lastK
        carry = 0
        if sum >= 10 {
            ans = append(ans, sum%10)
            carry = sum / 10
        } else {
            ans = append(ans, sum)
        }
    }
    
    for k != 0 {
        ans = append(ans,k%10)
        k /= 10
    }
    reverse(ans)
    return ans
    
}

func reverse(in []int) {
    for i := 0; i < len(in)/2; i++ {
        in[i],in[len(in)-1-i] = in[len(in)-1-i], in[i]
    }
}