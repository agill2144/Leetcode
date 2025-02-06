func groupStrings(strings []string) [][]string {
    grps := map[string][]string{}    
    for i := 0; i < len(strings); i++ {
        key := hash(strings[i])
        grps[key] = append(grps[key],strings[i])
    }
    out := [][]string{}
    for _, v := range grps {
        out = append(out, v)
    }
    return out

}

func hash(word string) string {
    if len(word) == 1 {return ""}
    res := new(strings.Builder)
    for i := 1; i < len(word); i++ {
        diff := int(word[i-1]-'a') - int(word[i]-'a')
        if diff < 0 {diff += 26}
        res.WriteString(fmt.Sprintf("%v-", diff))
    }
    return res.String()
}