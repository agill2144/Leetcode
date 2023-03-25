/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */

// classic dfs :shrug
func crawl(startUrl string, htmlParser HtmlParser) []string {
    startUrlSplit := strings.Split(startUrl, "/")
    hostName := "http://"+startUrlSplit[2]
    result := map[string][]string{hostName: []string{}}
    visited := map[string]struct{}{}
    var dfs func(url string)
    dfs = func(url string) {
        // base
        if _, ok := visited[url]; ok {return}
        visited[url] = struct{}{}
        if !strings.HasPrefix(url, hostName) {
            return
        }
        
        // logic
        result[hostName] = append(result[hostName], url)
        urls := htmlParser.GetUrls(url)
        for _, u := range urls {
            dfs(u)
        }
    }
    dfs(startUrl)
    return result[hostName]
}