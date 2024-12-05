import { input } from "./input.js";
import '../array.js';
import { parseInput } from "../utils.js";

let [rules, updates] = input.split(`\n\n`);
rules = rules.split(`\n`);
updates = parseInput(updates, l => l.split(',').map(n => +n));

const isOrdered = update => {
    let ordered = true;
    for (let i = 0; i < update.length - 1; i++) {
        for (let j = i + 1; j < update.length; j++) {
            const rule = update[i] + '|' + update[j];
            if (rules.indexOf(rule) === -1) {
                update.swap(i, j);
                ordered = false;
            }
        }
    }

    return ordered;
}

const middlePage = update => update[(update.length - 1) / 2];

let part1 = 0;
let part2 = 0;
for (const update of updates) {
    if (isOrdered(update)) {
        part1 += middlePage(update);
    } else {
        part2 += middlePage(update);
    }
}

console.log("Part 1", part1);
console.log("Part 2", part2);