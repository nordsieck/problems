"use strict";
const test = require('../../test.js');

var hammingDistance = function(x, y) { return bits(x ^ y); };

function bits(num) {
    let count = 0;
    while (num > 0) {
	if (num % 2 != 0) { count++; num -= 1; }
	num /= 2;
    }
    return count;
};

test.Run(
    test.test(bits(0), 0),
    test.test(bits(1), 1),
    test.test(bits(3), 2),
    test.test(bits(1024), 1),
    test.test(bits(1023), 10),

    test.test(hammingDistance(1, 4), 2),
    test.test(hammingDistance(1024, 1024), 0),
);
