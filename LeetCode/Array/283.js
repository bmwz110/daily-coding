// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:
// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]

// test data:
arr = [1, 0, 2, 3, 0 ,5];


// best answer
// O(n) O(1)
function sortArr(arr) {
  let i = 0;

  // 第一次遍历的时候，j指针记录非0的个数，只要是非0的统统都赋给nums[j]
  for (let j = 0; j < arr.length; j++) {
    if (arr[j] !== 0) {
      arr[i++] = arr[j];
    }
  }

  // 非0元素统计完了，剩下的都是0了,第二次遍历把末尾的元素都赋为0即可
  for (let j = i; j < arr.length; j++) {
    arr[j] = 0;
  }

  return arr;
}



// best answer - like quick sort
function sortArr(arr) {
  let i = 0;

  for (let j = 0; j < arr.length; j++) {
    // if 当前元素!=0，就把其交换到左边，等于0的交换到右边
    if (arr[j] !== 0) {
      [arr[i++], arr[j]] = [arr[j], arr[i]];
    }
  }

  return arr;
}

console.log(sortArr(arr));
