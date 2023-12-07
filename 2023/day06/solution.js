import '../array.js';
import { isDigit, makeArray, parseInput } from '../utils.js';
import { input } from './input.js';

const [times, distances] = parseInput(input, s =>
    s.split(" ").filter(s => s != '').slice(1).map(s => parseInt(s))
);
const races = makeArray(times.length);
for (let i = 0; i < times.length; i++) {
    races[i] = { time: times[i], distance: distances[i] };
}

const numWins = r => {
    let wins = 0;
    for (let wait = 1; wait < r.time; wait++) {
        const distance = (r.time - wait) * wait;
        if (distance > r.distance) {
            wins++;
        }
    }
    return wins;
};

console.log("Part 1", races.map(numWins).multiply());

const [time, distance] = parseInput(input, s => 
    parseInt(s.split('').filter(isDigit).join(""))
);
console.log("Part 2", numWins({ time, distance }));
