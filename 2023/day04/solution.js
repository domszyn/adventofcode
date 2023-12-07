import '../array.js';
import { makeArray, parseInput } from '../utils.js';
import { input } from './input.js';

var rows = parseInput(input);
var cards = makeArray(rows.length, 1);

const getNumbers = c => c.split(' ').filter(s => s != '').map(s => parseInt(s.trim()));

const matches = rows.map((s, idx) => {
    const [, card] = s.split(': ');
    const [c1, c2] = card.split(' | ');
    const winningNumbers = new Set(getNumbers(c1));
    const numbers = getNumbers(c2);
    var matches = numbers.filter(n => winningNumbers.has(n)).length;

    for (let i = 1; i <= matches && i < cards.length - idx; i++) {
        cards[idx + i] += cards[idx];
    }

    return matches == 0 ? 0 : Math.pow(2, matches - 1);
});

console.log("Part 1", matches.sum());
console.log("Part 2", cards.sum());