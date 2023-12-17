/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function (nums) {
  const prefix = [];
  let sum = 0;
  for (const n of nums) {
    sum += n;
    prefix.push(sum);
  }

  for (let index = 0; index < prefix.length; index++) {
    const left = index - 1 >= 0 ? prefix[index - 1] : 0;
    const right = sum - prefix[index];
    if (left == right) return index;
  }

  return -1;
};
