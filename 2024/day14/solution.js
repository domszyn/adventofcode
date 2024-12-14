import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';
import { groupIntoRegions } from "../day12/solution.js";

const robots = parseInput(input, line => {
    const [p, v] = line.split(' ');
    const [px, py] = p.slice(2).split(',').map(Number);
    const [vx, vy] = v.slice(2).split(',').map(Number);
    return { px, py, vx, vy };
});

const maxX = 101, maxY = 103;
const mx = (maxX - 1) / 2, my = (maxY - 1) / 2;

const printRobots = (seconds) => {
    const grid = new Array(maxY);
    for (let i = 0; i < grid.length; i++) {
        grid[i] = new Array(maxX);
        for (let j = 0; j < maxX; j++) {
            grid[i][j] = '.';
        }
    }
    for (const r of robots) {
        grid[r.py][r.px] = '*';
    }

    console.log(grid.map(l => l.join('')).join('\n'));
}

for (let i = 0; i < 100000; i++) {
    for (let j = 0; j < robots.length; j++) {
        const { px, py, vx, vy } = robots[j];
        let x = px + vx, y = py + vy;
        if (x < 0) x += maxX;
        if (x >= maxX) x %= maxX;
        if (y < 0) y += maxY;
        if (y >= maxY) y %= maxY;
        robots[j].px = x;
        robots[j].py = y;
    }

    if (i == 99) {
        const counts = [
            robots.filter(({ px, py }) => px < mx && py < my).length,
            robots.filter(({ px, py }) => px > mx && py < my).length,
            robots.filter(({ px, py }) => px < mx && py > my).length,
            robots.filter(({ px, py }) => px > mx && py > my).length,
        ];

        console.log('Part 1', counts.multiply());
    }

    const regions = groupIntoRegions(robots.map(({ px, py }) => ({ x: px, y: py })));
    if (regions.some(r => r.length > 50)) {
        console.log("Part 2", i + 1);
        printRobots(i + 1);
        break;
    }
}
