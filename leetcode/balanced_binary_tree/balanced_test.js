"use strict"

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

function test(fn, input, expected) {
    let [out, exp] = [JSON.stringify(fn(input)), JSON.stringify(expected)];

    if (out != exp) { process.stdout.write("expected: " + exp + "\ngot:     " + out + "\n"); return false; }
    return true;
}

function Run() {
    let failure = false;
    for (let i in arguments) { if (!arguments[i]) { failure = true; }}
    if (failure) { process.exit(1); }
}

Run(
    test(isBalanced, new TreeNode(0, new TreeNode(0, new TreeNode(0, null, null), null), null), false),
    test(isBalanced, new TreeNode(0, null, null), true),
    test(isBalanced, null, true),
    test(heightAndBalance, new TreeNode(0, new TreeNode(0, new TreeNode(0, null, null), null), null), [3, false]),
    test(heightAndBalance, new TreeNode(0, null, null), [1, true]),
    test(heightAndBalance, null, [0, true])
);
