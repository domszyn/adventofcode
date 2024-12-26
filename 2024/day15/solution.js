import { parseInput } from '../utils.js';
import { movements, warehouse } from './input.js';

const warehouseLayout = parseInput(warehouse, l => [...l]);

let robot = { i: 0, j: 0 };

class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }
}

class Warehouse {
    #grid
    #robot

    #getLoc(i, j) {
        if (i < 0 || i >= this.#grid.length) return undefined;
        if (j < 0 || j >= this.#grid[0].length) return undefined;

        return this.#grid[i][j];
    }

    constructor(input) {
        this.#grid = parseInput(input, l => [...l]);
        for (let i = 0; i < this.#grid.length; i++) {
            for (let j = 0; j < this.#grid[i].length; j++) {
                if (this.#grid[i][j] === '@') {
                    this.#robot = { i, j };
                }
            }
        }
    }

    get score() {
        let s = 0;
        for (let i = 0; i < this.#grid.length; i++) {
            for (let j = 0; j < this.#grid[i].length; j++) {
                const e = this.#getLoc(i, j)
                if (e === 'O' || e === '[') {
                    s += 100 * i + j;
                }
            }
        }
        return s;
    }

    #findNumBoxesToMoveUp(i, j) {
        let boxes = 0;
        let up;
        do {
            i--;
            up = this.#getLoc(i, j);
            if (up === 'O') {
                boxes++;
            }
        } while (up === 'O');

        if (up === '.') {
            return boxes;
        } else {
            return -1;
        }
    }

    #findNumBoxesToMoveDown = (i, j) => {
        let boxes = 0;
        let down;
        do {
            i++;
            down = this.#getLoc(i, j);
            if (down === 'O') {
                boxes++;
            }
        } while (down === 'O');

        if (down === '.') {
            return boxes;
        } else {
            return -1;
        }
    }

    #findNumBoxesToMoveLeft = (i, j) => {
        let boxes = 0;
        let left;
        do {
            j--;
            left = this.#getLoc(i, j);
            if (left === 'O') {
                boxes++;
            }
        } while (left === 'O');

        if (left === '.') {
            return boxes;
        } else {
            return -1;
        }
    }

    #findNumBoxesToMoveRight = (i, j) => {
        let boxes = 0;
        let right;
        do {
            j++;
            right = this.#getLoc(i, j);
            if (right === 'O') {
                boxes++;
            }
        } while (right === 'O');

        if (right === '.') {
            return boxes;
        } else {
            return -1;
        }
    }

    #moveUp() {
        const numBoxes = this.#findNumBoxesToMoveUp(this.#robot.i, this.#robot.j);
        if (numBoxes === -1) {
            return;
        }

        if (numBoxes > 0) {
            this.#grid[this.#robot.i - numBoxes - 1][this.#robot.j] = 'O';
        }
        this.#grid[this.#robot.i - 1][this.#robot.j] = '@';
        this.#grid[this.#robot.i][this.#robot.j] = '.';
        this.#robot.i--;
    }

    #moveDown() {
        const numBoxes = this.#findNumBoxesToMoveDown(this.#robot.i, this.#robot.j);
        if (numBoxes === -1) {
            return;
        }
        if (numBoxes > 0) {
            this.#grid[this.#robot.i + numBoxes + 1][this.#robot.j] = 'O';
        }
        this.#grid[this.#robot.i + 1][this.#robot.j] = '@';
        this.#grid[this.#robot.i][this.#robot.j] = '.';
        this.#robot.i++;
    }

    #moveLeft() {
        const numBoxes = this.#findNumBoxesToMoveLeft(this.#robot.i, this.#robot.j);
        if (numBoxes === -1) {
            return;
        }
        if (numBoxes > 0) {
            this.#grid[this.#robot.i][this.#robot.j - numBoxes - 1] = 'O';
        }
        this.#grid[this.#robot.i][this.#robot.j - 1] = '@';
        this.#grid[this.#robot.i][this.#robot.j] = '.';
        this.#robot.j--;
    }

    #moveRight() {
        const numBoxes = this.#findNumBoxesToMoveRight(this.#robot.i, this.#robot.j);
        if (numBoxes === -1) {
            return;
        }
        if (numBoxes > 0) {
            this.#grid[this.#robot.i][this.#robot.j + numBoxes + 1] = 'O';
        }
        this.#grid[this.#robot.i][this.#robot.j + 1] = '@';
        this.#grid[this.#robot.i][this.#robot.j] = '.';
        this.#robot.j++;
    }

    move(dir) {
        switch (dir) {
            case '^':
                this.#moveUp();
                break;
            case 'v':
                this.#moveDown();
                break;
            case '<':
                this.#moveLeft();
                break;
            case '>':
                this.#moveRight();
                break;
        }
    }
}

const getLoc = (grid, i, j) => {
    if (i < 0 || i >= grid.length) return undefined;
    if (j < 0 || j >= grid[0].length) return undefined;

    return grid[i][j];
}

for (let i = 0; i < warehouseLayout.length; i++) {
    for (let j = 0; j < warehouseLayout[i].length; j++) {
        if (getLoc(warehouseLayout, i, j) === '@') {
            robot.i = i;
            robot.j = j;
        }
    }
}

const getScore = wh => {
    let score = 0;
    for (let i = 0; i < wh.length; i++) {
        for (let j = 0; j < wh[i].length; j++) {
            const e = getLoc(wh, i, j)
            if (e === 'O' || e === '[') {
                score += 100 * i;
                score += j;
            }
        }
    }
    return score;
}

const findNumBoxesToMoveUp = (warehouse, i, j) => {
    let boxes = 0;
    let up;
    do {
        i--;
        up = getLoc(warehouse, i, j);
        if (up === 'O') {
            boxes++;
        }
    } while (up === 'O');

    if (up === '.') {
        return boxes;
    } else {
        return -1;
    }
}

const findNumBoxesToMoveDown = (warehouse, i, j) => {
    let boxes = 0;
    let down;
    do {
        i++;
        down = getLoc(warehouse, i, j);
        if (down === 'O') {
            boxes++;
        }
    } while (down === 'O');

    if (down === '.') {
        return boxes;
    } else {
        return -1;
    }
}

const findNumBoxesToMoveLeft = (warehouse, i, j) => {
    let boxes = 0;
    let left;
    do {
        j--;
        left = getLoc(warehouse, i, j);
        if (left === 'O') {
            boxes++;
        }
    } while (left === 'O');

    if (left === '.') {
        return boxes;
    } else {
        return -1;
    }
}

const findNumBoxesToMoveRight = (warehouse, i, j) => {
    let boxes = 0;
    let right;
    do {
        j++;
        right = getLoc(warehouse, i, j);
        if (right === 'O') {
            boxes++;
        }
    } while (right === 'O');

    if (right === '.') {
        return boxes;
    } else {
        return -1;
    }
}

for (const m of movements) {
    let numBoxes;
    switch (m) {
        case '^':
            numBoxes = findNumBoxesToMoveUp(warehouseLayout, robot.i, robot.j);
            if (numBoxes === -1) {
                continue;
            } else if (numBoxes > 0) {
                warehouseLayout[robot.i - numBoxes - 1][robot.j] = 'O';
            }
            warehouseLayout[robot.i - 1][robot.j] = '@';
            warehouseLayout[robot.i][robot.j] = '.';
            robot.i = robot.i - 1;
            break;
        case 'v':
            numBoxes = findNumBoxesToMoveDown(warehouseLayout, robot.i, robot.j);
            if (numBoxes === -1) {
                continue;
            } else if (numBoxes > 0) {
                warehouseLayout[robot.i + numBoxes + 1][robot.j] = 'O';
            }
            warehouseLayout[robot.i + 1][robot.j] = '@';
            warehouseLayout[robot.i][robot.j] = '.';
            robot.i = robot.i + 1;
            break;
        case '<':
            numBoxes = findNumBoxesToMoveLeft(warehouseLayout, robot.i, robot.j);
            if (numBoxes === -1) {
                continue;
            } else if (numBoxes > 0) {
                warehouseLayout[robot.i][robot.j - numBoxes - 1] = 'O';
            }
            warehouseLayout[robot.i][robot.j - 1] = '@';
            warehouseLayout[robot.i][robot.j] = '.';
            robot.j = robot.j - 1;
            break;
        case '>':
            numBoxes = findNumBoxesToMoveRight(warehouseLayout, robot.i, robot.j);
            if (numBoxes === -1) {
                continue;
            } else if (numBoxes > 0) {
                warehouseLayout[robot.i][robot.j + numBoxes + 1] = 'O';
            }
            warehouseLayout[robot.i][robot.j + 1] = '@';
            warehouseLayout[robot.i][robot.j] = '.';
            robot.j = robot.j + 1;
            break;
    }
}

const wl1 = new Warehouse(warehouse);
for (const m of movements) {
    wl1.move(m);
}
console.log("Part 1", wl1.score);

const expandedWarehouse = warehouse
    .replaceAll('#', '##')
    .replaceAll('O', '[]')
    .replaceAll('.', '..')
    .replaceAll('@', '@.');

const wl2 = parseInput(expandedWarehouse, l => [...l]);
let boxes = [];

for (let i = 0; i < wl2.length; i++) {
    for (let j = 0; j < wl2[i].length; j++) {
        const loc = getLoc(wl2, i, j)
        if (loc === '@') {
            robot.i = i;
            robot.j = j;
        } else if (loc == '[') {
            boxes.push({
                left: { i, j },
                right: { i, j: j + 1 },
            });
        }
    }
}

const findNeighbourBoxesLeft = (grid, i, j) => {
    const loc = getLoc(grid, i, j - 1)
    if (loc == '.' || loc == '#') {
        return [];
    }

    const box = {
        left: { i, j: j - 2 },
        right: { i, j: j - 1 },
    };
    return [
        box,
        ...findNeighbourBoxesLeft(grid, box.left.i, box.left.j)
    ];
}

const findNeighbourBoxesRight = (grid, i, j) => {
    const loc = getLoc(grid, i, j + 1)
    if (loc == '.' || loc == '#') {
        return [];
    }

    const box = {
        left: { i, j: j + 1 },
        right: { i, j: j + 2 },
    };
    return [
        box,
        ...findNeighbourBoxesRight(grid, box.right.i, box.right.j)
    ];
}

const findNeighbourBoxesUp = (grid, i, j) => {
    let from = j, to = j;
    let boxes = [];

    while (i > 0) {
        i--;
        if (from === to && grid[i][j] === ']') {
            from--;
        }
        if (from === to && grid[i][j] === '[') {
            to++;
        }
        const above = grid[i].slice(from, to + 1).join('');
        if (above.includes('#')) {
            return {
                canMove: false,
                boxes
            }
        } else if (above.split('').every(c => c === '.')) {
            return {
                canMove: true,
                boxes
            }
        }

        if (grid[i][to] === '[') {
            to++;
        }
        if (grid[i][from] === ']') {
            from--;
        }
        while (grid[i][to] === '.') {
            to--;
        }

        while (grid[i][from] === '.') {
            from++;
        }


        for (let k = from; k < to; k++) {
            if (grid[i][k] === '[') {
                boxes.push({
                    left: { i, j: k },
                    right: { i, j: k + 1 }
                });
            }
        }
    }
}

const findNeighbourBoxesDown = (grid, i, j) => {
    let from = j, to = j;
    let boxes = [];

    while (i < grid.length - 1) {
        i++;
        if (from === to && grid[i][j] === ']') {
            from--;
        }
        if (from === to && grid[i][j] === '[') {
            to++;
        }
        const below = grid[i].slice(from, to + 1).join('');
        if (below.includes('#')) {
            return {
                canMove: false,
                boxes
            }
        } else if (below.split('').every(c => c === '.')) {
            return {
                canMove: true,
                boxes
            }
        }

        if (grid[i][to] === '[') {
            to++;
        }
        if (grid[i][from] === ']') {
            from--;
        }
        while (grid[i][to] === '.') {
            to--;
        }

        while (grid[i][from] === '.') {
            from++;
        }

        for (let k = from; k < to; k++) {
            if (grid[i][k] === '[') {
                boxes.push({
                    left: { i, j: k },
                    right: { i, j: k + 1 }
                });
            }
        }
    }
}


for (const m of movements.replace('\n', '')) {
    let canMove = false;
    let boxes = [];
    switch (m) {
        case '^':
            let { canMove: canMoveUp, boxes: boxesUp } = findNeighbourBoxesUp(wl2, robot.i, robot.j);
            if (canMoveUp) {
                boxesUp.sort((a, b) => a.left.i - b.left.i);
                for (const b of boxesUp) {
                    wl2[b.left.i - 1][b.left.j] = '[';
                    wl2[b.right.i - 1][b.right.j] = ']';
                    wl2[b.left.i][b.left.j] = '.';
                    wl2[b.right.i][b.right.j] = '.';
                }

                wl2[robot.i - 1][robot.j] = '@';
                wl2[robot.i][robot.j] = '.';
                robot.i = robot.i - 1;
            }
            break;
        case 'v':
            let { canMove: canMoveDown, boxes: boxesDown } = findNeighbourBoxesDown(wl2, robot.i, robot.j);
            if (canMoveDown) {
                boxesDown.sort((a, b) => b.left.i - a.left.i);
                for (const b of boxesDown) {
                    wl2[b.left.i + 1][b.left.j] = '[';
                    wl2[b.right.i + 1][b.right.j] = ']';
                    wl2[b.left.i][b.left.j] = '.';
                    wl2[b.right.i][b.right.j] = '.';
                }

                wl2[robot.i + 1][robot.j] = '@';
                wl2[robot.i][robot.j] = '.';
                robot.i = robot.i + 1;
            }
            break;
        case '<':
            boxes = findNeighbourBoxesLeft(wl2, robot.i, robot.j);
            canMove = getLoc(wl2, robot.i, robot.j - 2 * boxes.length - 1) === '.';
            if (canMove) {
                for (let i = 0; i < boxes.length; i++) {
                    boxes[i].left.j = boxes[i].left.j - 1;
                    boxes[i].right.j = boxes[i].right.j - 1;
                    wl2[boxes[i].left.i][boxes[i].left.j] = '[';
                    wl2[boxes[i].right.i][boxes[i].right.j] = ']';
                }
                wl2[robot.i][robot.j - 1] = '@';
                wl2[robot.i][robot.j] = '.';
                robot.j = robot.j - 1;
            }
            break;
        case '>':
            boxes = findNeighbourBoxesRight(wl2, robot.i, robot.j);
            canMove = getLoc(wl2, robot.i, robot.j + 2 * boxes.length + 1) === '.';
            if (canMove) {
                for (let i = 0; i < boxes.length; i++) {
                    boxes[i].left.j = boxes[i].left.j + 1;
                    boxes[i].right.j = boxes[i].right.j + 1;
                    wl2[boxes[i].left.i][boxes[i].left.j] = '[';
                    wl2[boxes[i].right.i][boxes[i].right.j] = ']';
                }
                wl2[robot.i][robot.j + 1] = '@';
                wl2[robot.i][robot.j] = '.';
                robot.j = robot.j + 1;

            }
            break;
    }
}

console.log("Part 2", getScore(wl2));
