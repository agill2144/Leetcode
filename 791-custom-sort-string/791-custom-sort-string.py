class Solution:
    def customSortString(self, order: str, s: str) -> str:
        Map = dict()
        for i in range(len(s)):
            c = s[i]
            if c not in Map:
                Map[c] = 1
            else:
                Map[c] = Map[c] + 1
        strbuilder = ""
        for j in range(len(order)):
            c = order[j]
            if c in Map:
                count = Map.get(c,0)
                for k in range(count):
                    strbuilder += c
                Map.pop(c,None)
        #collect remaining characters
        for key, value in Map.items():
            count = Map[key]
            for k in range(count):
                strbuilder += key
        return strbuilder