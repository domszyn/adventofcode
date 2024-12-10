import { parseInput } from "../utils.js";
import { input } from "./input.js";
import '../array.js';

const grid = parseInput(input, l => l.split('').map(Number));

let part1 = 0, part2 = 0;
for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[0].length; x++) {
        if (grid[y][x] === 0) {
            const paths = grid.walk({ x, y }, 9, (v1, v2) => v2 - v1 === 1);
            part1 += new Set(paths.map(p => {
                const { x, y } = p.last();
                return x + ',' + y;
            })).size;
            part2 += paths.length
        }
    }
}

console.log("Part 1", part1);
console.log("Part 2", part2);
