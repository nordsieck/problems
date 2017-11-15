"use strict";
const test = require("../../test.js");

var lastRemaining = function(n) { return forward(n); };

function forward(n) {
    if (n === 1) { return 1; }
    return 2 * reverse(n / 2 | 0);
}

function reverse(n) {
    if (n === 1) { return 1; }
    if (n % 2 === 0) {
	return 2 * forward(n / 2 | 0) - 1;
    } else {
	return 2 * forward(n / 2 | 0);
    }
}

test.Run(
    test.test(lastRemaining(9), 6),
    test.test(lastRemaining(1000000000), 534765398),
);
