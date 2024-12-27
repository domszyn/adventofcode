import { diagram } from "./input.js";

const move = ({ row, column, direction }) => {
    let ray = '';
    switch (direction) {
        case 'down':
            ray = diagram.slice(row + 1).map(r => r[column]).join('');
            break;
        case 'up':
            ray = diagram.slice(0, row).map(r => r[column]).reverse().join('');
            break;
        case 'left':
            ray = [...diagram[row].slice(0, column)].reverse().join('');
            break;
        case 'right':
            ray = diagram[row].slice(column + 1);
            break;
    }
    if (!ray.includes('+')) {
        const index = [...ray].findIndex(c => c >= 'A' && c <= 'Z');
        return [, ray[index], index + 1];
    }

    ray = ray.substring(0, ray.indexOf('+') + 1);
    const letter = [...ray].find(c => c >= 'A' && c <= 'Z');

    switch (direction) {
        case 'down':
            row += ray.length;
            if (column > 0 && diagram[row][column - 1] == '-') {
                direction = 'left';
            } else if (column < diagram[row].length - 1 && diagram[row][column + 1] == '-') {
                direction = 'right';
            }
            break;
        case 'up':
            row -= ray.length;
            if (column > 0 && diagram[row][column - 1] == '-') {
                direction = 'left';
            } else if (column < diagram[row].length - 1 && diagram[row][column + 1] == '-') {
                direction = 'right';
            }
            break;
        case 'left':
            column -= ray.length;
            if (row > 0 && diagram[row - 1][column] == '|') {
                direction = 'up';
            } else if (row < diagram.length - 1 && diagram[row + 1][column] == '|') {
                direction = 'down';
            }
            break;
        case 'right':
            column += ray.length;
            if (row > 0 && diagram[row - 1][column] == '|') {
                direction = 'up';
            } else if (row < diagram.length - 1 && diagram[row + 1][column] == '|') {
                direction = 'down';
            }
            break;
    }

    return [
        { row, column, direction },
        letter,
        ray.length
    ]
};

let loc = {
    row: 0,
    column: diagram[0].indexOf('|'),
    direction: 'down'
};

let letters = '';
let steps = 1;

while (loc) {
    let letter, count;
    [loc, letter, count] = move(loc);
    if (letter) {
        letters += letter;
    }
    steps += count;
}

console.log("Part 1", letters);
console.log("Part 2", steps);