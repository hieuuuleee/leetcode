/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let n = height.length;
    let res = 0;
    let i = 0, j = n - 1;
    while (i !== j) {
        let area = (j - i) * Math.min(height[i], height[j]);
        res = Math.max(area, res);
        if (height[i] <= height[j]) {
            i++;
        } else {
            j--;
        }
    }
    return res;
};
