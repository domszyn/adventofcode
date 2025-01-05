import { makeArray } from "../utils.js";
import "../array.js";
import { numMarbles, numPlayers } from "./input.js";

const solve = (numPlayers, numMarbles) => {
    const scores = makeArray(numPlayers, 0);
    const marbles = new Array(numMarbles - 1);
    for (let i = 0; i < numMarbles - 1; i++) {
        marbles[i] = i + 1;
    }

    let circle = [0];
    let currentIdx = 0;

    let maxScore = Number.MIN_SAFE_INTEGER;
    let tail = [];
    for (let i = 0; i < marbles.length; i++) {
        const marble = marbles[i];
        if (marble % 23 == 0) {
            currentIdx = (currentIdx + circle.length - 7) % circle.length;
            const [removed] = circle.splice(currentIdx, 1);
            const player = i % numPlayers;
            scores[player] += marble + removed;
            if (scores[player] > maxScore) {
                maxScore = scores[player];
            }
        } else {
            const prevIdx = currentIdx;
            currentIdx = (currentIdx + 1) % circle.length + 1;
            if (currentIdx < prevIdx && tail.length > 0) {
                currentIdx += circle.length;
                circle = [
                    ...circle,
                    ...tail,
                ];
                tail = [];
            }
            circle.splice(currentIdx, 0, marble);
            if ((marble + 1) % 23 == 0 && currentIdx > 7) {
                tail = [
                    ...tail,
                    ...circle.slice(0, currentIdx - 7)
                ];
                circle = circle.slice(currentIdx - 7);
                currentIdx = 7;
            }
        }
    }
    return maxScore;
}

console.log(solve(numPlayers, numMarbles));
console.log(solve(numPlayers, numMarbles * 100));