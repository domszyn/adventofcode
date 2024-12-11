import { parseInput } from "../utils.js";
import { input } from "./input.js";
import '../array.js';

let stones = input.split(' ').map(Number);

const archive = new Map();
const readArchive = (stone, gen) => archive.get(stone + '@' + gen);
const writeArchive = (stone, gen, count) => archive.set(stone + '@' + gen, count);

const blink = (stone, gen, end) => {
    if (gen === end - 1) {
        if (stone === 0) {
            writeArchive(stone, gen, 1);
            return 1;
        } else if (String(stone).length % 2 === 0) {
            writeArchive(stone, gen, 2);
            return 2;
        } else {
            writeArchive(stone, gen, 1);
            return 1;
        }
    }

    if (stone === 0) {
        const count = readArchive(1, gen + 1) ?? blink(1, gen + 1, end);
        writeArchive(stone, gen, count);
        return count;
    } else {
        const s = String(stone);
        let count = readArchive(stone, gen);
        if (count !== undefined) {
            return count;
        }
        if (s.length % 2 === 0) {
            const left = Number(s.slice(0, s.length / 2));
            const right = Number(s.slice(s.length / 2));
            const blinkLeft = readArchive(left, gen + 1) ?? blink(left, gen + 1, end);
            const blinkRight = readArchive(right, gen + 1) ?? blink(right, gen + 1, end);
            writeArchive(stone, gen, blinkLeft + blinkRight);
            return blinkLeft + blinkRight
        } else {
            const ns = stone * 2024;
            const count = readArchive(ns, gen + 1) ?? blink(ns, gen + 1, end);
            writeArchive(stone, gen, count)
            return count;
        }
    }
}

for (let i = 74; i >= 0; i--) {
    for (let j = 0; j < stones.length; j++) {
        blink(stones[j], i, 75);
    }
}

console.log("Part 1", stones.map(s => readArchive(s, 50)).sum());
console.log("Part 2", stones.map(s => readArchive(s, 0)).sum());