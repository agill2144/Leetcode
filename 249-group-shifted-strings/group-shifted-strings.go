func groupStrings(strings []string) [][]string {
    groups := map[string][]string{}
    for i := 0; i < len(strings); i++ {
        key := hash(strings[i])
        groups[key] = append(groups[key], strings[i])
    }
    out := [][]string{}
    for _, v := range groups {
        out = append(out, v)
    }
    return out
}

func hash(word string) string {
    if len(word) == 1 {return "0"} 
    res := new(strings.Builder)
    for i := 1; i < len(word); i++ {
        diff := int(word[i]-'a') - int(word[i-1]-'a')
        if diff < 0 {diff += 26}
        res.WriteString(fmt.Sprintf("%v-", diff))
    }
    return res.String()
}