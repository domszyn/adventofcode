import { designs, patterns } from "./input.js";
import '../array.js';

const impossibleDesigns = new Set();
const counts = new Map();

const possible = design => {
    if (impossibleDesigns.has(design)) {
        return 0;
    } else  if (design.length == 0) {
        return 1;
    } else if (counts.has(design)) {
        return counts.get(design);
    }

    const remainingDesigns = patterns
        .filter(p => design.startsWith(p))
        .map(p => design.slice(p.length))
        .filter(d => !impossibleDesigns.has(d))
        .map(possible);

    if (remainingDesigns.length == 0) {
        impossibleDesigns.add(design);
        return 0;
    } else {
        const count = remainingDesigns.sum()
        counts.set(design, count);
        return count;
    }
}

console.log("Part 1", designs.filter(possible).length);
console.log("Part 2", designs.map(possible).sum());