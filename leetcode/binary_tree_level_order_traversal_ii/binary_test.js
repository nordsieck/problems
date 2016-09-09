function TreeNode(val, left, right) { [this.val, this.left, this.right] = [val, left, right]; }

var levelOrderBottom = function(root) {
    var [result, vals, next] = [[], [], [root]];
    
    while (next.length !== 0) {
	[vals, next] = level(next);
	result.unshift(vals);
    }
    return result;
};

function level(nodes) {
    var [vals, next] = [[], []];
    
    for (var i = 0; i < nodes.length; i++) {
	vals.push(nodes[i].val);
	if (nodes[i].left !== null) { next.push(nodes[i].left); }
	if (nodes[i].right !== null) { next.push(nodes[i].right); }
    }

    return [vals, next];
}

function test(fn, nodes, expected) {
    var [out, exp] = [JSON.stringify(fn(nodes)), JSON.stringify(expected)];

    if (out !== exp) { process.stdout.write("expected: " + exp + "\ngot:      " + out + "\n"); return false; }
    return true;
}

function Test() {
    var failure = false;
    for (i in arguments) { if (!arguments[i]) { failure = true; }}
    if (failure) { process.exit(1); }
}

Test(
    test(levelOrderBottom, new TreeNode(1, null, null), [[1]]),
    test(levelOrderBottom, new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null)), [[2, 3], [1]]),
    test(level, [new TreeNode(1, null, null)], [[1], []]),
    test(level, [new TreeNode(1, new TreeNode(2, null, null), new TreeNode(3, null, null))], [[1], [new TreeNode(2, null, null), new TreeNode(3, null, null)]]),
    test(level, [new TreeNode(1, null, null), new TreeNode(2, null, null)], [[1, 2], []])
);
