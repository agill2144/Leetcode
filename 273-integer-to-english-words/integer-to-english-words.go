func numberToWords(num int) string {
    if num == 0 {return "Zero"}
    belowTwentys := []string {
        "",
        "One","Two","Three",
        "Four","Five","Six",
        "Seven","Eight","Nine",
        "Ten","Eleven","Twelve",
        "Thirteen","Fourteen","Fifteen",
        "Sixteen","Seventeen","Eighteen","Nineteen",
    }
    tens := []string{
        "",
        "Ten","Twenty","Thirty",
        "Forty", "Fifty","Sixty",
        "Seventy","Eighty","Ninety",
    }
    thousands := []string{
        "", "Thousand", "Million", "Billion",
    }

    var helper func(n int) string
    helper = func(n int) string {
        // base
        if n == 0 {return ""}

        // logic
        if n < 20 {
            return belowTwentys[n] + " "
        } else if n < 100 {
            return tens[n/10] + " " + helper(n%10)
        } else {
            // 123 = 23
            return belowTwentys[n/100] + " Hundred " + helper(n%100)
        }
    }

    i := 0
    ans := ""
    for num != 0 {

        //1. get a triplet from the back
        // example: 123456; we want 456
        // thats num % 1000
        triplet := num % 1000

        //2. process a triplet
        // we are processing each triplet pair from right to left
        if triplet != 0 {ans = helper(triplet) + thousands[i] + " " + ans}

        //3. update num to be the remaining number
        // after processing tripplet
        // example: 123456 = 123
        // thats num / 1000
        num /= 1000
        i++
    }
    return strings.TrimSpace(ans)
}

