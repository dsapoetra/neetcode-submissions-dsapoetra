class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    missingNumber(nums) {
        let res = nums.length;

        for (let i = 0; i < nums.length; i++) {
            console.log("===")
            // console.log(res);
            console.log(i);
            console.log(nums[i]);
            console.log("===")

            res += i - nums[i];
        }

        return res;
    }
}
