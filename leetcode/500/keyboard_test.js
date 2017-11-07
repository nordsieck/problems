"use strict";
const test = require("../../test.js");

var findWords = function(words) {
    let ret = [];
    let letters = [
	new Set(["q", "w", "e", "r", "t", "y", "u", "i", "o", "p"]),
	new Set(["a", "s", "d", "f", "g", "h", "j", "k", "l"]),
	new Set(["z", "x", "c", "v", "b", "n", "m"]),
    ];

    for (let w of words) {
	for (let l of letters) {
	    if (!w.toLowerCase().split("").reduce((reject, val) => {
		if (l.has(val)) { return false || reject; }
		return true || reject;
	    }, false)) {
		ret.push(w);
		break;
	    }
	}
    }
    return ret;
};

test.Run(
    test.test(findWords(["q"]), ["q"]),
    test.test(findWords(["qwerty"]), ["qwerty"]),
    test.test(findWords(["q", "a", "z"]), ["q", "a", "z"]),
    test.test(findWords(["qaz"]), []),
);
