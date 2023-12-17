/**
 * @param {number[]} nums
 */
var NumArray = function (nums) {
  this.prefix = [];
  let total = 0;
  for (const n of nums) {
    total += n;
    this.prefix.push(total);
  }
};

/**
 * @param {number} left
 * @param {number} right
 * @return {number}
 */
NumArray.prototype.sumRange = function (left, right) {
  return left - 1 >= 0
    ? this.prefix[right] - this.prefix[left - 1]
    : this.prefix[right];
};

/**
 * Your NumArray object will be instantiated and called as such:
 * var obj = new NumArray(nums)
 * var param_1 = obj.sumRange(left,right)
 */
