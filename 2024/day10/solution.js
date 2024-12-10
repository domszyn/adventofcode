import { parseInput } from "../utils.js";
import { input } from "./input.js";
import '../array.js';

const grid = parseInput(input, l => l.split('').map(Number));

let trailheads = [];
for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[0].length; x++) {
        if (grid[y][x] === 0) {
            trailheads.push({ x, y });
        }
    }
}

const getHeight = ({ x, y }) => {
    if (y < 0 || y >= grid.length) return undefined;
    if (x < 0 || x >= grid[0].length) return undefined;

    return grid[y][x];
}

const score = ({ x, y }) => {
    const currentHeight = getHeight({ x, y });

    if (currentHeight === 9) {
        return [[{ x, y }]];
    } else {
        let trails = [];
        if (getHeight({ x: x + 1, y }) === currentHeight + 1) {
            trails = [...trails, ...score({ x: x + 1, y })];
        }
        if (getHeight({ x: x - 1, y }) === currentHeight + 1) {
            trails = [...trails, ...score({ x: x - 1, y })];
        }
        if (getHeight({ x, y: y + 1 }) === currentHeight + 1) {
            trails = [...trails, ...score({ x, y: y + 1 })];
        }
        if (getHeight({ x, y: y - 1 }) === currentHeight + 1) {
            trails = [...trails, ...score({ x, y: y - 1 })];
        }
        return trails.map(t => [{ x, y }, ...t]);
    }
}

const countPeaks = (th) => {
    let reachable = new Set();
    for (const trail of score(th)) {
        const { x, y } = trail.last();
        reachable.add(x + ',' + y);
    }
    return reachable.size;
}

console.log("Part 1", trailheads.map(countPeaks).sum());
console.log("Part 2", trailheads.map(th => score(th).length).sum());
