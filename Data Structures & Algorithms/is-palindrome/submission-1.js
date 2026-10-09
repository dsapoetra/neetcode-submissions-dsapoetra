class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */

    isAlphaNumeric(char) {
        return (char >= 'a' && char <= 'z') ||
               (char >= 'A' && char <= 'Z') ||
               (char >= '0' && char <= '9');
    }

    isPalindrome(s) {
        let res = ''

        for (let c of s) {
            if (this.isAlphaNumeric(c)) {
                res += c.toLowerCase();
            }
        }

        return res === res.split('').reverse().join('');
    }
}
