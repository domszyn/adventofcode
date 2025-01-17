import { input } from "./input.js";
import { parseInput } from "../utils.js";

let carts = [];
const tracks = parseInput(input, (l, y) => {
    const row = [...l];
    for (let x = 0; x < row.length; x++) {
        switch (row[x]) {
            case '^':
            case 'v':
                carts.push({ x, y, dir: row[x], int: 0, id: carts.length });
                row[x] = '|';
                break;
            case '<':
            case '>':
                carts.push({ x, y, dir: row[x], int: 0, id: carts.length });
                row[x] = '-';
                break;
        }
    }
    return row;
});

let crashes = [];
race: while (carts.length > 1) {
    carts.sort((a, b) => a.y * 10000 + a.x - b.y * 1000 - b.x);

    for (let i = 0; i < carts.length; i++) {
        let cart = carts[i];
        if (cart.crashed) {
            continue;
        };

        switch (cart.dir) {
            case '^':
                cart.y--;
                break;
            case 'v':
                cart.y++;
                break;
            case '<':
                cart.x--;
                break;
            case '>':
                cart.x++;
                break;
        }

        switch (tracks[cart.y][cart.x]) {
            case '/':
                switch (cart.dir) {
                    case '^':
                        cart.dir = '>';
                        break;
                    case 'v':
                        cart.dir = '<';
                        break;
                    case '<':
                        cart.dir = 'v';
                        break;
                    case '>':
                        cart.dir = '^';
                        break;
                }
                break;
            case '\\':
                switch (cart.dir) {
                    case '^':
                        cart.dir = '<';
                        break;
                    case 'v':
                        cart.dir = '>';
                        break;
                    case '<':
                        cart.dir = '^';
                        break;
                    case '>':
                        cart.dir = 'v';
                        break;
                }
                break;
            case '+':
                switch (cart.int % 3) {
                    case 0:
                        switch (cart.dir) {
                            case '^':
                                cart.dir = '<';
                                break;
                            case 'v':
                                cart.dir = '>';
                                break;
                            case '<':
                                cart.dir = 'v';
                                break;
                            case '>':
                                cart.dir = '^';
                                break;
                        }
                        break;
                    case 2:
                        switch (cart.dir) {
                            case '^':
                                cart.dir = '>';
                                break;
                            case 'v':
                                cart.dir = '<';
                                break;
                            case '<':
                                cart.dir = '^';
                                break;
                            case '>':
                                cart.dir = 'v';
                                break;
                        }
                        break;
                }
                cart.int++;
                break;
        }

        if (carts.filter(c => c.x == cart.x && c.y == cart.y).length > 1) {
            carts = carts.map(c => ({
                ...c,
                crashed: c.crashed || (c.x == cart.x && c.y == cart.y)
            }));
            crashes.push([cart.x, cart.y].join());
        }
    }

    carts = carts.filter(c => !c.crashed);
}

console.log("Part 1", crashes[0]);
console.log("Part 2", [carts[0].x, carts[0].y].join());