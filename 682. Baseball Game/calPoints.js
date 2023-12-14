/**
 * @param {string[]} operations
 * @return {number}
 */
var calPoints = function (operations) {
  const scoreRecord = [];
  let lastPtr = -1;

  operations.forEach((op) => {
    switch (op) {
      case "+":
        scoreRecord.push(scoreRecord[lastPtr] + scoreRecord[lastPtr - 1]);
        lastPtr++;
        break;
      case "D":
        scoreRecord.push(scoreRecord[lastPtr] * 2);
        lastPtr++;
        break;
      case "C":
        scoreRecord.pop();
        lastPtr--;
        break;
      default:
        scoreRecord.push(op * 1);
        lastPtr++;
    }
  });

  if (scoreRecord.length === 0) {
    return 0;
  }

  return scoreRecord.reduce((acc, cur) => acc + cur);
};
