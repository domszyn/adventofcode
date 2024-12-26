import { input } from './input.js';
import { parseInput, primeFactors } from '../utils.js';
import '../array.js';

let maxDepth = 0;
const firewalls = parseInput(input, s => {
    var [depth, range] = s.split(": ").map(Number);
    if (depth > maxDepth) {
        maxDepth = depth;
    }
    return { depth, range };
});

console.log("Part 1", firewalls
    .filter(({ depth, range }) => (depth) % (2 * range - 2) == 0)
    .map(({ depth, range }) => depth * range)
    .sum()
);

for (let delay = 0; ; delay++) {
    let caught = false;
    for (const { depth, range } of firewalls) {
        if ((delay + depth) % (2 * range - 2) == 0) {
            // console.log(`Caught with delay ${delay} at depth ${depth}`);
            caught = true;
            break;
        }
    }
    if (caught) continue;
    console.log("Part 2", delay);
    break;
}

