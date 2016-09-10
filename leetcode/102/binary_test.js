"use strict";
const test = require('../../test.js');

function TreeNode(val, left, right) { [this.val, this.left, this.right] = [val, left, right]; }

var levelOrder = function(root) {
    if (root === null) { return []; }

    let [result, vals, next] = [[], [], [root]];
    
    while (next.length !== 0) {
	[vals, next] = level(next);
	result.push(vals);
    }
    return result;
};

function level(nodes) {
    let [vals, next] = [[], []];
    
    for (let i = 0; i < nodes.length; i++) {
	vals.push(nodes[i].val);
	if (nodes[i].left !== null) { next.push(nodes[i].left); }
	if (nodes[i].right !== null) { next.push(nodes[i].right); }
    }

    return [vals, next];
}

test.Run(
    test.test(levelOrder(null), []),
    test.test(levelOrder(new TreeNode(1, null, null)), [[1]]),
    test.test(levelOrder(new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null))), [[1], [2, 3]]),
    test.test(level([new TreeNode(1, null, null)]), [[1], []]),
    test.test(level([new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null))]), [[1], [new TreeNode(2, null, null), new TreeNode(3, null, null)]]),
    test.test(level([new TreeNode(1, null, null), new TreeNode(2, null, null)]), [[1, 2], []])
);
