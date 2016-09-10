"use strict";
const test = require('../../test.js');

var longestCommonPrefix = function(strs) {
    if (strs === null || strs.length === 0) { return ""; }
    for (let i = 0; i < strs[0].length; i++) {
	for (let j = 1; j < strs.length; j++) {
	    if (strs[j][i] !== strs[0][i]) { return strs[0].slice(0, i); }
	}
    }
    return strs[0];
};

test.Run(
    test.test(longestCommonPrefix(["abc", "abc", "abd"]), "ab"),
    test.test(longestCommonPrefix(["a", "a", "b"]), ""),
    test.test(longestCommonPrefix(["a", "a", "a"]), "a"),
    test.test(longestCommonPrefix(["aa", "ba"]), ""),
    test.test(longestCommonPrefix(["aa", "ab"]), "a"),
    test.test(longestCommonPrefix(["aa", "aa"]), "aa"),
    test.test(longestCommonPrefix(["a", "a"]), "a"),
    test.test(longestCommonPrefix(["a", "b"]), ""),
    test.test(longestCommonPrefix([]), ""),
    test.test(longestCommonPrefix(null), "")
);
