import { designs, patterns } from "./input.js";
import '../array.js';

const counts = new Map();
counts.set('', 1);

const possible = design => {
    if (counts.has(design)) {
        return counts.get(design);
    }
    
    const count = patterns
        .filter(p => design.startsWith(p))
        .map(p => design.slice(p.length))
        .map(possible)
        .sum();

    counts.set(design, count);
    return count;
}

console.log("Part 1", designs.filter(possible).length);
console.log("Part 2", designs.map(possible).sum());