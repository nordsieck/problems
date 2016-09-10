"use strict"
const test = require('../../test.js');

function TreeNode(val, left, right) { [this.val, this.left, this.right] = [val, left, right]; }

var isBalanced = function(root) {
    let [height, balance] = heightAndBalance(root);
    return balance;
}

function heightAndBalance(node) {
    if (node === null) { return [0, true]; }
    let [[lh, lb], [rh, rb]] = [heightAndBalance(node.left), heightAndBalance(node.right)];
    return [Math.max(lh, rh) + 1, lb && rb && Math.abs(lh - rh) < 2];
}

test.Run(
    test.test(isBalanced(new TreeNode(0, new TreeNode(0, new TreeNode(0, null, null), null), null)), false),
    test.test(isBalanced(new TreeNode(0, null, null)), true),
    test.test(isBalanced(null), true),
    test.test(heightAndBalance(new TreeNode(0, new TreeNode(0, new TreeNode(0, null, null), null), null)), [3, false]),
    test.test(heightAndBalance(new TreeNode(0, null, null)), [1, true]),
    test.test(heightAndBalance(null), [0, true])
);
