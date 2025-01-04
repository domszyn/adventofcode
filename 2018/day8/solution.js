import { input } from "./input.js";
import '../array.js';

const numbers = input.split(' ').map(Number);
const readTree = (nodes) => {
    const [numChildren, numMetadata] = nodes;
    nodes = nodes.slice(2);
    const children = new Array(numChildren);
    for (let i = 0; i < numChildren; i++) {
        const c = readTree(nodes);
        children[i] = { children: c.children, part1: c.part1, part2: c.part2 };
        nodes = c.rest;
    }
    
    const metadata = nodes.slice(0, numMetadata);
    let part2 = 0;
    if (numChildren == 0) {
        part2 = metadata.sum();
    } else {
        part2 = metadata
            .map(m => m - 1)
            .filter(m => m >= 0 && m < children.length)
            .map(m => children[m].part2)
            .sum();
    }

    return {
        children,
        part1: children.map(c => c.part1).sum() + metadata.sum(),
        part2,
        rest: nodes.slice(numMetadata)
    };
}
const root = readTree(numbers);
console.log(root.part1);
console.log(root.part2);