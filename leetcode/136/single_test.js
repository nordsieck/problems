"use strict"
const test = require('../../test.js');

var singleNumber = function(nums) {
    let dict = {};
    for (let i in nums) {
	if (nums[i] in dict) { delete dict[nums[i]]; }
	else { dict[nums[i]] = true; }
    }
    for (let d in dict) { return parseInt(d); }
};

test.Run(
    test.test(singleNumber([0, 1, 2, 0, 2]), 1),
    test.test(singleNumber([1, 0, 0]), 1),
    test.test(singleNumber([0, 1, 0]), 1),
    test.test(singleNumber([0, 0, 1]), 1),
    test.test(singleNumber([1]), 1)
);
