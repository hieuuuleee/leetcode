const largestTriangleArea = (points) => {
  let max2 = 0;
  const n = points.length;
  for (let i = 0; i < n; ++i) {
    const [x1, y1] = points[i];
    for (let j = i + 1; j < n; ++j) {
      const [x2, y2] = points[j];
      for (let k = j + 1; k < n; ++k) {
        const [x3, y3] = points[k];
        const area2 = Math.abs((x1-x2)*(y1+y2)+(x2-x3)*(y2+y3)+(x3-x1)*(y3+y1))
        if (area2 > max2) max2 = area2;
      }
    }
  }
  return max2 / 2;
};
