import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

const grid = parseInput(input, l => [...l]);

let antennas = new Map();

for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
        let elem = grid[i][j]
        if (elem !== '.') {
            if (!antennas.has(elem)) {
                antennas.set(elem, []);
            }

            antennas.set(elem, [...antennas.get(elem), { x: j, y: i }]);
        }
    }
}

let antinodes = new Set();
let antinodes2 = new Set();

for (const [a, nodes] of antennas) {
    for (let i = 0; i < nodes.length; i++) {
        for (let j = 0; j < nodes.length; j++) {
            if (i === j) continue;
            let { x: ax, y: ay } = nodes[i];
            let { x: bx, y: by } = nodes[j];

            let dx = bx - ax;
            let dy = by - ay;

            let cx = ax;
            let cy = ay;

            let part1added = false;

            while (cx >= 0 && cx < grid[0].length && cy >= 0 && cy < grid.length) {
                if (!part1added && cx != ax && cy != ay) {
                    antinodes.add(cx + ',' + cy);
                    part1added = true;
                }

                antinodes2.add(cx + ',' + cy);
                cx -= dx;
                cy -= dy;
            }
        }
    }

}

console.log("Part 1", antinodes.size);
console.log("Part 2", antinodes2.size);