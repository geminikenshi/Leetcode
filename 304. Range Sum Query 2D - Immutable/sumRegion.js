/**
 * @param {number[][]} matrix
 */
var NumMatrix = function (matrix) {
  this.prefixSum = [];
  for (const row of matrix) {
    let total = 0;
    const sum = [];
    for (const n of row) {
      total += n;
      sum.push(total);
    }
    this.prefixSum.push(sum);
  }
};

/**
 * @param {number} row1
 * @param {number} col1
 * @param {number} row2
 * @param {number} col2
 * @return {number}
 */
NumMatrix.prototype.sumRegion = function (row1, col1, row2, col2) {
  let regionSum = 0;
  for (const i = row1; i <= row2; i++) {
    regionSum +=
      col1 - 1 >= 0
        ? this.prefixSum[i][col2] - this.prefixSum[i][col1 - 1]
        : this.prefixSum[i][col2];
  }
  return regionSum;
};

/**
 * Your NumMatrix object will be instantiated and called as such:
 * var obj = new NumMatrix(matrix)
 * var param_1 = obj.sumRegion(row1,col1,row2,col2)
 */
