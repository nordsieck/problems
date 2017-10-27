"use strict";
const test = require("../../test.js");

function parse(s) {
    let split = s.split("+");
    split[1] = split[1].slice(0, split[1].length - 1);
    return [parseInt(split[0]), parseInt(split[1])];
}

function format(i) { return i[0].toString() + "+" + i[1].toString() + "i"; }

var complexNumberMultiply = function(a, b) {
    let aa = parse(a);
    let bb = parse(b);
    let cc = [aa[0] * bb[0] - aa[1] * bb[1], aa[0] * bb[1] + aa[1] * bb[0]];
    return format(cc);
};

test.Run(
    test.test(parse("1+1i"), [1, 1]),
    test.test(parse("0+-1i"), [0, -1]),
    test.test(format([1, 1]), "1+1i"),
    test.test(format([0, -1]), "0+-1i"),
    test.test(complexNumberMultiply("1+1i", "1+1i"), "0+2i"),
);
