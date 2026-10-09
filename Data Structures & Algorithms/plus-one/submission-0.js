class Solution {
    /**
     * @param {number[]} digits
     * @return {number[]}
     */
    plusOne(digits) {
        let digitStr = digits.join('');
        // console.log(digitStr)
        let digitNums = parseInt(digitStr);
        digitNums++;

        let ret = digitNums.toString();
        ret.split('');

        return ret
    }
}
