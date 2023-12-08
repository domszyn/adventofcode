import '../array.js';
import { parseInput } from '../utils.js';
import { input } from "./input.js";

let lines = parseInput(input);
const [steps] = lines.splice(0, 2);

const directions = new Map(lines.map(s => {
    const [location, next] = s.split(" = ");
    const [left, right] = next.substring(1, next.length - 1).split(", ");
    return [location, {
        left,
        right
    }];
}));

function countSteps(location) {
    let stepsCounter = 0;

    while (location[2] != 'Z') {
        let i = stepsCounter % steps.length;
        const nextLocation = directions.get(location);
        location = steps[i] == 'L' ? nextLocation.left : nextLocation.right;
        stepsCounter++;
    }

    return stepsCounter;
}

console.log("Part 1", countSteps('AAA'));

const stepsNeeded = [...directions.keys()].filter(s => s[2] == 'A').map(l => countSteps(l));
console.log("Part 2", stepsNeeded.lcm());
