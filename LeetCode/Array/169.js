// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1:
// 输入: [3,2,3]
// 输出: 3

// test data:
let arr = [1, 1, 2, 2, 2, 5, 5, 5, 5, 5, 5];

// origin answer - Violence
// O(n2) O(1)
function majorityEle(arr) {
  let majorityCount = arr.length / 2;

  for (let i = 0; i < arr.length; i++) {
    let count = 0;
    for (let j = 0; j < arr.length; j++) {
      if (arr[i] === arr[j]) {
        count += 1;
      }
    }

      if (count > majorityCount) {
        return arr[i];
    }
  }

  return false;
}



// best answer - Hash Map
// O(n) O(n)
function majorityEle(arr) {
  let tempMap = new Map();

  for (let i = 0; i < arr.length; i++) {
    if (tempMap.get(arr[i]) === undefined) {
      tempMap.set(arr[i], 1);
    } else {
      tempMap.set(arr[i], tempMap.get(arr[i]) + 1)
    }
  }

  let maxCount = 0;
  let finalItem = 0;
  for (let item of tempMap.keys()) {
    if (tempMap.get(item) > maxCount) {
      maxCount = tempMap.get(item);
      finalItem = item;
    } else {
      maxCount = maxCount;
    }
  }
  return finalItem;
}



// best answer - sort
// O(nlogn) O(1)/O(n) 取决于是否能就地排序
// 注意：本题定义的众数为出现次数大于 n/2 ，故排序后定为 ceil(n/2) 处数字
function majorityEle(arr) {
  arr.sort(function(a, b) {
    return a - b;
  });

  return arr[Math.ceil(arr.length/2)];
}



// opt answer - Boyer-Moore vote algorithm
// 通过 +1 和 -1 计数，删掉同样数目的众数和非众数
// 注意：本题定义的众数为出现次数大于 n/2 
function majorityEle(arr) {
  let count = 0;
  let candidate = null;

  for (let i = 0; i < arr.length; i++) {
    if (count === 0) {
      candidate = arr[i];
    }

    count += (arr[i] === candidate) ? 1 : -1;
  }

  return candidate;
}


console.log(majorityEle(arr));

