"use strict";
const test = require("../../test.js");

/**
 * @param {number} capacity
 */
var LRUCache = function(capacity) {
    this.capacity = capacity;
    this.cache = {};
    this.list = Object.create(List);
};

/** 
 * @param {number} key
 * @return {number}
 */
LRUCache.prototype.get = function(key) {
    if (this.cache[key] === undefined) { return -1; }
    this.list.promote(this.cache[key].node);
    return this.cache[key].val;
};

/** 
 * @param {number} key 
 * @param {number} value
 * @return {void}
 */
LRUCache.prototype.put = function(key, value) {
    if (this.cache[key] === undefined) {
	if (Object.keys(this.cache).length === this.capacity) {
	    let n = this.list.trim();
	    delete this.cache[n.key];
	}
	
	this.list.add(new node(key));
	this.cache[key] = new Val(value, this.list.head);
    } else {
	this.cache[key].val = value;
	this.list.promote(this.cache[key].node);
    }
};

function Val(val, node) {
    this.val = val;
    this.node = node;
}

function node(key, head = null, tail = null) {
    this.key = key;
    this.head = head;
    this.tail = tail;
}

let List = {
    head: null,
    tail: null,
    add: function(n) {
	if (this.tail === null) { this.tail = n; }
	if (this.head === null) { this.head = n; }
	else {
	    n.tail = this.head;
	    n.tail.head = n;
	    this.head = n;
	}
    },
    promote: function(n) {
	if (n.head != null) { n.head.tail = n.tail; }
	if (n.tail != null) { n.tail.head = n.head; }
	if (this.head === n) { this.head = n.tail; }
	if (this.tail === n) { this.tail = n.head; }
	n.head = n.tail = null;
	this.add(n);
    },
    trim: function() {
	if (this.tail === null) { return null; }
	if (this.tail.head === null) {
	    let n = this.tail;
	    this.tail = this.head = null;
	    return n;
	}
	let n = this.tail;
	this.tail = this.tail.head;
	this.tail.tail = null;
	return n;
    }
}

let l = Object.create(List);
let zero = new node(0);
let one = new node(1);

let cache = new LRUCache(2);

test.Run(
    
    // empty
    test.test(cache.get("foo"), -1),

    // foo
    test.test(cache.put("foo", "f"), undefined),
    test.test(cache.get("foo"), "f"),

    // foo, bar
    test.test(cache.put("bar", "b"), undefined),
    test.test(cache.get("foo"), "f"),
    test.test(cache.get("bar"), "b"),

    // bar, baz
    test.test(cache.put("baz", "z"), undefined),
    test.test(cache.get("foo"), -1),
    test.test(cache.get("bar"), "b"),
    test.test(cache.get("baz"), "z"),
    
    
    // empty
    test.test(l.trim(), null),

    // one
    test.test(l.add(zero), undefined),
    test.test(l.tail.key, 0),
    test.test(l.head.key, 0),
    test.test(l.trim(), zero),
    test.test(l.head, null),
    test.test(l.tail, null),
    test.test(l.add(zero), undefined),

    // two
    test.test(l.add(one), undefined),
    test.test(l.tail.key, 0),
    test.test(l.head.key, 1),
    test.test(l.head.tail.key, 0),
    test.test(l.promote(zero), undefined),
    test.test(l.tail.key, 1),
    test.test(l.head.key, 0),
    test.test(l.trim(), one),
);
