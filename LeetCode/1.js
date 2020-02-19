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
let twoSum = function(nums, target) {
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
