import '../array.js';
import { parseInput } from '../utils.js';
import { input } from './input.js';

function countHand(hand) {
    const counts = new Map();
    for (let i = 0; i < hand.length; i++) {
        const c = hand[i];
        counts.set(c, (counts.get(c) || 0) + 1);
    }

    return [[...counts.values()].sort((a, b) => b - a), counts.get('1') || 0];
}

function getType(hand) {
    const [counts, jokers] = countHand(hand);
    if (counts.eq([5])) {
        return 7;
    }
    if (counts.eq([4, 1])) {
        return jokers > 0 ? 7 : 6;
    }
    if (counts.eq([3, 2])) {
        return jokers > 0 ? 7 : 5;
    }
    if (counts.eq([3, 1, 1])) {
        return jokers > 0 ? 6 : 4;
    }
    if (counts.eq([2, 2, 1])) {
        return jokers > 0 ? jokers + 4 : 3;
    }
    if (counts.eq([2, 1, 1, 1])) {
        return jokers > 0 ? 4 : 2;
    }
    if (counts.eq([1, 1, 1, 1, 1])) {
        return jokers > 0 ? 2 : 1;
    }
    return 0;
}

const mapHand = (s, withJokers) => {
    const [hand, bidAmount] = s.split(" ");
    let comparableHand = hand
        .replaceAll("A", "X")
        .replaceAll("K", "W")
        .replaceAll("Q", "V")
        .replaceAll("J", withJokers ? "1" : "U");

    return {
        hand: comparableHand,
        bid: parseInt(bidAmount),
        type: getType(comparableHand),
    };
}

const compareHands = (a, b) => {
    if (a.type == b.type) {
        return a.hand > b.hand ? 1 : -1;
    } else {
        return a.type - b.type;
    }
};

const getScore = ({ bid }, idx) => bid * (idx + 1);

const hands = parseInput(input, s => mapHand(s, false)).sort(compareHands);
const handsWithJokers = parseInput(input, s => mapHand(s, true)).sort(compareHands);

console.log("Part 1:", hands.map(getScore).sum());
console.log("Part 2:", handsWithJokers.map(getScore).sum());