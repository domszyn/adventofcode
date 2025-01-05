import { gridSerialNumber } from "./input.js";

const grid = new Array(300);
for (let y = 0; y < grid.length; y++) {
    grid[y] = new Array(300);
    for (let x = 0; x < grid[y].length; x++) {
        const rackId = x + 11;
        const powerLevel = rackId * (rackId * (y + 1) + gridSerialNumber);
        const hundreds = Math.floor((powerLevel % 1000) / 100);
        grid[y][x] = hundreds - 5;
    }
}

const getPower = (grid, startX, startY, size) => {
    let power = 0;
    for (let y = startY; y < startY + size; y++) {
        for (let x = startX; x < startX + size; x++) {
            power += grid[y][x];
        }
    }

    return power;
}

let maxPower3x3 = -45;
let maxCell3x3 = [];
let maxPower = -45;
let maxCell = [];
for (let size = 1; size <= 300; size++) {
    for (let y = 0; y < grid.length - size + 1; y++) {
        for (let x = 0; x < grid[y].length - size + 1; x++) {
            const power = getPower(grid, x, y, size);
            if (power > maxPower) {
                maxPower = power;
                maxCell = [x + 1, y + 1, size];
            }
            if (size == 3 && power > maxPower3x3) {
                maxPower3x3 = power;
                maxCell3x3 = [x + 1, y + 1];
            }
        }
    }
}

console.log(maxCell3x3.join());
console.log(maxCell.join());