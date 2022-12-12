func numberToWords(num int) string {
    
    below20 := []string{
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
    
    if num == 0 {return "Zero"}
 
    var helper func(tr int) string
    helper = func(tr int) string {
        // base
        if tr == 0 {return ""}
        
        // logic
        if tr < 20 {
            return below20[tr] + " "
        } else if tr >= 20 && tr <= 99 {
            return tens[tr/10] + " " + helper(tr%10) 
        } else { // >= 100
            return below20[tr/100] + " Hundred " + helper(tr%100)
        }
    }
    
    var result string
    // pointing to the thousands array
    i := 0
    for num > 0 {
        if num % 1000 != 0 {
            result = helper(num % 1000) + thousands[i] + " " + result
        }
        i++
        num = num / 1000
    }
    return strings.TrimSpace(result)
}
