import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

const grid = parseInput(input, l => [...l]);

const plants = new Map();
for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
        const plant = grid[i][j];
        if (!plants.has(plant)) {
            plants.set(plant, []);
        }
        plants.set(plant, [...plants.get(plant), { x: j, y: i }]);
    }
}

const getNeighbors = (p) => [
    { x: p.x - 1, y: p.y },
    { x: p.x + 1, y: p.y },
    { x: p.x, y: p.y - 1 },
    { x: p.x, y: p.y + 1 },
];

const groupPlantsIntoRegions = (plants) => {
    const regions = [];

    while (plants.length > 0) {
        let p = plants[0];
        plants = plants.slice(1);
        let region = [p];
        let neighborsFound;
        do {
            neighborsFound = false;
            for (const p of region) {
                for (const n of getNeighbors(p)) {
                    const idx = plants.findIndex(({ x, y }) => x === n.x && y === n.y);
                    if (idx >= 0) {
                        neighborsFound = true;
                        region = [...region, n];
                        plants.splice(idx, 1);
                    }
                }
            }
        } while (neighborsFound);
        regions.push(region);
    }

    return regions;
}

const getPlant = ({ x, y }) => {
    if (x < 0 || x >= grid[0].length || y < 0 || y >= grid.length) {
        return undefined;
    }

    return grid[y][x];
}

const getPerimeterOld = (region) => {
    const plant = getPlant(region[0]);
    return region.map(p => {
        const neighbors = getNeighbors(p);
        let diffPlantCount = 0;
        for (const n of neighbors) {
            if (getPlant(n) !== plant) {
                diffPlantCount++;
            }
        }

        return diffPlantCount;
    }).sum();
}

const getPerimeter = (region, compact) => {
    let vectors = [];
    for (const { x, y } of region) {
        const rvectors = [
            { from: { x, y }, to: { x: x + 1, y } },
            { from: { x: x + 1, y }, to: { x: x + 1, y: y - 1 } },
            { from: { x: x + 1, y: y - 1 }, to: { x, y: y - 1 } },
            { from: { x, y: y - 1 }, to: { x, y } }
        ];

        for (const rv of rvectors) {
            const idx = vectors.findIndex(v => v.to.x === rv.from.x &&
                v.to.y == rv.from.y &&
                v.from.x === rv.to.x &&
                v.from.y == rv.to.y);
            if (idx < 0) {
                vectors.push(rv);
            } else {
                vectors.splice(idx, 1);
            }
        }
    }


    if (compact) {
        let connected;
        do {
            connected = false;

            for (let i = 0; i < vectors.length; i++) {
                const v = vectors[i];
                let idx = -1;

                if (v.from.x === v.to.x) {
                    idx = vectors.findIndex(rv => rv.from.x === rv.to.x &&
                        v.to.x === rv.from.x &&
                        v.to.y == rv.from.y);
                } else if (v.from.y == v.to.y) {
                    idx = vectors.findIndex(rv => rv.from.y === rv.to.y &&
                        v.to.x === rv.from.x &&
                        v.to.y == rv.from.y);
                }

                if (idx >= 0) {
                    vectors[i].to = { ...vectors[idx].to };
                    vectors.splice(idx, 1);
                    connected = true;
                    break;
                }
            }

        } while (connected)
    }


    return vectors.length;
}

const plantRegions = [...plants.keys()].map(p => groupPlantsIntoRegions(plants.get(p)));

console.log('Part 1', plantRegions.map(pr => pr.map(r => getPerimeter(r) * r.length).sum()).sum());

console.log('Part 2', plantRegions.map(pr => pr.map(r => getPerimeter(r, true) * r.length).sum()).sum());
