/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (numbers, target) {
  let index1 = 1;
  let index2 = numbers.length;
  while (index1 !== index2) {
    if (numbers[index1 - 1] + numbers[index2 - 1] < target) {
      index1++;
    } else if (numbers[index1 - 1] + numbers[index2 - 1] > target) {
      index2--;
    } else {
      return [index1, index2];
    }
  }
};
