"use strict";
const test = require("../../test.js");

var encode = function(longUrl) { return longUrl; };
var decode = function(shortUrl) { return shortUrl; };

test.Run(
    test.test(decode(encode("https://google.com")), "https://google.com"),
);
