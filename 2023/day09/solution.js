import '../array.js';
import { makeArray, parseInput } from '../utils.js';
import { input } from "./input.js";

const historicalValues = parseInput(input, s => s.split(' ').map(s => parseInt(s)));

function predictValue(history, backwards) {
    let triangle = [history];

    while (triangle[0].some(n => n != 0)) {
        const tmp = makeArray(triangle[0].length - 1);
        for (let i = 0; i < tmp.length; i++) {
            tmp[i] = triangle[0][i + 1] - triangle[0][i];
        }
        triangle = [
            tmp,
            ...triangle
        ];
    }

    for (let i = 0; i < triangle.length; i++) {
        if (i == 0) {
            triangle[i] = [0, ...triangle[i], 0];
        } else {
            triangle[i] = [triangle[i][0] - triangle[i - 1][0], ...triangle[i], triangle[i][triangle[i].length - 1] + triangle[i - 1][triangle[i - 1].length - 1]];
        }
    }

    const historyWithPrediction = triangle[triangle.length - 1];
    return backwards ? historyWithPrediction[0] : historyWithPrediction[historyWithPrediction.length - 1];
}

console.log("Part 1", historicalValues.map(history => predictValue(history)).sum());
console.log("Part 2", historicalValues.map(history => predictValue(history, true)).sum());

