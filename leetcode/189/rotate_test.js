"use strict";
const test = require('../../test.js');

var rotate = function(nums, k) {
    if (k === 0) { return; }
    for (let i = 0, g = gcf(nums.length, k); i < g; i++) { rotateSlice(nums, k, i); }
};

function rotateSlice(nums, k, idx) {
    if (k === 0) { return; }
    
    let [val, i, next] = [nums[idx], idx, 0];
    while (true) {
	next = (i + k) % nums.length;
	[val, nums[next]] = [nums[next], val];
	if (next == idx) { break; }
	i = next;
    }
}

function gcf(a, b) {
    let [g, i] = [1, 2];
    for (let i = 2; i <= a && i <= b; i++) {
	while (a % i === 0 && b % i === 0) { a /= i; b /= i; g *= i; }
    }
    return g;
}

function resultSlice(nums, k, idx) { rotateSlice(nums, k, idx); return nums; }
function result(nums, k) { rotate(nums, k); return nums; }

test.Run(
    test.test(gcf(49, 21), 7),
    test.test(gcf(25, 15), 5),
    test.test(gcf(2, 4), 2),
    test.test(gcf(2, 2), 2),
    test.test(gcf(1, 7), 1),
    test.test(gcf(1, 1), 1),

    test.test(resultSlice([1, 2, 3], 2, 0), [2, 3, 1]),
    test.test(resultSlice([1, 2, 3], 1, 2), [3, 1, 2]),
    test.test(resultSlice([1, 2, 3], 1, 1), [3, 1, 2]),
    test.test(resultSlice([1, 2, 3], 1, 0), [3, 1, 2]),
    test.test(resultSlice([1, 2, 3, 4], 2, 0), [3, 2, 1, 4]),
    test.test(resultSlice([1, 2, 3, 4], 1, 0), [4, 1, 2, 3]),
    test.test(resultSlice([1], 0, 0), [1]),
    test.test(resultSlice([], 0, 0), []),
    
    test.test(result([1, 2, 3, 4, 5, 6, 7], 3), [5, 6, 7, 1, 2, 3, 4]),
    test.test(result([1, 2, 3, 4], 2), [3, 4, 1, 2]),
    test.test(result([1, 2, 3, 4], 1), [4, 1, 2, 3]),
    test.test(result([1, 2, 3, 4], 0), [1, 2, 3, 4]),
    test.test(result([1, 2, 3], 2), [2, 3, 1]),
    test.test(result([1, 2, 3], 1), [3, 1, 2]),
    test.test(result([1, 2, 3], 0), [1, 2, 3]),
    test.test(result([1, 2], 1), [2, 1]),
    test.test(result([1], 0), [1]),
    test.test(result([], 0), [])
);
