class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        s = s.replace('-', '')
        r = len(s) % k
        strs = []
        if r != 0:
            strs = [s[:r].upper()]

        while r < len(s):
            strs.append(s[r:r+k].upper())
            r += k

        return '-'.join(strs)
