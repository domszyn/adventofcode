import { groupIntoRegions } from "../../2024/day12/solution.js";
import { parseInput } from "../utils.js";
import { input } from "./input.js";
import "../array.js";

let points = parseInput(input, l => {
    const vIdx = l.indexOf('v');
    const [loc, vel] = [l.slice(0, vIdx - 1), l.slice(vIdx)].map(t => {
        let coords = t.split('=')[1];
        const [x, y] = coords
            .slice(1, coords.length - 1)
            .split(',')
            .map(s => s.trim(' '))
            .map(Number);

        return { x, y };
    });

    return { loc, vel };
});

const printPoints = (points) => {
    const minX = points.map(p => p.loc.x).min();
    const minY = points.map(p => p.loc.y).min();
    const maxX = points.map(p => p.loc.x).max();
    const maxY = points.map(p => p.loc.y).max();

    const lines = new Array(maxY - minY + 1);
    for (let y = 0; y < lines.length; y++) {
        lines[y] = new Array(maxX - minX + 1).fill('.');
    }

    for (const { x, y } of points.map(p => p.loc)) {
        lines[y - minY][x - minX] = '#';
    }

    console.log(lines.map(l => l.join('')).join('\n'));
}

let numRegions = [];

let timer = 10550;
points = points.map(p => ({
    ...p,
    loc: {
        x: p.loc.x + timer * p.vel.x,
        y: p.loc.y + timer * p.vel.y
    }
}));

while (numRegions.length != 8) {
    timer++;
    points = points.map(p => ({
        ...p,
        loc: {
            x: p.loc.x + p.vel.x,
            y: p.loc.y + p.vel.y
        }
    }));
    numRegions = groupIntoRegions(points.map(p => p.loc), true);
}

printPoints(points);
console.log(timer);