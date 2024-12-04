import { parseInput } from '../utils.js';
import { input } from './input.js';
import '../array.js';

const grid = parseInput(input, line => [...line]);

const getLetter = (i, j) => {
    if (i < 0 || i >= grid.length) return undefined;
    if (j < 0 || j >= grid[0].length) return undefined;

    return grid[i][j];
}

const checks = [
    /* up */(i, j) => getLetter(i, j) === 'X' && getLetter(i - 1, j) === 'M' && getLetter(i - 2, j) === 'A' && getLetter(i - 3, j) === 'S',
    /* up right*/(i, j) => getLetter(i, j) === 'X' && getLetter(i - 1, j + 1) === 'M' && getLetter(i - 2, j + 2) === 'A' && getLetter(i - 3, j + 3) === 'S',
    /*right*/(i, j) => getLetter(i, j) === 'X' && getLetter(i, j + 1) === 'M' && getLetter(i, j + 2) === 'A' && getLetter(i, j + 3) === 'S',
    /*down right*/(i, j) => getLetter(i, j) === 'X' && getLetter(i + 1, j + 1) === 'M' && getLetter(i + 2, j + 2) === 'A' && getLetter(i + 3, j + 3) === 'S',
    /*down*/(i, j) => getLetter(i, j) === 'X' && getLetter(i + 1, j) === 'M' && getLetter(i + 2, j) === 'A' && getLetter(i + 3, j) === 'S',
    /*down left*/(i, j) => getLetter(i, j) === 'X' && getLetter(i + 1, j - 1) === 'M' && getLetter(i + 2, j - 2) === 'A' && getLetter(i + 3, j - 3) === 'S',
    /*left*/(i, j) => getLetter(i, j) === 'X' && getLetter(i, j - 1) === 'M' && getLetter(i, j - 2) === 'A' && getLetter(i, j - 3) === 'S',
    /*up left*/(i, j) => getLetter(i, j) === 'X' && getLetter(i - 1, j - 1) === 'M' && getLetter(i - 2, j - 2) === 'A' && getLetter(i - 3, j - 3) === 'S',
];

const checkXmas = [
    (i, j) => getLetter(i, j) === 'M' && getLetter(i, j + 2) === 'S' && getLetter(i + 1, j + 1) === 'A' && getLetter(i + 2, j) === 'M' && getLetter(i + 2, j + 2) === 'S',
    (i, j) => getLetter(i, j) === 'S' && getLetter(i, j + 2) === 'S' && getLetter(i + 1, j + 1) === 'A' && getLetter(i + 2, j) === 'M' && getLetter(i + 2, j + 2) === 'M',
    (i, j) => getLetter(i, j) === 'S' && getLetter(i, j + 2) === 'M' && getLetter(i + 1, j + 1) === 'A' && getLetter(i + 2, j) === 'S' && getLetter(i + 2, j + 2) === 'M',
    (i, j) => getLetter(i, j) === 'M' && getLetter(i, j + 2) === 'M' && getLetter(i + 1, j + 1) === 'A' && getLetter(i + 2, j) === 'S' && getLetter(i + 2, j + 2) === 'S',
];

let part1 = 0;
let part2 = 0;

for (let row = 0; row < grid.length; row++) {
    for (let column = 0; column < grid[0].length; column++) {
        for (const check of checks) {
            if (check(row, column)) part1++;
        }
        for (const check of checkXmas) {
            if (check(row, column)) part2++;
        }
    }
}

console.log("Part 1", part1);
console.log("Part 2", part2);