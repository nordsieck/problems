"use strict";
const test = require("../../test.js");

var integerReplacement = function(n) {
    let steps = 0;
    while (n > 1) {
	if (n % 2 === 0 ) { n /= 2; }
	else if (n == 3 || n % 4 == 1) { n--; }
	else { n++; }
	steps++;
    }
    return steps;
};

test.Run(
    test.test(integerReplacement(8), 3),
    test.test(integerReplacement(7), 4),
);
