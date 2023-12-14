/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
  if (s.length % 2 !== 0) return false;

  const stack = [];

  for (chr of s) {
    switch (chr) {
      case "(":
        stack.push(chr);
        break;
      case "{":
        stack.push(chr);
        break;
      case "[":
        stack.push(chr);
        break;
      case ")":
        if (stack[stack.length - 1] === "(") {
          stack.pop();
          break;
        }
        return false;
      case "}":
        if (stack[stack.length - 1] === "{") {
          stack.pop();
          break;
        }
        return false;

      case "]":
        if (stack[stack.length - 1] === "[") {
          stack.pop();
          break;
        }
        return false;
    }
  }

  return stack.length === 0;
};
