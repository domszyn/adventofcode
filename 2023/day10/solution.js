import '../array.js';
import { makeArray, parseInput } from "../utils.js";
import { input } from "./input.js";

const verticalPipe = '|';
const horizontalPipe = '-';
const bend90ne = 'L';
const bend90nw = 'J';
const bend90sw = '7';
const bend90se = 'F';
const ground = '.';
const start = 'S';

const grid = parseInput(input, s => s.split(""));
const rows = grid.length;
const columns = grid[0].length;
let startLocation;
for (let i = 0; i < rows; i++) {
    let idx = grid[i].indexOf(start);
    if (idx >= 0) {
        startLocation = { row: i, column: idx };
    }
}

const getConnectedTiles = (row, column) => {
    const pipeType = grid[row][column];
    switch (pipeType) {
        case verticalPipe:
            return [
                { row: row - 1, column },
                { row: row + 1, column }
            ];
        case horizontalPipe:
            return [
                { row, column: column - 1 },
                { row, column: column + 1 }
            ];
        case bend90ne:
            return [
                { row: row - 1, column },
                { row, column: column + 1 }
            ];
        case bend90nw:
            return [
                { row: row - 1, column },
                { row, column: column - 1 }
            ];
        case bend90sw:
            return [
                { row: row + 1, column },
                { row, column: column - 1 }
            ];
        case bend90se:
            return [
                { row: row + 1, column },
                { row, column: column + 1 }
            ];
    }
    return [];
}



function buildLoop() {
    const { row: startRow, column: startColumn } = startLocation;
    const [leftOfStart] = [
        { row: startRow - 1, column: startColumn },
        { row: startRow + 1, column: startColumn },
        { row: startRow, column: startColumn - 1 },
        { row: startRow, column: startColumn + 1 },
    ].filter(({ row, column }) => {
        if (row >= 0 && row < rows && column >= 0 && column < columns) {
            for (const tile of getConnectedTiles(row, column)) {
                if (tile.row == startRow && tile.column == startColumn) {
                    return true;
                }
            }
        }

        return false;
    });

    let loop = [startLocation, leftOfStart];
    while (true) {
        const prevLocation = loop[loop.length - 2];
        const currentLocation = loop.last();
        const nextLocation = getConnectedTiles(currentLocation.row, currentLocation.column)
            .find(({ row, column }) => row != prevLocation.row || column != prevLocation.column);
        if (nextLocation.row == startLocation.row && nextLocation.column == startLocation.column) {
            break;
        } else {
            loop = [...loop, nextLocation];
        }
    }

    return loop;
}

const loop = buildLoop();
console.log("Part 1", loop.length / 2);

const cleanGrid = makeArray(rows);
for (let i = 0; i < rows; i++) {
    cleanGrid[i] = makeArray(columns, '.');
}

for (const { row, column } of loop) {
    cleanGrid[row][column] = grid[row][column];
}

let lockedTiles = 0;
for (let i = 0; i < rows; i++) {
    for (let j = 0; j < columns; j++) {
        if (cleanGrid[i][j] != ground) {
            continue;
        }

        const intersections = cleanGrid[i].slice(j + 1).filter(c => c != ground).join("")
            .replaceAll(/F\-*7/g, "II")
            .replaceAll(/L\-*J/g, "II")
            .replaceAll(/L\-*7/g, "I")
            .replaceAll(/F\-*J/g, "I");
            
        if (intersections.length % 2 == 1) {
            lockedTiles++;
        }
    }
}

console.log("Part 2", lockedTiles);
