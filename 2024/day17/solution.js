import { input, example } from "./input.js";
import '../array.js';

let [a, b, c, , program] = input.split('\n');
[, a] = a.split(': ');
a = +a;
[, b] = b.split(': ');
b = +b;
[, c] = c.split(': ');
c = +c;
[, program] = program.split(': ');
program = program.split(',').map(Number);

const cache = new Map();
const execute = (a, b, c) => {
    if (cache.has(a)) {
        return cache.get(a);
    }
    const startA = a;
    const getCombo = operand => {
        switch (operand) {
            case 0n:
            case 1n:
            case 2n:
            case 3n:
                return operand;
            case 4n:
                return a;
            case 5n:
                return b;
            case 6n:
                return c;
            case 7n:
                throw new Error('invalid state');
        }
    }

    let pos = 0;
    let output = [];
    while (true) {
        if (pos >= program.length) break;
        let opcode = program[pos];
        let operand = BigInt(program[pos + 1]);

        switch (opcode) {
            case 0:
                a = a >> getCombo(operand);
                pos += 2;
                break;
            case 1:
                b ^= operand;
                pos += 2;
                break;
            case 2:
                b = getCombo(operand) % 8n;
                pos += 2;
                break;
            case 3:
                if (a != 0) {
                    pos = Number(operand);
                } else {
                    pos += 2;
                }
                break;
            case 4:
                const bb = b ^ c;
                b = bb;
                pos += 2;
                break;
            case 5:
                let o = getCombo(operand) % 8n;
                output.push(o.toString());
                pos += 2;
                break;
            case 6:
                b = a >> getCombo(operand);
                pos += 2;
                break;
            case 7:
                c = a >> getCombo(operand);
                pos += 2;
                break;
        }
    }

    cache.set(a, output);
    return output;
}

console.log("Part 1: ", execute(BigInt(a), 0n, 0n).join());

const tailMatch = (program, output, from) => program.slice(from).eq(output.map(Number).slice(from));

let queue = [Math.pow(2, 45).toString(8).split('').map(Number)];
for (let i = 0; i < 16; i++) {
    let nextQueue = [];
    for (const q of queue) {
        for (let j = 0; j < 8; j++) {
            q[i] = j;
            nextQueue.push([...q]);
        }
    }
    queue = [];

    for (const q of nextQueue) {
        const a = BigInt(parseInt(q.join(''), 8));
        const output = execute(a, 0, 0);
        if (tailMatch(program, output, 15 - i)) {
            queue.push([...q]);
        }
    }
}

console.log(queue.map(q => parseInt(q.join(''), 8)).min());