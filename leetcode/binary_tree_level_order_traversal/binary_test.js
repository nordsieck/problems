"use strict";

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

function test(fn, nodes, expected) {
    let [out, exp] = [JSON.stringify(fn(nodes)), JSON.stringify(expected)];

    if (out !== exp) { process.stdout.write("expected: " + exp + "\ngot:      " + out + "\n"); return false; }
    return true;
}

function Test() {
    let failure = false;
    for (let i in arguments) { if (!arguments[i]) { failure = true; }}
    if (failure) { process.exit(1); }
}

Test(
    test(levelOrder, null, []),
    test(levelOrder, new TreeNode(1, null, null), [[1]]),
    test(levelOrder, new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null)), [[1], [2, 3]]),
    test(level, [new TreeNode(1, null, null)], [[1], []]),
    test(level, [new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null))], [[1], [new TreeNode(2, null, null), new TreeNode(3, null, null)]]),
    test(level, [new TreeNode(1, null, null), new TreeNode(2, null, null)], [[1, 2], []])
);
