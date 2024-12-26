import { input } from "./input.js";
import { parseInput } from '../utils.js';
import { PriorityQueue } from "../minHeap.js";
import '../array.js';

const grid = parseInput(input, l => [...l]);
const getKey = ({ x, y }) => `${x},${y}`;
const walls = new Set();
let start, end;
for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
        if (grid[i][j] === 'S') {
            start = { x: j, y: i };
        }
        if (grid[i][j] === 'E') {
            end = { x: j, y: i };
        }
        if (grid[i][j] === '#') {
            walls.add(getKey({ x: j, y: i }))
        }
    }
}
const dx = [1, 0, -1, 0];
const dy = [0, -1, 0, 1];

const findPath = (start, end, walls) => {
    const queue = new PriorityQueue((a, b) => a.score < b.score);
    queue.push({ cell: start, score: 0, path: [start] });


    const visited = new Map();
    let minPath = []
    while (!queue.isEmpty()) {
        let { cell, score, path } = queue.pop();
        if (cell.x == end.x && cell.y == end.y) {
            minPath = [...path];
            break;
        };

        for (let i = 0; i < 4; i++) {
            let ny = cell.y + dy[i];
            let nx = cell.x + dx[i];
            if (ny < 0 || nx < 0 || ny >= grid.length || nx >= grid[0].length) {
                continue
            };

            const cacheKey = getKey({ x: nx, y: ny });
            let s = score + 1;
            if ((visited.has(cacheKey) && visited.get(cacheKey) <= s)) {
                continue;
            }

            if (!walls.has(cacheKey)) {
                visited.set(cacheKey, s);
                const to = { x: nx, y: ny };
                queue.push({
                    cell: to,
                    score: s,
                    path: [...path, to]
                });
            }
        }
    }

    return [minPath, visited.get(getKey(end))];
}

const [minPath] = findPath(start, end, walls);
const distance = (a, b) => Math.abs(a.x - b.x) + Math.abs(a.y - b.y);

let part1=0, part2=0;
for (let i = 0; i < minPath.length; i++) {
    for (let j = i + 1; j < minPath.length; j++) {
        const m = distance(minPath[i], minPath[j]);
        if (m <= 2 && j - i - m >= 100) {
            part1++;
        }
        if (m <= 20 && j - i - m >= 100) {
            part2++;
        }
    }
}

console.log("Part 1", part1);
console.log("Part 2", part2);