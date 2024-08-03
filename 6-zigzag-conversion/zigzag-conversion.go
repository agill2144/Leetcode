func convert(s string, numRows int) string {
    if numRows == 1 { return s}
    out := new(strings.Builder)
    inc := (numRows-1)*2
    n := len(s)
    for i := 0; i < numRows; i++ {
        idx := i
        for idx < n {
            out.WriteByte(s[idx])
            if i != 0 && i != numRows-1 {
                inBetween := idx + (inc - 2*i)
                if inBetween < n {out.WriteByte(s[inBetween])}
            }
            idx += inc
        }
    }
    return out.String()
}