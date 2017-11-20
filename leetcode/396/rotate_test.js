"use strict"
const test = require("../../test.js")

var maxRotateFunction = function(A) {
    if (A.length === 0) { return 0; }

    let max = Number.MIN_SAFE_INTEGER;
    for (let i = 0; i < A.length; i++) {
        let a = A.slice(i).concat(A.slice(0, i));
        let current = F(a);
        if (current > max) { max = current; }
    }
    return max;
}

function F(A) {
    let sum = 0;
    for (let i = 1; i < A.length; i++) { sum += i * A[i]; }
    return sum;
}

test.Run(
    test.test(maxRotateFunction([]), 0),
    test.test(maxRotateFunction([4, 3, 2, 6]), 26),
    test.test(maxRotateFunction([-2147483648,-2147483648]), -2147483648),
);
