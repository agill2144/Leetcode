# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # 1. create child to parent map ( all node values are uniq )
        node_to_parent = {root:root}
        def dfs(r: 'TreeNode'):
            nonlocal node_to_parent
            if not r:
                return
            
            if r.left:
                node_to_parent[r.left] = r
                dfs(r.left)
            if r.right:
                node_to_parent[r.right] = r
                dfs(r.right)
        dfs(root)
        
        # 2. if either p or q is not found, we exit
        if p not in node_to_parent or q not in node_to_parent:
            return None
        
        p_to_root = {}
        while p is not root:
            p_to_root[p] = True
            p = node_to_parent[p]
        
        while q is not root:
            if q in p_to_root:
                return q
            q = node_to_parent[q]
        if p == q:
            return p
        return None