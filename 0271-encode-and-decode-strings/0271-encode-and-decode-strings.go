/*
    approach: using a non ascii character as a delimeter
    - The input strs will only ever contain 256 ascii values
    - there are non-ascii values too ( 128 )
    - we can use one of the non ascii value as a delimeter
    - in encode we will go thru the list of strings and append each to a stringBuilder
    - and add a non-ascii character in between
    - in decode, we will split our string on that non-ascii delimeter so that can we can restore back the original input
    
    - in golang, the max/last ascii value can be retrieved by unicode.MaxASCII const
    - to get a non-ascii value, unicode.MaxASCII+1
*/
type Codec struct {
    data string
}

// Encodes a list of strings to a single string.
// time: o(n) to loop over all strs
// space: o(1) since string builder is used as output, that does not count
func (codec *Codec) Encode(strs []string) string {
    out := new(strings.Builder)
    for i := 0; i < len(strs); i++ {
        out.WriteString(strs[i])
        if i != len(strs)-1 {
            out.WriteRune(unicode.MaxASCII+1)
        }
    }
    return out.String()
}

// Decodes a single string to a list of strings.
// time: o(n) - split under the hood is going character by character and as soon as it finds the delimeter it splits ( does that at each delimeter )
// space: o(1)
func (codec *Codec) Decode(strs string) []string {
    return strings.Split(strs, string(unicode.MaxASCII+1))
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));