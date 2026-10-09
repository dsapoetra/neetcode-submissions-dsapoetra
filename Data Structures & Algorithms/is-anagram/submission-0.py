class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        s = sorted(s)
        t = sorted(t)

        for x in range(len(s)):
            if s[x] != t[x]:
                return False
        return True
