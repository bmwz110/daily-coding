// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

// 示例
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]


// my answer
let twoSum = function(nums, target) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i+1; j < nums.length; j++) {
      if ((nums[i] + nums[j]) === target) {
        return [i, j]
      }
    }
  }
  return false
};

// opt answer
// HashMap
var twoSum = function(nums, target) {
  let map = new Map();
  for (let i = 0; i < nums.length; i++) {
    let gapValue = target - nums[i];
    if (map.get(gapValue) !== undefined) {
        return [map.get(gapValue), i];
    } else {
        map.set(nums[i], i);
    }
  }
  return false
};
