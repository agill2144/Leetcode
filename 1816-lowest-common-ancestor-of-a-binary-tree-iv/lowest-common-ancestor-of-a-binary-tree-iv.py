# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', nodes: 'List[TreeNode]') -> 'TreeNode':
        nodeSet = set(nodes)
        def dfs(r :'TreeNode'):
            # base
            if not r:
                return None

            # logic
            if r in nodeSet:
                return r
            left = dfs(r.left)
            right = dfs(r.right)
            if left and right:
                return r
            if left:
                return left
            return right
        return dfs(root)