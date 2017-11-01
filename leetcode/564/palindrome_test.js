"use strict";
const test = require("../../test.js");

// we have to do all our processing on strings because some of the test cases are
// larger than 32-bit ints.
var nearestPalindromic = function(n) {
    if (n === "0") { return "1"; }

    // 1: get the obvious palindrome
    let pal = n.slice(0, Math.floor((n.length + 1) / 2));
    pal += n.slice(0, Math.floor((n.length) / 2)).split("").reverse().join("");

    // 2: choose the better of the two candidates
    let u, d;
    
    switch(compare(pal, n)) {
    case 0:
	u = up(pal);
	d = down(pal);
	break;
    case -1:
	u = pal;
	d = down(pal);
	break;
    case 1:
	u = up(pal);
	d = pal;
    }
    
    if (parseInt(u) - parseInt(n) < parseInt(n) - parseInt(d)) { return u; }
    return d;
};

function update(num, newDigit, pos) { return num.slice(0, pos) + newDigit + num.slice(pos + 1); }

function compare(a, b) {
    if (a === b) { return 0; }
    if (a.length > b.length) { return -1; }
    if (b.length > a.length) { return 1; }
    for (let i = 0; i < a.length; i++) {
	if (a[i] > b[i]) { return -1; }
	if (b[i] > a[i]) { return 1; }
    }
}

function down(num) {
    if (num === "1") { return "0"; }
    let top = (num.length - 1) / 2 | 0;
    let bot = num.length / 2 | 0;

    for (let i = 0; i <= top; i++) {
	let val = num[top - i];

	num = update(num, dec(parseInt(val)).toString(), top - i);
	num = update(num, dec(parseInt(val)).toString(), bot + i);

	if (i === top && num[num.length - 1] === "0") { num = update(num, "9", num.length - 1).slice(1); }
	if (val != "0") { break; }
    }
    return num;
}

function up(num) {
    if (nines(num)) { return "1" + "0".repeat(num.length - 1) + "1"; }
    let top = (num.length - 1) / 2 | 0;
    let bot = num.length / 2 | 0;

    for (let i = 0; i <= top; i++) {
	let val = num[top - i];

	num = update(num, inc(parseInt(val)).toString(), top - i);
	num = update(num, inc(parseInt(val)).toString(), bot + i);

	if (val != "9") { break; }
    }
    return num;
}
    
function nines(num) {
    for (let c of num) { if (c != "9") { return false; }}
    return true;
}

function dec(num) { return (num + 9) % 10; }
function inc(num) { return (num + 1) % 10; }

function w(val, label = "") { process.stdout.write(label + val.toString() + "\n"); }

test.Run(
    test.test(nines("9"), true),
    test.test(nines("999999"), true),
    test.test(nines("99199"), false),
    
    test.test(update("0", "1", 0), "1"),
    test.test(update("123", "4", 0), "423"),
    test.test(update("123", "4", 1), "143"),
    test.test(update("123", "4", 2), "124"),
    
    test.test(compare("0", "1"), 1),
    test.test(compare("0", "0"), 0),
    test.test(compare("100", "10"), -1),
    
    test.test(nearestPalindromic("0"), "1"),
    test.test(nearestPalindromic("1"), "0"),
    test.test(nearestPalindromic("2"), "1"),
    test.test(nearestPalindromic("9"), "8"),
    test.test(nearestPalindromic("10"), "9"),
    test.test(nearestPalindromic("11"), "9"),
    test.test(nearestPalindromic("12"), "11"),
    test.test(nearestPalindromic("22"), "11"),
    test.test(nearestPalindromic("99"), "101"),
    test.test(nearestPalindromic("100"), "99"),
    test.test(nearestPalindromic("101"), "99"),
    test.test(nearestPalindromic("111"), "101"),
    test.test(nearestPalindromic("123"), "121"),
    test.test(nearestPalindromic("999"), "1001"),
    test.test(nearestPalindromic("10001"), "9999"),
    test.test(nearestPalindromic("11011"), "11111"),
    test.test(nearestPalindromic("20002"), "19991"),
    test.test(nearestPalindromic("21012"), "21112"),
    test.test(nearestPalindromic("22022"), "22122"),
    
    test.test(dec(4), 3),
    test.test(dec(0), 9),

    test.test(inc(4), 5),
    test.test(inc(9), 0),
    
    test.test(down("1"), "0"),
    test.test(down("2"), "1"),
    test.test(down("11"), "9"),
    test.test(down("22"), "11"),
    test.test(down("101"), "99"),
    test.test(down("222"), "212"),
    test.test(down("10001"), "9999"),
    test.test(down("20002"), "19991"),

    test.test(up("1"), "2"),
    test.test(up("9"), "11"),
    test.test(up("11"), "22"),
    test.test(up("111"), "121"),
    test.test(up("9999"), "10001"),
);
