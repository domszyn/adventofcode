import '../array.js';
import { parseInput } from '../utils.js';
import { input } from './input.js';

const ids = parseInput(input, l => l.split('   ').map(n => +n));
const lists = [
    ids.map(id => id[0]).sort(),
    ids.map(id => id[1]).sort(),
];

console.log("Part 1", lists[0].map((n, i) => Math.abs(n - lists[1][i])).sum());
console.log("Part 2", lists[0].map((n) => n * lists[1].filter(nn => nn == n).length).sum());