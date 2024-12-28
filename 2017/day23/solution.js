import { input } from "./input.js";
import { parseInput } from "../utils.js";

const registers = new Map();
let part1 = 0, part2 = 0;
const operations = new Map([
    ['set', (x, y) => {
        registers.set(x, Number.isInteger(y) ? y : (registers.get(y) ?? 0));
        return 1;
    }],
    ['sub', (x, y) => {
        registers.set(x, (registers.get(x) ?? 0) - (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
        return 1;
    }],
    ['mul', (x, y) => {
        registers.set(x, (registers.get(x) ?? 0) * (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
        part1++;
        return 1;
    }],
    ['jnz', (x, y) => {
        const val = Number.isInteger(x) ? x : (registers.get(x) ?? 0);
        if (val != 0) {
            return Number.isInteger(y) ? y : (registers.get(y) ?? 0);
        } else {
            return 1;
        }
    }],
]);

const instructions = parseInput(input, l => {
    let [op, a, b] = l.split(' ');

    return {
        name: op,
        op: operations.get(op),
        a: Number.isNaN(+a) ? a : +a,
        b: Number.isNaN(+b) ? b : +b,
    };
});

for (let i = 0; i < instructions.length;) {
    const { name, op, a, b } = instructions[i];
    i += op(a, b);
}

for (let b = 108400; b <= 125400; b += 17) {
    for (let d = 2; d < b / 2; d++) {
        if (b % d == 0) {
            part2++;
            break;
        };
    }
}

console.log("Part 1", part1);
console.log("Part 2", part2);