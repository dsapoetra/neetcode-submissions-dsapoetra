class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums) {
        const map = new Map();

        for (let i = 0; i < nums.length; i++) {

            const isTarget = map.has(nums[i]);

            if (isTarget) {
                return true;
            }

            map.set(nums[i], i)
        }

        return false
    }
}
