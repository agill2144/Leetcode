# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution(object):
    def pathSum(self, root, targetSum):
        """
        :type root: TreeNode
        :type targetSum: int
        :rtype: List[List[int]]
        """
        result = []
        self.inorderDfs(root, targetSum, [], result)
        return result
    
    def inorderDfs(self, root, target, paths, result):
        if root is None:
            return
        
        target = target - root.val
        paths.append(root.val)
        
        self.inorderDfs(root.left, target, paths, result)
        if target == 0 and root.left is None and root.right is None:
            result.append(list(paths))
        
        self.inorderDfs(root.right, target, paths, result)
        paths.pop()
