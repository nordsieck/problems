"use strict";
const test = require("../../test.js");

function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

var constructMaximumBinaryTree = function(nums) {
    if (nums.length == 0) { return null; }
    let idx = 0;
    for (let i = 1; i < nums.length; i++) {
	if (nums[i] > nums[idx]) { idx = i; }
    }
    let tn = new TreeNode(nums[idx]);
    tn.left = constructMaximumBinaryTree(nums.slice(0, idx));
    tn.right = constructMaximumBinaryTree(nums.slice(idx+1, nums.length));
    return tn;
};

function n(val, left = null, right = null) {
    let tn = new TreeNode(val);
    tn.left = left;
    tn.right = right;
    return tn;
}

test.Run(
    test.test(constructMaximumBinaryTree([1]), n(1)),
    test.test(constructMaximumBinaryTree([1, 3, 2]), n(3, n(1), n(2))),
);
