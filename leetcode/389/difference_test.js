"use strict";
const test = require('../../test.js');

var findTheDifference = function(s, t) {
    let [freqs, freqt] = [{}, {}];
    
    for (let c of s) { freqs[c] = freqs[c] === undefined ? 1 : freqs[c] + 1; }
    for (let c of t) { freqt[c] = freqt[c] === undefined ? 1 : freqt[c] + 1; }

    for (let char in freqt) {
	if (freqs[char] === undefined)   { return char; }
	if (freqs[char] !== freqt[char]) { return char; }
    }
};

test.Run(
    test.test(findTheDifference("aa", "aaa"), "a"),
    test.test(findTheDifference("a", "ab"), "b"),
    test.test(findTheDifference("a", "ba"), "b"),
    test.test(findTheDifference("", "t"), "t")
);
