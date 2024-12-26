import { input } from "./input.js";
import { knotHash } from "../day10/solution.js";

let used = [];
for (let i = 0; i < 128; i++) {
    let hash = knotHash(`${input}-${i}`, 64, true);
    for (let j = 0; j < hash.length; j++) {
        const bits = parseInt(hash[j], 16).toString(2).padStart(4, '0');
        for (let k = 0; k < 4; k++) {
            if (bits[k] == '1') {
                used.push({ i, j: 4 * j + k });
            }
        }
    }
}

console.log("Part 1", used.length);

const regions = [];
while (used.length > 0) {
    let square = used[0];
    used = used.slice(1);
    let region = [square];
    let neighborsFound;
    do {
        neighborsFound = false;
        for (const s of region) {
            const neighbors = [
                { i: s.i - 1, j: s.j },
                { i: s.i + 1, j: s.j },
                { i: s.i, j: s.j - 1 },
                { i: s.i, j: s.j + 1 },
            ];
            for (const n of neighbors) {
                const idx = used.findIndex(({ i, j }) => i === n.i && j === n.j);
                if (idx >= 0) {
                    neighborsFound = true;
                    region = [...region, n];
                    used.splice(idx, 1);
                }
            }
        }
    } while (neighborsFound);
    regions.push(region);
}

console.log("Part 2", regions.length);
