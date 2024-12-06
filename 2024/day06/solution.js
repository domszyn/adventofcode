import { parseInput } from "../utils.js";
import { input } from "./input.js";
import '../array.js';

let posColumn = 0, posRow = 0;
const grid = parseInput(input, (l, idx) => {
    if (l.indexOf('^') >= 0) {
        posColumn = l.indexOf('^');
        posRow = idx;
    }
    return [...l];
});

const loc = (row, column) => row + ',' + column;

const move = (from, direction) => {
    let to = { ...from };
    switch (direction) {
        case 0:
            to.row--;
            break;
        case 1:
            to.column++;
            break;
        case 2:
            to.row++;
            break;
        case 3:
            to.column--;
            break;
    }

    if (to.row < 0 || to.row >= grid.length || to.column < 0 || to.column >= grid[0].length) {
        return {
            finished: true,
            turned: false,
            direction,
            position: from
        };
    }

    if (grid[to.row][to.column] !== '#') {
        return {
            finished: false,
            turned: false,
            direction,
            position: to,
        }
    } else {
        return {
            finished: false,
            turned: true,
            direction: (direction + 1) % 4,
            position: from
        }
    }
}

const start = ({ row, column, countVisited }) => {
    let visited = new Set();
    let turnPoints = [];
    let direction = 0;
    while (true) {
        const pos = loc(row, column);
        if (countVisited === true) {
            visited.add(pos);
        }
        const step = move({ row, column }, direction);
        if (step.finished) {
            break;
        }

        if (step.turned) {
            turnPoints.push(pos);
            direction = step.direction
        } else {
            row = step.position.row;
            column = step.position.column;
        }

        if (turnPoints.length > 8) {
            for (let i = 8; i < turnPoints.length; i += 8) {
                let tail = turnPoints.slice(-i);
                let left = tail.slice(0, i / 2);
                let right = tail.slice(-i / 2);
                if (left.eq(right)) {
                    return [true, right];
                }

            }
        }
    }

    return [false, visited.size];
}

const [, visited] = start({ row: posRow, column: posColumn, countVisited: true });
console.log("Part 1", visited);

let part2 = 0;
for (let row = 0; row < grid.length; row++) {
    for (let column = 0; column < grid[0].length; column++) {
        if (grid[row][column] !== '.') {
            continue;
        }

        grid[row][column] = '#';
        const [stuck] = start({ row: posRow, column: posColumn });
        if (stuck) {
            part2++;
        }
        grid[row][column] = '.';
    }
}

console.log("Part 2", part2);
