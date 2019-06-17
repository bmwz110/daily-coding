var maximumGap = function(arr) {
  let len = arr.length
  if (len < 2) {
    return false
  }
  for (let i = len; i > 0; i--) {
    for (let j = 0; j < i; j++) {
      if (arr[j+1] < arr[j]) {
        let tmp = arr[j]
        arr[j] = arr[j+1]
        arr[j+1] = tmp
      }
    }
  }
  console.log(arr)
  let gap = 0
  for (let i = 0; i < len-1; i++) {
    let curGap = Math.abs(arr[i+1] - arr[i])
    if (curGap > gap) {
      gap = curGap
    }
  }
  return gap
};