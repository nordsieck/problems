"use strict";
const test = require('../../test.js');

var compareVersion = function(version1, version2) {
    let [v1, v2] = [version1.split(".").map(pI), version2.split(".").map(pI)];
    while (v1[v1.length-1] === 0) { v1.pop(); }
    while (v2[v2.length-1] === 0) { v2.pop(); }

    for (let i = 0; i < v1.length && i < v2.length; i++) {
	if (v1[i] > v2[i]) { return 1; }
	if (v1[i] < v2[i]) { return -1; }
    }
    if (v1.length > v2.length) { return 1; }
    if (v1.length < v2.length) { return -1; }
    return 0;
}

function pI(i) { return parseInt(i, 10); }

test.Run(
    test.test(compareVersion("0.0.0", "0.0.0.1.0"), -1),
    test.test(compareVersion("01", "1"), 0),
    test.test(compareVersion("1", "0"), 1),
    test.test(compareVersion("0.10", "0.9"), 1),
    test.test(compareVersion("0.0.1", "0.0.2"), -1),
    test.test(compareVersion("0.0.0", "0.0.0.0"), 0),
    test.test(compareVersion("0", "0"), 0)
);
