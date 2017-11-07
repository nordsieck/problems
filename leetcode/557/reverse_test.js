"use strict";
const test = require("../../test.js");

var reverseWords = function(s) {
    return s.split(" ").map((word) => word.split("").reverse().join("")).join(" ");
}

test.Run(
    test.test(reverseWords(""), ""),
    test.test(reverseWords("abc"), "cba"),
    test.test(reverseWords("foo bar"), "oof rab"),
);
