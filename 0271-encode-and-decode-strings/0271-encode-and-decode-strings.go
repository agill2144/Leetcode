type Codec struct {
    data string
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
    out := new(strings.Builder)
    for i := 0; i < len(strs); i++ {
        out.WriteString(strs[i])
        if i != len(strs)-1 {
            out.WriteString(string(unicode.MaxASCII+150))
        }
    }
    codec.data = out.String()
    return codec.data
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
    deli := string(unicode.MaxASCII+150)
    return strings.Split(codec.data, deli)
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));