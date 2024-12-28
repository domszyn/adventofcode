import { input } from "./input.js";
import { parseInput } from "../utils.js";

const grid = parseInput(input, l => [...l]);
const infectionStatus = new Set()

for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
        let x = j - (grid[0].length - 1) / 2;
        let y = i - (grid.length - 1) / 2;
        if (grid[i][j] == '#') {
            infectionStatus.add([x, y].join());
        }
    }
}

const directions = ['up', 'right', 'down', 'left'];
const dirRight = dir => directions[(directions.findIndex(d => d == dir) + 1) % directions.length];
const dirLeft = dir => directions[(directions.findIndex(d => d == dir) + 3) % directions.length];

let virusCarrier = {
    loc: [0, 0],
    dir: 'up'
};

let infectedNodes = 0;

for (let i = 0; i < 10000; i++) {
    const currentLoc = virusCarrier.loc.join();
    if (infectionStatus.has(currentLoc)) {
        virusCarrier.dir = dirRight(virusCarrier.dir);
        infectionStatus.delete(currentLoc);
    } else {
        virusCarrier.dir = dirLeft(virusCarrier.dir);
        infectionStatus.add(currentLoc);
        infectedNodes++;
    }

    const [x, y] = virusCarrier.loc;
    switch (virusCarrier.dir) {
        case 'up':
            virusCarrier.loc = [x, y - 1];
            break;
        case 'right':
            virusCarrier.loc = [x + 1, y];
            break;
        case 'down':
            virusCarrier.loc = [x, y + 1];
            break;
        case 'left':
            virusCarrier.loc = [x - 1, y];
            break;
    }
}

console.log(infectedNodes);