"use strict";
const test = require("../../test.js");

var findComplement = function(num) {
    return fromBits(reverse(toBits(num)));
};

function toBits(num) {
    if (num === 0) { return [0]; }
    
    let bits = [];
    while (num > 0) {
	bits.push(num % 2);
	num = Math.floor(num / 2);
    }
    return bits;
}

function fromBits(bits) {
    let mul = 1;
    let num = 0;

    while (bits.length > 0) {
	num = num + mul * bits[0];
	mul = mul * 2;
	bits.shift();
    }
    return num;
}

function reverse(bits) {
    bits = bits.map((x) => (x + 1) % 2);
    while (bits[bits.length - 1] === 0) {
	bits.length--;
    }
    if (bits.length === 0) { return [0]; }
    return bits;
}

test.Run(
    test.test(toBits(0), [0]),
    test.test(toBits(1), [1]),
    test.test(toBits(2), [0, 1]),
    test.test(toBits(3), [1, 1]),
    test.test(fromBits([0]), 0),
    test.test(fromBits([1]), 1),
    test.test(fromBits([0, 1]), 2),
    test.test(fromBits([1, 1]), 3),
    test.test(reverse([0]), [1]),
    test.test(reverse([1]), [0]),
    test.test(reverse([0, 1]), [1]),
    test.test(reverse([1, 1]), [0]),
    test.test(reverse([0, 0, 1]), [1, 1]),
    test.test(findComplement(5), 2),
    test.test(findComplement(1), 0),
);
