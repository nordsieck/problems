"use strict";
const test = require(`../../test.js`);

var mergeTrees = function(t1, t2) {
    if (t1 === null) { return t2; }
    if (t2 === null) { return t1; }
    let tn = new TreeNode(t1.val + t2.val);
    tn.left = mergeTrees(t1.left, t2.left);
    tn.right = mergeTrees(t1.right, t2.right);
    return tn;
};

function TreeNode(val) {
    this.val = val;
    this.left = this.right = null;
}

function n(val, left = null, right = null) {
    let tn = new TreeNode(val)
    tn.left = left;
    tn.right = right;
    return tn;
}

test.Run(
    test.test(mergeTrees(null, null), null),
    test.test(mergeTrees(null, n(1)), n(1)),
    test.test(mergeTrees(n(1), n(2)), n(3)),
    test.test(mergeTrees(n(1, n(3, n(5), null), n(2)), n(2, n(1, null, n(4)), n(3, null, n(7)))), n(3, n(4, n(5), n(4)), n(5, null, n(7)))),
);
