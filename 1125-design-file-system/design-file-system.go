
type FileSystem struct {
    root *trieNode
}


func Constructor() FileSystem {
    return FileSystem{
        root: &trieNode{
            val: -1, 
            childrens: map[string]*trieNode{},
        },
    }
}


func (this *FileSystem) CreatePath(path string, value int) bool {
    return mkdir(this.root, path, value)
}


func (this *FileSystem) Get(path string) int {
    curr := this.root
    if len(path) == 1 {return -1}
    pathList := strings.Split(path, "/")
    for i := 1; i < len(pathList); i++ {
        val, ok := curr.childrens[pathList[i]]
        if !ok {return -1}
        curr = val
    }
    return curr.val
}

type trieNode struct {
    val int
    childrens map[string]*trieNode
}

// reusing from "Design In-Memory File System" question
func mkdir(root *trieNode, path string, val int) bool {
    curr := root
    if len(path) == 1 {return true}
    pathList := strings.Split(path, "/") // ["", "a", "b", "c"]
    // go to the second last dir
    for i := 1; i < len(pathList)-1; i++ {
        // if at any point a parent dir does not exist
        // we can exit early and return false
        if curr.childrens[pathList[i]] == nil {return false}
        // otherwise keep going down that path
        curr = curr.childrens[pathList[i]]
    }
    // if path already exists, return false
    if curr.childrens[pathList[len(pathList)-1]] != nil {return false}

    // successfully made to second last dir in path
    // create new node under 2nd last dir
    curr.childrens[pathList[len(pathList)-1]] = new(trieNode)
    curr.childrens[pathList[len(pathList)-1]].childrens = map[string]*trieNode{}
    curr.childrens[pathList[len(pathList)-1]].val = val
    return true
}


/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */