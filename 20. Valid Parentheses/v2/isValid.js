/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  const closeParenMap = {
    ")": "(",
    "}": "{",
    "]": "[",
  };
  const stack = [];

  for (const ch of s) {
    if (!(ch in closeParenMap)) {
      stack.push(ch);
    } else {
      if (stack.at(-1) === closeParenMap[ch]) {
        stack.pop();
        continue;
      }
      return false;
    }
  }

  return stack.length === 0;
};
