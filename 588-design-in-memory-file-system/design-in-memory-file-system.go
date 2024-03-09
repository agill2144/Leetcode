type FileSystem struct {
    root *trieNode
}


func Constructor() FileSystem {
    return FileSystem{
        root: &trieNode{
            files: map[string]*strings.Builder{}, // {fileName: "fileContents"}
            childrens: map[string]*trieNode{},
        },
    }
}


func (this *FileSystem) Ls(path string) []string {
    curr := this.root
    pathList := strings.Split(path, "/")
    lastStr := pathList[len(pathList)-1]
    if len(path) > 1 {
        // path = "/a"
        for i := 1; i < len(pathList); i++ {
            if curr.childrens[pathList[i]] != nil {
                curr = curr.childrens[pathList[i]]
            }
        }
    }
    if curr.files != nil {
        if _, ok := curr.files[lastStr]; ok {
            return []string{lastStr}
        }
    }
    out := []string{}

    // add files
    for fileName, _ := range curr.files {
        out = append(out, fileName)
    }    
    // add other dirs
    for dirName, _ := range curr.childrens {
        out = append(out, dirName)
    }
    sort.Strings(out)
    return out

}


func (this *FileSystem) Mkdir(path string)  {
    _ = mkdir(this.root, path, false)
}


func (this *FileSystem) AddContentToFile(filePath string, content string)  {
    pathList := strings.Split(filePath, "/")
    fileName := pathList[len(pathList)-1]
    curr := mkdir(this.root,filePath, true)
    if curr.files == nil {
        curr.files = map[string]*strings.Builder{}
    }
    if curr.files[fileName] == nil {
        curr.files[fileName] = new(strings.Builder)
    }
    curr.files[fileName].WriteString(content)
}


func (this *FileSystem) ReadContentFromFile(filePath string) string {
    pathList := strings.Split(filePath, "/")
    fileName := pathList[len(pathList)-1]
    curr := this.root
    for i := 1; i < len(pathList)-1; i++ {
        curr = curr.childrens[pathList[i]]
    }
    return curr.files[fileName].String()
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */

type trieNode struct {
    files map[string]*strings.Builder
    childrens map[string]*trieNode
}

func mkdir(root *trieNode, path string, isFile bool) *trieNode {
    curr := root
    // path = "/"
    if len(path) == 1 {return curr}
    // path = "/a/b/c"
    pathList := strings.Split(path, "/") // ["", "a", "b", "c"]
    end := len(pathList)
    if isFile {end--}
    for i := 1; i < end; i++ {
        if curr.childrens[pathList[i]] == nil {
            curr.childrens[pathList[i]] = new(trieNode)
            curr.childrens[pathList[i]].files = map[string]*strings.Builder{}
            curr.childrens[pathList[i]].childrens = map[string]*trieNode{}
        }
        curr = curr.childrens[pathList[i]]
    }
    return curr
}