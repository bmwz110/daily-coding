// 方法1 - split
var reverseWords = function(s) {
  return s.split(' ').map(item => {
      return item.split('').reverse().join('')
  }).join(' ')
};

// // 方法2 - 正则
// var reverseWords = function(s) {
//   return s.split(/\s/g).map(item => {
//       return item.split('').reverse().join('')
//   }).join(' ')
// };

// // 方法3 - match
// var reverseWords = function(s) {
//   return s.match(/[\w']+/g).map(item => {
//       return item.split('').reverse().join('')
//   }).join(' ')
// };