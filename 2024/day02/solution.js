import '../array.js';
import { parseInput } from '../utils.js';
import { input } from './input.js';

const isSafe = (levels) => {
    let increasing = false;
    for (let i = 0; i < levels.length - 1; i++) {
        const diff = levels[i + 1] - levels[i]
        if (diff > 0 && diff <= 3) {
            if (i == 0) {
                increasing = true
            } else if (!increasing) {
                return false
            }
        } else if (diff >= -3 && diff < 0) {
            if (i == 0) {
                increasing = false
            } else if (increasing) {
                return false
            }
        } else {
            return false
        }
    }

    return true
}

const isSafeWithTolerate = (levels) => levels
    .map((_, i) => levels.toSpliced(i, 1))
    .some(tl => isSafe(tl));

let reports = parseInput(input, report => report.split(' ').map(l => +l));

console.log("Part 1", reports.map(levels => isSafe(levels)).filter(_ => _).length);
console.log("Part 2", reports.map(levels => isSafeWithTolerate(levels)).filter(_ => _).length);