# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        p_path = []
        q_path = []
        def dfs(r, path):
            nonlocal p_path
            nonlocal q_path
            if not r:
                return
            
            path.append(r)
            if r is p:
                new_p = path.copy()
                p_path = new_p
                
            if r is q:
                new_q = path.copy()
                q_path = new_q
                
            if len(p_path) > 0 and len(q_path) > 0:
                return
            dfs(r.left, path)
            dfs(r.right, path)
            path.pop()
        dfs(root, [])
        if len(p_path) == 0 or len(q_path) == 0:
            return None
        
        out = None
        p_ptr = 0
        q_ptr = 0
        while p_ptr < len(p_path) and q_ptr < len(q_path):
            if p_path[p_ptr] == q_path[q_ptr]:
                out = p_path[p_ptr]
                p_ptr += 1
                q_ptr += 1
                continue
            break
        return out

    # def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
    #     # 1. create child to parent map ( all node values are uniq )
    #     node_to_parent = {root:root}
    #     def dfs(r: 'TreeNode'):
    #         nonlocal node_to_parent
    #         if not r:
    #             return
            
    #         if r.left:
    #             node_to_parent[r.left] = r
    #             dfs(r.left)
    #         if r.right:
    #             node_to_parent[r.right] = r
    #             dfs(r.right)
    #     dfs(root)
        
    #     # 2. if either p or q is not found, we exit
    #     if p not in node_to_parent or q not in node_to_parent:
    #         return None
        
    #     p_to_root = {}
    #     while p is not root:
    #         p_to_root[p] = True
    #         p = node_to_parent[p]
        
    #     while q is not root:
    #         if q in p_to_root:
    #             return q
    #         q = node_to_parent[q]
    #     if p == q:
    #         return p
    #     return None