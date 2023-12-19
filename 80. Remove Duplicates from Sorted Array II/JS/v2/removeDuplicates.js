/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  if (nums.length <= 2) return nums.length;

  let unique = 2;

  for (let i = 2; i < nums.length; i++) {
    if (nums[i] !== nums[unique - 2]) {
      nums[unique] = nums[i];
      unique++;
    }
  }
  return unique;
};
