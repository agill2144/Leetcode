func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    left := 0
    right := len(products)-1
    str := new(strings.Builder)
    out := [][]string{}
    for i := 0; i < len(searchWord); i++ {
        str.WriteByte(searchWord[i])
        
        for left < len(products) && !strings.HasPrefix(products[left],str.String()) {left++}
        for right >= 0 && !strings.HasPrefix(products[right],str.String()) {right--}
        if left >= len(products) || right < 0 {out = append(out, []string{}); continue}
        
        if right-left+1 > 3 {
            out = append(out, products[left:right+1][:3])         
        } else {
            out = append(out, products[left:right+1])
        }
    }
    return out
}