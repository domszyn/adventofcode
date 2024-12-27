import { input } from "./input.js";
import { parseInput } from "../utils.js";

const registers = new Map();
const operations = new Map([
    ['snd', (x) => {
        registers.set('snd', Number.isInteger(x) ? x : (registers.get(x) ?? 0));
        return 1
    }],
    ['set', (x, y) => {
        registers.set(x, Number.isInteger(y) ? y : (registers.get(y) ?? 0));
        return 1;
    }],
    ['add', (x, y) => {
        registers.set(x, (registers.get(x) ?? 0) + (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
        return 1;
    }],
    ['mul', (x, y) => {
        registers.set(x, (registers.get(x) ?? 0) * (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
        return 1;
    }],
    ['mod', (x, y) => {
        registers.set(x, (registers.get(x) ?? 0) % (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
        return 1;
    }],
    ['rcv', (x) => {
        const val = Number.isInteger(x) ? x : (registers.get(x) ?? 0);
        if (val > 0) {
            console.log(registers.get('snd'));
            return input.length;
        } else {
            return 1;
        }
    }],
    ['jgz', (x, y) => {
        const val = Number.isInteger(x) ? x : (registers.get(x) ?? 0);
        if (val > 0) {
            return Number.isInteger(y) ? y : (registers.get(y) ?? 0);
        } else {
            return 1;
        }
    }],
]);

const instructions = parseInput(input, l => {
    let [op, a, b] = l.split(' ');

    return {
        op: operations.get(op),
        a: Number.isNaN(+a) ? a : +a,
        b: Number.isNaN(+b) ? b : +b,
    };
});

for (let i = 0; i < instructions.length;) {
    const { op, a, b } = instructions[i];
    i += op(a, b);
}