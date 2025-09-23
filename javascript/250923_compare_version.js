var compareVersion = function(version1, version2) {
    const v1 = version1.split('.').map(Number);
    const v2 = version2.split('.').map(Number);

    const maxLength = Math.max(v1.length, v2.length);
    while (v1.length < maxLength) v1.push(0);
    while (v2.length < maxLength) v2.push(0);

    for (let i = 0; i < maxLength; i++) {
        if (v1[i] < v2[i]) return -1;
        if (v1[i] > v2[i]) return 1;
    }

    return 0;
};