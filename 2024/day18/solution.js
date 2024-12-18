import { parseInput } from "../utils.js";
import { input } from "./input.js";
import { PriorityQueue } from "../minHeap.js";

const getKey = ({ x, y }) => `${x},${y}`;
const corrupted = parseInput(input, l => {
    const [x, y] = l.split(',').map(Number);
    return { x, y };
});


const findPath = (from, end, corruptedMap) => {
    const queue = new PriorityQueue((a, b) => a.score < b.score)
    queue.push({ cell: from, score: 0 });
    const dx = [1, 0, -1, 0];
    const dy = [0, -1, 0, 1];

    const visited = new Map();
    while (!queue.isEmpty()) {
        const { cell, score, } = queue.pop();
        if (cell.x == end.x && cell.y == end.y) {
            break;
        };

        for (let i = 0; i < 4; i++) {
            const ny = cell.y + dy[i];
            const nx = cell.x + dx[i];
            if (ny < 0 || nx < 0 || ny > end.y || nx > end.x) {
                continue
            };

            const cacheKey = getKey({ x: nx, y: ny });
            if (corruptedMap.has(cacheKey) || (visited.has(cacheKey) && visited.get(cacheKey) <= score + 1)) {
                continue;
            }
            visited.set(cacheKey, score + 1);

            queue.push({
                cell: { x: nx, y: ny },
                score: score + 1,
            });
        }
    }


    return visited.get(getKey(end));
}

const corruptedMap = new Set(corrupted.slice(0, 1024).map(c => getKey(c)));
console.log("Part 1", findPath({ x: 0, y: 0 }, { x: 70, y: 70 }, corruptedMap));

for (let i = 1024; i < corrupted.length; i++) {
    const nextCorrupted = getKey(corrupted[i]);
    corruptedMap.add(nextCorrupted);
    if (findPath({ x: 0, y: 0 }, { x: 70, y: 70 }, corruptedMap) == undefined) {
        console.log("Part 2", nextCorrupted);
        break;
    }
}