"use strict";
const test = require('../../test.js');

var judgeCircle = function(moves) {
    let u = 0;
    let l = 0;
    for (let i = 0; i < moves.length; i++) {
	if (moves[i] === "U") { u++; }
	else if (moves[i] === "D") { u--; }
	else if (moves[i] === "L") { l++; }
	else { l--; }
    }
    return u === 0 && l === 0;
};

test.Run(
    test.test(judgeCircle(""), true),
    test.test(judgeCircle("D"), false),
    test.test(judgeCircle("UD"), true),
    test.test(judgeCircle("LR"), true),
    test.test(judgeCircle("LL"), false),
);
