# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', nodes: 'List[TreeNode]') -> 'TreeNode':
        nodeSet = set(nodes)
        def dfs(r: TreeNode) -> TreeNode:
            # base
            if r == None:
                return None

            # logic
            if r in nodeSet:
                return r
            left = dfs(r.left)
            right = dfs(r.right)
            if left != None and right != None:
                return r
            if left != None:
                return left
            return right

        return dfs(root)