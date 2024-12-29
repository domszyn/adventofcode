import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

let components = parseInput(input, l => l.split('/').map(Number));
let bridges = [];
for (let i = 0; i < components.length; i++) {
    components[i].sort((a, b) => a - b);
    if (components[i][0] == 0) {
        bridges.push([[...components[i]], components.filter((v, idx) => i != idx)]);
    }
}
let maxStrength = Number.MIN_SAFE_INTEGER;
let maxLength = Number.MIN_SAFE_INTEGER;
let maxStrengthLongest = Number.MIN_SAFE_INTEGER;
while (bridges.length > 0) {
    let newBridges = [];
    let [b, remaining] = bridges[0];

    const port = b.last();

    for (let i = 0; i < remaining.length; i++) {
        if (remaining[i].includes(port)) {
            const brick = [...remaining[i]];
            if (brick.first() != port) {
                brick.reverse()
            }
            let newRemaining = [...remaining];
            newRemaining.splice(i, 1);
            newBridges.push([[...b, ...brick], newRemaining]);
        }
    }

    if (newBridges.length == 0) {
        const strength = b.sum();
        if (strength > maxStrength) {
            maxStrength = strength;
        }
        if (b.length > maxLength) {
            maxLength = b.length;
            maxStrengthLongest = strength;
        }  else if (b.length == maxLength && strength > maxStrengthLongest) {
            maxStrengthLongest = strength;
        }
    }

    bridges = [
        ...newBridges,
        ...bridges.slice(1)
    ];
}

console.log("Part 1", maxStrength);
console.log("Part 2", maxStrengthLongest);