class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        word_set = set(wordDict)
        out = []

        def dfs(start: int, path: List[str]):
            # base case
            if start == len(s):
                path_str = " ".join(path)
                out.append(path_str)
                return

            # recursive case
            for i in range(start, len(s)):
                sub_str = s[start:i + 1]
                if sub_str in word_set:
                    path.append(sub_str)
                    dfs(i + 1, path)
                    path.pop()

        dfs(0, [])
        return out
