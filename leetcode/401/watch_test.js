"use strict";
const test = require("../../test.js");

const times = [1, 2, 4, 8, 16, 32, // minutes
	       60, 120, 240, 480]; // hours

var readBinaryWatch = function(num) {
    counters = [];
    for (let i = 0; i < num; i++) { counters[i] = i; }
};

// TODO
function increment(counters, times) {
    counters[counters.length - 1] += 1;
    for (let i = counters.length - 1; i >= 0; i--) {
	if (counters[i] == times.length - 1) {
	    counters[i - 1] += 1;
	    counters[i]
	}
    }
}

function format(idx) {
    let m = 0, h = 0;
    for (let i in idx) {
	let val = times[idx[i]];
	if (val < 60) { m += val; } else { h += val; }
    }
    return String(h/60) + ":" + (m < 10 ? "0" : "") + String(m);
}

function legal(idx) {
    if (idx.length == 0) { return false; }
    let m = 0, h = 0;
    for (let i in idx) {
	let val = times[idx[i]];
	if (val < 60) { m += val; } else { h += val; }
    }
    if (m >= 60) { return false; }
    if (h > 12 * 60) { return false; }
    return true;
}

test.Run(
    test.test(format([0, 1, 3, 4, 5, 8, 9]), "12:59"),
    test.test(format([6]), "1:00"),
    test.test(format([0]), "0:01"),
    test.test(format([]), "0:00"),
    
    test.test(legal([0, 1, 3, 4, 5, 8, 9]), true),
    test.test(legal([8, 9]), true),
    test.test(legal([0, 1, 3, 4, 5]), true),
    test.test(legal([6, 7, 8, 9]), false),
    test.test(legal([0, 1, 2, 3, 4, 5]), false),
    test.test(legal([0, 6]), true),
    test.test(legal([6]), true),
    test.test(legal([0]), true),
    test.test(legal([]), false)
);
