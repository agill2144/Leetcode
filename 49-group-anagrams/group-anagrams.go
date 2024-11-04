func groupAnagrams(strs []string) [][]string {
    groups := map[float64][]string{}
    for i := 0; i < len(strs); i++ {
        hashVal := hash(strs[i])
        groups[hashVal] = append(groups[hashVal], strs[i])
    }
    out := [][]string{}
    for _, v := range groups { out = append(out, v)}
    return out
}

func hash(key string) float64 {
    primes := []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    prod := 1.0
    for i := 0; i < len(key); i++ {
        idx := int(key[i] - 'a')
        prod *= primes[idx]
    }
    return prod
}