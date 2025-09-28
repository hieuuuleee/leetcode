/**
 * @param {number[]} nums
 * @return {number}
 */
var largestPerimeter = function(nums) {
    let res=0;
    nums.sort((a,b) => a-b);
    for (let i=nums.length-1; i>1; i--)
        if (nums[i-2]+nums[i-1]>nums[i])
            return nums[i-2]+nums[i-1]+nums[i];
    return 0;
};
