/** Pair class to store key-value pairs */
// class Pair {
//   /**
//    * @param {number} key The key to be stored in the pair
//    * @param {string} value The value to be stored in the pair
//    */
//   constructor(key, value) {
//       this.key = key;
//       this.value = value;
//   }
// }
class Solution {
  /**
   * @param {Pair[]} pairs
   * @returns {Pair[]}
   */
  mergeSort(pairs) {
    return this.mergeSortHelper(pairs, 0, pairs.length - 1);
  }

  mergeSortHelper(pairs, start, end) {
    if (end - start <= 0) return pairs;
    const mid = Math.floor((start + end) / 2);
    this.mergeSortHelper(pairs, start, mid);
    this.mergeSortHelper(pairs, mid + 1, end);

    this.merge(pairs, start, mid, end);

    return pairs;
  }

  merge(arr, start, mid, end) {
    const left = arr.slice(start, mid + 1);
    const right = arr.slice(mid + 1, end + 1);

    let [i, j, k] = [0, 0, start];
    while (i < left.length && j < right.length) {
      if (left[i].key <= right[j].key) {
        arr[k] = left[i];
        i++;
      } else {
        arr[k] = right[j];
        j++;
      }
      k++;
    }
    while (i < left.length) {
      arr[k] = left[i];
      i++;
      k++;
    }
    while (j < right.length) {
      arr[k] = right[j];
      j++;
      k++;
    }
  }
}
