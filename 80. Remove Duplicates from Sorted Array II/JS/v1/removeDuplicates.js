/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
  let twice = false;
  let unique2 = 1;
  let left = 1;
  for (let i = 1; i < nums.length; i++) {
    if (nums[i] === nums[left]) {
      if (!twice) {
        nums[left] = nums[i];
        left++;
        twice = true;
        unique2++;
      }
    } else {
      nums[left] = nums[i];
      left++;
      twice = false;
      unique2++;
    }
  }

  return unique2;
};
