// cur记录当前连续的字符数，last记录字符变更之前连续的字符数
// 若last >= cur，例如 01、001、0001，则必有一个解
var countBinarySubstrings = function(s) {
  s = s.split('')
  let last, cur, res
  last = res = 0
  cur = 1
  for (let i = 1; i < s.length; i++) {
    if (s[i] === s[i-1]) {
      cur ++
    } else {
      last = cur
      cur = 1
    }

    if (last >= cur) {
      res ++
    }
  }
  return res
};