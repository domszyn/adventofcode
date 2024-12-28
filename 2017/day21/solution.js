import { input } from "./input.js";
import { parseInput } from '../utils.js';

const rules = new Map(parseInput(input, l => l.split(" => ")));

const fold = pattern => pattern.map(p => p.join('')).join('/');
const unfold = pattern => pattern.split('/').map(l => [...l]);

const split = pattern => {
    let squareSize = pattern.length % 2 == 0 ? 2 : 3;
    let numSquares = pattern.length / squareSize;
    const squares = new Array(numSquares * numSquares).fill('');

    for (let i = 0; i < pattern.length; i += squareSize) {
        for (let j = 0; j < pattern.length; j += squareSize) {
            const idx = (i / squareSize) * numSquares + (j / squareSize);
            squares[idx] = pattern
                .slice(i, i + squareSize)
                .map(p => p.slice(j, j + squareSize).join(''))
                .join('/');
        }
    }

    return squares;
}

const join = squares => {
    if (squares.length == 1) {
        return squares[0];
    }

    const squareSize = squares[0].length;
    const squaresPerRow = Math.round(Math.sqrt(squares.length));
    const pattern = new Array(squareSize * squaresPerRow);
    for (let i = 0; i < squares.length; i++) {
        let row = Math.floor(i / squaresPerRow);
        let column = i % squaresPerRow;
        for (let j = 0; j < squareSize; j++) {
            const rj = row * squareSize + j;
            if (pattern[rj] === undefined) {
                pattern[rj] = new Array(squareSize * squaresPerRow);
            }
            for (let k = 0; k < squareSize; k++) {
                const rc = column * squareSize + k;
                pattern[rj][rc] = squares[i][j][k];
            }
        }
    }
    return pattern;
}

const flipH = square => fold(unfold(square).map(p => p.reverse()));
const rotateRight = square => {
    const pattern = unfold(square);
    const rotated = new Array(pattern.length);
    for (let i = 0; i < rotated.length; i++) {
        rotated[i] = pattern.map(r => r[i]).reverse();
    }
    return fold(rotated);
}

for (const input of rules.keys()) {
    const output = rules.get(input);
    const rr = rotateRight(input);
    const fh = flipH(input);

    if (!rules.has(rr)) {
        rules.set(rr, output);
    }

    if (!rules.has(fh)) {
        rules.set(fh, output);
    }
}

const countPixelsOn = pattern => [...fold(pattern)].filter(c => c == '#').length;

let part1;
let pattern = unfold(".#./..#/###");
for (let i = 0; i < 18; i++) {
    if (i == 5) {
        part1 = countPixelsOn(pattern);
    }
    const squares = split(pattern);
    const enhancedSquares = squares
        .map(s => rules.get(s))
        .map(s => unfold(s));
    pattern = join(enhancedSquares);
}

console.log("Part 1", part1);
console.log("Part 2", countPixelsOn(pattern));