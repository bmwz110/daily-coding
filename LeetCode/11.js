// my answer
// o(n2)
var maxArea = function(height) {
  let area = 0;
  for (let i = 0; i < height.length; i++) {
    for (let j = i; j < height.length; j++) {
      let newArea = (Math.min.apply(null, [height[i],height[j]]) * Math.abs(i - j));
      if (newArea > area) {
        area = newArea;
      }
    }
  }
  return area;
    
};

console.log(maxArea([1,8,6,2,5,4,8,3,7]));


// opt answer
// o(n)
var maxArea = function(height) {
  let area = 0;
  let iniI = 0;
  let iniJ = height.length - 1;
  while (iniI < iniJ) {
    newArea = Math.min.apply(null ,[height[iniI] ,height[iniJ]]) * Math.abs(iniI - iniJ);
    if (area < newArea) {
      area = newArea;
    }

    if (height[iniI] < height[iniJ]) {
      iniI ++;
    } else {
      iniJ --;
    }
  }
  return area;
};

console.log(maxArea([1,8,6,2,5,4,8,3,7]));
