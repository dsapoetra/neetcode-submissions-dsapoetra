class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        const res = new Map();
        for (let s of strs) {
            let sortedS = s.split('').sort().join('');

            if (!res[sortedS]) {
                res[sortedS] = []
            }

            res[sortedS].push(s);
        }
       return Object.values(res);
    }
}
