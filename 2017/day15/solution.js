import { a, b } from './input.js';

class Generator {
    constructor(start, factor, multipleOf) {
        this.value = start;
        this.factor = factor;
        this.multipleOf = multipleOf ?? 1;
    }

    next() {
        this.value = this.value * this.factor % 2147483647;
        if (this.value % this.multipleOf == 0) {
            return this.value << 16;
        } else {
            return this.next();
        }
    }
}

let genA = new Generator(a, 16807), genB = new Generator(b, 48271);

let part1 = 0;
for (let i = 0; i < 40000000; i++) {
    if (genA.next() == genB.next()) {
        part1++;
    }
}
console.log(part1);

genA = new Generator(a, 16807, 4);
genB = new Generator(b, 48271, 8);

let part2 = 0;
for (let i = 0; i < 5000000; i++) {
    if (genA.next() ==  genB.next()) {
        part2++;
    }
}
console.log(part2);