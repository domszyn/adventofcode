import { input } from "./input.js";
import { parseInput } from "../utils.js";

const program = async (id, inQueue, outQueue) => {
    const registers = new Map();
    registers.set('p', id);
    let sent = 0;
    let received = 0;
    const operations = new Map([
        ['snd', async (x) => {
            const val = Number.isInteger(x) ? x : (registers.get(x) ?? 0);
            sent++;
            // console.log(`Program ${id} sent ${val}, ${sent} values in total`);
            await new Promise(resolve => setTimeout(resolve, 2))
            outQueue.push(val);
            return 1
        }],
        ['set', async (x, y) => {
            registers.set(x, Number.isInteger(y) ? y : (registers.get(y) ?? 0));
            return 1;
        }],
        ['add', async (x, y) => {
            registers.set(x, (registers.get(x) ?? 0) + (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
            return 1;
        }],
        ['mul', async (x, y) => {
            registers.set(x, (registers.get(x) ?? 0) * (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
            return 1;
        }],
        ['mod', async (x, y) => {
            registers.set(x, (registers.get(x) ?? 0) % (Number.isInteger(y) ? y : (registers.get(y) ?? 0)));
            return 1;
        }],
        ['rcv', async (x) => {
            let totalTimeout = 0;
            while (inQueue.length == 0) {
                await new Promise(resolve => setTimeout(resolve, 1));
                totalTimeout += 5;
                if (totalTimeout > 5000) {
                    if (id == '1') {
                        console.log("Part 2", sent);
                    }
                    return input.length;
                }
            }
            const val = inQueue.shift();
            received++;
            // console.log(`Program ${id} received ${val}, ${received} values in total`);
            registers.set(x, val);
            return 1;
        }],
        ['jgz', async (x, y) => {
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

    return async () => {
        for (let i = 0; i < instructions.length;) {
            const { op, a, b } = instructions[i];
            i += await op(a, b);
        }
    }
}

const queue1 = [];
const queue2 = [];
const p1 = await program(0, queue1, queue2);
const p2 = await program(1, queue2, queue1);
await Promise.all([p1(), p2()]);