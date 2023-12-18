/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let unique = 1;

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] !== nums[unique - 1]) {
      nums[unique] = nums[i];
      unique++;
    }
  }
  return unique;
};
