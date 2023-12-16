// sort and do equality check
func isAnagram(s string, t string) bool {
    sList := strings.Split(s,"")
    tList := strings.Split(t,"")
    sort.Strings(sList)
    sort.Strings(tList)
    return strings.Join(sList,"") == strings.Join(tList, "")
}