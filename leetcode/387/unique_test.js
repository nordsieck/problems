"use strict";
const test = require('../../test.js');

var firstUniqChar = function(s) {
    let freqs = {};

    for (let ch of s)                  { freqs[ch] = freqs[ch] === undefined ? 1 : freqs[ch] + 1; }
    for (let i = 0; i < s.length; i++) { if (freqs[s[i]] === 1) { return i; }}
    return -1;
};

test.Run(
    test.test(firstUniqChar("aba"), 1),
    test.test(firstUniqChar("aaaa"), -1),
    test.test(firstUniqChar("aa"), -1),
    test.test(firstUniqChar("a"), 0),
    test.test(firstUniqChar(""), -1)
);
