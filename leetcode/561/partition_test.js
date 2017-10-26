"use strict";
const test = require("../../test.js");

var arrayPairSum = function(nums) {
    var nums = nums.sort((a,b) => a - b );
    let sum = 0;
    for (let i = 0; i < nums.length; i += 2) {
	sum += nums[i];
    }

    return sum;
};

test.Run(
    test.test(arrayPairSum([1, 4, 3, 2]), 4),
    test.test(arrayPairSum([1, 2, 3, 2]), 3),
);
