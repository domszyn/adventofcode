
const input = ``;

var rows = input.split('\n');
var cards = new Array(rows.length).fill(1);

const getNumbers = c => c.split(' ').filter(s => s != '').map(s => parseInt(s.trim()));
const sum = arr => arr.reduce((a, b) => a + b, 0);

const part1 = sum(rows.map((s, idx) => {
    const [, card] = s.split(': ');
    const [c1, c2] = card.split(' | ');
    const winningNumbers = new Set(getNumbers(c1));
    const numbers = getNumbers(c2);
    var matches = numbers.filter(n => winningNumbers.has(n)).length;

    for (let i = 1; i <= matches && i < cards.length - idx; i++) {
        cards[idx + i] += cards[idx];
    }

    return matches == 0 ? 0 : Math.pow(2, matches - 1);
}));

console.log("Part 1", part1);
console.log("Part 2", sum(cards));