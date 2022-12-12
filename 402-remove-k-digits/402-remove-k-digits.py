class Solution:
    def removeKdigits(self, num: str, k: int) -> str:
        # null case
        if num == "":
            return num
        if len(num) == k:
            return "0"
        result = ""
        n = len(num)
        stack = []
        stack.append(int(num[0]))
        for i in range(1, len(num)):
            incoming =int(num[i])
            
            if stack: #check stack.top
                while(stack and incoming < stack[-1] and k):
                    stack.pop()
                    k -= 1
                    print(len(stack))
            stack.append(incoming)
            #print(str(stack))
        
        while k and stack:
            stack.pop()
            k -= 1

        
        for i in range(len(stack)):
            elem = stack.pop()
            result = str(elem) + result
        
        return str(int(result))