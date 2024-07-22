func sortPeople(names []string, heights []int) []string {
    heightToPerson := map[int]string{}
    for i := 0; i < len(heights); i++ {
        heightToPerson[heights[i]] = names[i]
    }
    sort.Slice(heights, func(i, j int)bool{return heights[i] > heights[j]})
    out := []string{}
    for i := 0; i < len(heights); i++ {
        out = append(out, heightToPerson[heights[i]])
    }
    return out
}
// func sortPeople(names []string, heights []int) []string {
//     type mergedNode struct {
//         person string
//         height int
//     }
//     merged := []*mergedNode{}
//     for i := 0; i < len(names); i++ {
//         merged = append(merged, &mergedNode{names[i], heights[i]})
//     }
//     sort.Slice(merged, func(i, j int)bool{
//         return merged[i].height > merged[j].height
//     })
//     out := []string{}
//     for i := 0; i < len(merged); i++ {
//         out = append(out, merged[i].person)
//     }
//     return out
// }