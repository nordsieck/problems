"use strict";
const test = require('../../test.js');

var countPrimes = function(n) {
    let [field, primes] = [[], []];
    for (let i = 2; i < n; i++) {
	if (field[i] === 0) { continue; }
	field[i] = 1;
	for (let j = 2 * i; j < n; j += i) { field[j] = 0; }
    }

    for (let i = 2; i < n; i++) { if (field[i] === 1) { primes.push(field[i]); }}
    return primes.length;
};

test.Run(
    test.test(countPrimes(7920), 1000),
    test.test(countPrimes(6), 3),
    test.test(countPrimes(4), 2),
    test.test(countPrimes(3), 1),
    test.test(countPrimes(2), 0),
    test.test(countPrimes(1), 0),
    test.test(countPrimes(0), 0)
);
