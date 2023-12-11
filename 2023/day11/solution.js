import '../array.js';
import { input } from "./input.js";
import { parseInput } from '../utils.js';

const galaxy = parseInput(input, s => s.split(""));

function expand(galaxy) {
    const extraRows = [];
    const extraColumns = [];
    for (let i = 0; i < galaxy.length; i++) {
        if (galaxy[i].every(c => c == '.')) {
            extraRows.push(i);
        }
    }

    for (let j = 0; j < galaxy[0].length; j++) {
        let emptyColumn = true;
        for (let i = 0; i < galaxy.length; i++) {
            if (galaxy[i][j] == '#') {
                emptyColumn = false;
                break;
            }
        }

        if (emptyColumn) {
            extraColumns.push(j);
        }
    }

    return [extraRows, extraColumns];
}

function findStars(galaxy) {
    const stars = [];
    for (let i = 0; i < galaxy.length; i++) {
        for (let j = 0; j < galaxy[i].length; j++) {
            if (galaxy[i][j] == '#') {
                stars.push([i, j]);
            }
        }
    }
    return stars;
}

function solve(extraRows, extraColumns, expansionFactor) {
    const stars = findStars(galaxy);
    const distances = [];
    for (let i = 0; i < stars.length; i++) {
        for (let j = i + 1; j < stars.length; j++) {
            let distance = 0;
            const rows = [stars[i][0], stars[j][0]].sort((a, b) => a - b);
            const columns = [stars[i][1], stars[j][1]].sort((a, b) => a - b);
            const er = extraRows.filter(er => er > rows[0] && er < rows[1]).length;
            const ec = extraColumns.filter(ec => ec > columns[0] && ec < columns[1]).length;
            distance += rows[1] - rows[0];
            distance += columns[1] - columns[0];
            distance += expansionFactor * er;
            distance += expansionFactor * ec;
            distances.push(distance);
        }
    }
    return distances.sum();
}
const [extraRows, extraColumns] = expand(galaxy);
console.log("Part 1", solve(extraRows, extraColumns, 1));
console.log("Part 2", solve(extraRows, extraColumns, 999999));