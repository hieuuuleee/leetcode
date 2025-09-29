/**
 * @param {number[]} values
 * @return {number}
 */
var minScoreTriangulation = function(values) {
    let n = values.length;
    let d = Array.from({length:n}, ()=>Array(n).fill(0));
    for (let len=3; len<=n; len++) {
        for (let i=0; i+len-1<n; i++) {
            let j = i+len-1;
            d[i][j] = Infinity;
            for (let k=i+1; k<j; k++) {
                d[i][j] = Math.min(d[i][j],
                    d[i][k] + d[k][j] + values[i]*values[j]*values[k]);
            }
        }
    }
    return d[0][n-1];
};
