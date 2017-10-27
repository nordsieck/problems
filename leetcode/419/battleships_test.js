"use strict";
const test = require("../../test.js");

var countBattleships = function(board) {
    let count = 0;
    for (let i = 0; i < board.length; i++) {
	for (let j = 0; j < board[i].length; j++) {
	    if (detect(board, i, j)) { count++; }
	}
    }
    return count;
}

function detect(board, x, y) {
    if (board[x][y] != "X") { return false; }
    if (isX(board, x, y-1)) { return false; }
    if (isX(board, x-1, y)) { return false; }
    return true;
}

// allowed to call coords off the board in the negative direction
function isX(board, x, y) {
    if (x < 0) { return false; }
    if (y < 0) { return false; }
    if (board[x][y] === ".") { return false; }
    return true;
}

test.Run(
    test.test(countBattleships([[".", ".", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."]]), 0),
    test.test(countBattleships([["X", ".", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."]]), 1),
    test.test(countBattleships([["X", "X", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."], [".", ".", ".", "."]]), 1),
    test.test(countBattleships([["X", ".", ".", "X"], [".", ".", ".", "X"], [".", ".", ".", "X"]]), 2),
    test.test(countBattleships([["X", "."], [".", "X"]]), 2),
);

