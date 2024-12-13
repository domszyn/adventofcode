import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

const parseGame = game => {
    const [buttonA, buttonB, prize] = game.split('\n');
    return {
        a: parseGameLine(buttonA),
        b: parseGameLine(buttonB),
        prize: parseGameLine(prize)
    }
}

const parseGameLine = button => {
    const [, coords] = button.split(': ');
    const [x, y] = coords.split(', ');
    return {
        x: Number(x.slice(2)),
        y: Number(y.slice(2))
    }
}

const games = input.split('\n\n').map(parseGame);


const play = ({ a, b, prize }) => {
    const d = a.x * b.y - b.x * a.y;
    const d1 = prize.x * b.y - b.x * prize.y;
    const d2 = a.x * prize.y - prize.x * a.y;

    if (d1 % d !== 0 || d2 % d !== 0) {
        return undefined;
    }

    const minX = d1 / d;
    const minY = d2 / d;

    return 3 * minX + minY;
}


console.log("Part 1", games.map(play).filter(Boolean).sum());
console.log("Part 2", games.map(({ a, b, prize }) => ({
    a, b, prize: {
        x: prize.x + 10000000000000,
        y: prize.y + 10000000000000
    }
})).map(play).filter(Boolean).sum());

