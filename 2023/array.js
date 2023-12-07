Array.prototype.sum = function() {
    return this.reduce((a, b) => a + b, 0);
}

Array.prototype.multiply = function() {
    return this.reduce((a, b) => a * b, 1);
}

Array.prototype.min = function() {
    return this.reduce((a, b) => a < b ? a : b, Number.MAX_SAFE_INTEGER);
}

Array.prototype.max = function() {
    return this.reduce((a, b) => a > b ? a : b, Number.MIN_SAFE_INTEGER);
}

Array.prototype.flatten = function() {
    return this.reduce((a, b) => ([...a, ...b]), []);
}

Array.prototype.eq = function (arr) {
    return this.length == arr.length && arr.every((elem, idx) => elem == this[idx]);
};