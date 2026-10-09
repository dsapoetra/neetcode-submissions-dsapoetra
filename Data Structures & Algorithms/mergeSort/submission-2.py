# Definition for a pair.
# class Pair:
#     def __init__(self, key: int, value: str):
#         self.key = key
#         self.value = value
class Solution:
    def mergeSort(self, pairs: List[Pair]) -> List[Pair]:
        return self.mergeSortHelper(pairs, 0, len(pairs) - 1)

    def mergeSortHelper(self, pairs: List[Pair], s: int, e: int) -> List[Pair]:
        if e - s + 1 <= 1:
            return pairs
        m = (s + e) // 2

        self.mergeSortHelper(pairs, s, m)
        self.mergeSortHelper(pairs, m + 1, e)

        self.merge(pairs, s, m, e)

        return pairs

    def merge(self, arr: List[Pair], s: int, m: int, e: int) -> None:
        L = arr[s:m+1]
        R = arr[m+1: e+1]

        pointerL = 0
        pointerR = 0
        pointerRes = s

        while pointerL < len(L) and pointerR < len(R):
            if L[pointerL].key <= R[pointerR].key:
                arr[pointerRes] =L[pointerL]
                pointerL += 1
            else:
                arr[pointerRes] = R[pointerR]
                pointerR += 1
            pointerRes += 1

        while pointerL < len(L):
            arr[pointerRes] = L[pointerL]
            pointerL += 1
            pointerRes += 1

        while pointerR < len(R):
            arr[pointerRes] = R[pointerR]
            pointerR += 1
            pointerRes += 1