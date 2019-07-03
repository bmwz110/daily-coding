// 给定一个整数数组 nums 和一个目标值 target，
// 请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
var twoSum = function(nums, target) {
  for (let i = 0; i<nums.length; i++) {
    for (let j = i+1; j < nums.length; j++) {
      if (target - nums[i] === nums[j]) {
        return [i,j]
      } else {
        continue
      }
    }
  }
};