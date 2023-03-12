func getDirections(root *TreeNode, startValue int, destValue int) string {
    /*
        Get LCA (lowest common ancestor for start & dest)
        Get start & dest path from LCA 
        Combine the path
    */

    lca := lowestCommonAncestor(root, startValue, destValue)
    startPath, destPath := []string{}, []string{}
    dfs(lca, startValue, &startPath)
    dfs(lca, destValue, &destPath)

    return invert(startPath)+strings.Join(destPath,"")
}

func invert (path []string) string{
    //sLen := len(path)
	var upPath strings.Builder
    for _=range path{
        upPath.WriteString("U")
    }
    return upPath.String()
}

func dfs(cur *TreeNode, val int, path *[]string) bool{
    if cur == nil{
        return false
    }

    if cur.Val == val {
        return true
    }

    *path = append(*path, "L")
    hasValue := dfs(cur.Left, val, path)
    if hasValue {
        return true
    }
    *path = (*path)[:len(*path)-1]

    *path = append(*path, "R")
    hasValue = dfs(cur.Right, val, path)
    if hasValue {
        return true
    }
    *path = (*path)[:len(*path)-1]

    return false

}

func lowestCommonAncestor(root *TreeNode, s,d int) *TreeNode{

    if root == nil{
        return nil
    }

    if root.Val == s || root.Val == d {
        return root
    }

    left := lowestCommonAncestor(root.Left, s, d)
    right := lowestCommonAncestor(root.Right, s, d)

    if left!=nil && right!=nil{
        return root
    }

    if left == nil {
        return right
    }

    if right == nil{
        return left
    }

    return nil

}
