/**
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const prefix = [];
  const postfix = [];
  const answer = [];
  let preProduct = 1;
  let postProduct = 1;

  // get prefix, postfix sum (product)
  // nums.length is at least 2, there is no need to check out of bound
  for (let i = 0, j = -1; i < nums.length; i++, j--) {
    preProduct *= nums[i];
    postProduct *= nums.at(j);
    prefix.push(preProduct);
    postfix.push(postProduct);
  }

  for (let i = 0; i < nums.length; i++) {
    const pre = i - 1 >= 0 ? prefix[i - 1] : 1;
    const post =
      nums.length - 1 - 1 - i >= 0 ? postfix[nums.length - 2 - i] : 1;
    answer.push(pre * post);
  }

  return answer;
};
