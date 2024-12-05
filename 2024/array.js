Array.prototype.sum = function () {
    return this.reduce((a, b) => a + b, 0);
}

Array.prototype.multiply = function () {
    return this.reduce((a, b) => a * b, 1);
}

Array.prototype.min = function () {
    return this.reduce((a, b) => a < b ? a : b, Number.MAX_SAFE_INTEGER);
}

Array.prototype.max = function () {
    return this.reduce((a, b) => a > b ? a : b, Number.MIN_SAFE_INTEGER);
}

Array.prototype.flatten = function () {
    return this.reduce((a, b) => ([...a, ...b]), []);
}

Array.prototype.eq = function (arr) {
    return this.length == arr.length && arr.every((elem, idx) => elem == this[idx]);
};

Array.prototype.swap = function(i, j) {
    if (i >= 0 && i < this.length && j >= 0 && j < this.length) {
        var tmp = this[i];
        this[i] = this[j];
        this[j] = tmp;
    }
}

export const gcd = (a, b) => {
    while (b != 0) {
        let t = b;
        b = a % b;
        a = t;
    }
    return a;
}

Array.prototype.lcm = function () {
    return this.reduce((a, b) => b * a / gcd(a, b), this[0]);
}

Array.prototype.first = function () {
    return this.length > 0 ? this[0] : undefined;
}

Array.prototype.last = function () {
    return this.length > 0 ? this[this.length - 1] : undefined;
}
