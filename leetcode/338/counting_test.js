"use strict";
const test = require("../../test.js");

var countBits = function(num) {
    let bits = [0];

    while (bits.length <= num) {
	bits = bits.concat(bits.map((x) => x+1));
    }
    return bits.slice(0, num+1);
};

test.Run(
    test.test(countBits(0), [0]),
    test.test(countBits(1), [0, 1]),
    test.test(countBits(2), [0, 1, 1]),
    test.test(countBits(3), [0, 1, 1, 2]),
);
