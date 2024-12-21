import { input } from "./input.js";
import { parseInput } from "../utils.js";
import '../array.js';

const numPad = new Map([
    ['7', [0, 0]],
    ['8', [1, 0]],
    ['9', [2, 0]],
    ['4', [0, 1]],
    ['5', [1, 1]],
    ['6', [2, 1]],
    ['1', [0, 2]],
    ['2', [1, 2]],
    ['3', [2, 2]],
    ['0', [1, 3]],
    ['A', [2, 3]],
]);

const dirPad = new Map([
    ['^', [1, 0]],
    ['A', [2, 0]],
    ['<', [0, 1]],
    ['v', [1, 1]],
    ['>', [2, 1]],
]);

const dx = [0, -1, 1, 0];
const dy = [-1, 0, 0, 1];
const distance = (x1, y1, x2, y2) => Math.abs(x1 - x2) + Math.abs(y1 - y2);

const findNumSeqCache = new Map();
const findNumSeq = (from, to) => {
    const key = from + '->' + to;
    if (findNumSeqCache.has(key)) {
        return findNumSeqCache.get(key);
    }
    
    let [x1, y1] = numPad.get(from);
    const [x2, y2] = numPad.get(to);
    let moves = '';

    while (x1 != x2 || y1 != y2) {
        for (let i = 0; i < 4; i++) {
            const x = x1 + dx[i];
            const y = y1 + dy[i];
            if ((x == 0 && y == 3) || x < 0 || x > 2 || y < 0 || y > 3) {
                continue; // gap or outside pad
            }

            if (distance(x, y, x2, y2) < distance(x1, y1, x2, y2)) {
                x1 = x;
                y1 = y;
                if (dx[i] == -1) {
                    moves += '<';
                    break;
                }
                else if (dx[i] == 1) {
                    moves += '>';
                    break;
                }
                else if (dy[i] == -1) {
                    moves += '^';
                    break;
                }
                else if (dy[i] == 1) {
                    moves += 'v';
                    break;
                }
            }
        }
    }

    findNumSeqCache.set(key, moves + 'A');
    return moves + 'A';
}

const findDirSeqCache = new Map();
const findDirSeq = (from, to) => {
    const key = from + '->' + to;
    if (findDirSeqCache.has(key)) {
        return findDirSeqCache.get(key);
    }
    let [x1, y1] = dirPad.get(from);
    const [x2, y2] = dirPad.get(to);
    let moves = '';

    while (x1 != x2 || y1 != y2) {
        for (let i = 0; i < 4; i++) {
            const x = x1 + dx[i];
            const y = y1 + dy[i];
            if ((x == 0 && y == 0) || x < 0 || x > 2 || y < 0 || y > 1) {
                continue; // gap or outside pad
            }

            if (distance(x, y, x2, y2) < distance(x1, y1, x2, y2)) {
                x1 = x;
                y1 = y;
                if (dx[i] == -1) {
                    moves += '<';
                    break;
                }
                else if (dx[i] == 1) {
                    moves += '>';
                    break;
                }
                else if (dy[i] == -1) {
                    moves += '^';
                    break;
                }
                else if (dy[i] == 1) {
                    moves += 'v';
                    break;
                }
            }
        }
    }

    findDirSeqCache.set(key, moves + 'A');
    return moves + 'A';
}

const combinationsCache = new Map();
const getCombinations = (seq) => {
    if (seq.length == 0) { return [] };
    if (seq.length == 1) { return [seq] };

    if (combinationsCache.has(seq)) {
        return combinationsCache.get(seq);
    }

    let combinations = [];
    for (let i = 0; i < seq.length; i++) {
        let tmp = seq.substring(0, i) + seq.substring(i + 1);
        for (const c of getCombinations(tmp)) {
            combinations.push([
                seq[i], ...c
            ].join(''));
        }
    }

    combinationsCache.set(seq, [...new Set(combinations).keys()]);
    return combinationsCache.get(seq);
}

const isValidNumPath = (from, path) => {
    let [x, y] = numPad.get(from);
    for (let i = 0; i < path.length; i++) {
        switch (path[i]) {
            case '<':
                x--;
                break;
            case '>':
                x++;
                break;
            case '^':
                y--;
                break;
            case 'v':
                y++;
                break;
        }

        if (x == 0 && y == 3) {
            return false;
        }
    }

    return true;
}

const isValidDirPath = (from, path) => {
    let [x, y] = dirPad.get(from);
    for (let i = 0; i < path.length; i++) {
        switch (path[i]) {
            case '<':
                x--;
                break;
            case '>':
                x++;
                break;
            case '^':
                y--;
                break;
            case 'v':
                y++;
                break;
        }

        if (x == 0 && y == 0) {
            return false;
        }
    }

    return true;
}

const translatNumCache = new Map
const translateNum = (from, to) => {
    const key = from + "->" + to;
    if (translatNumCache.has(key)) {
        return translatNumCache.get(key);
    }
    let seq = findNumSeq(from, to);
    let translations = getCombinations(seq.substring(0, seq.length - 1))
        .filter(c => isValidNumPath(from, c));
    if (translations.length == 0) {
        translations = ['A'];
    } else {
        translations = translations.map(t => t + 'A');
    }

    translatNumCache.set(key, translations);
    return translations;
}

const translatDirCache = new Map
const translateDir = (from, to) => {
    const key = from + ":" + to;
    if (translatDirCache.has(key)) {
        return translatDirCache.get(key);
    }
    let seq = findDirSeq(from, to);
    let translations = getCombinations(seq.substring(0, seq.length - 1));
    translations = translations.filter(c => isValidDirPath(from, c));
    if (translations.length == 0) {
        translations = ['A'];
    } else {
        translations = translations.map(t => t + 'A');
    }
    translatDirCache.set(key, translations);
    return translations;
}

const translateNumToDir = l => {
    let curr = 'A';
    let moves = [];
    for (const c of l) {
        let translations = translateNum(curr, c);
        if (moves.length == 0) {
            moves = [...translations];
        } else {
            let newMoves = [];
            for (const m of moves) {
                for (const t of translations) {
                    newMoves.push(m + t);
                }
            }
            moves = newMoves;
        }
        curr = c;
    }

    return moves;
}

const getSteps = (moves) => {
    const steps = new Set();
    for (const m of moves) {
        for (let i = 0; i < m.length; i++) {
            if (i == 0) {
                steps.add('A' + m[i]);
            } else {
                steps.add(m[i - 1] + m[i]);
            }
        }
    }
    return steps;
}

const dfsCache = new Map();
const dfs = (from, to, level) => {
    const key = from + ':' + to + ':' + level;
    if (dfsCache.has(key)) {
        return dfsCache.get(key);
    }
    const dir = translateDir(from, to);
    if (level == 0) {
        const len = dir.map(d => d.length).min();
        dfsCache.set(key, len);
        return len;
    }

    const steps = getSteps(dir);
    const stepLens = new Map();
    for (const s of steps) {
        const len = dfs(s[0], s[1], level - 1);
        stepLens.set(s, len);
    }

    const minLen = dir.map(d => {
        let len = 0;
        for (let i = 0; i < d.length; i++) {
            if (i == 0) {
                len += stepLens.get('A' + d[i]);
            } else {
                len += stepLens.get(d[i - 1] + d[i]);
            }
        }
        return len;
    }).min();

    dfsCache.set(key, minLen);
    return minLen;
}

const getComplexity = (l, numDirPads) => {
    let moves = translateNumToDir(l);
    const minLen = moves.map(m => {
        let len = 0;
        for (let i = 0; i < m.length; i++) {
            if (i == 0) {
                len += dfs('A', m[i], numDirPads - 1);
            } else {
                len += dfs(m[i - 1], m[i], numDirPads - 1);
            }
        }
        return len;
    }).min();

    return minLen * Number(l.substring(0, l.length - 1));
}

console.log("Part 1", parseInput(input, l => getComplexity(l, 2)).sum());
console.log("Part 2", parseInput(input, l => getComplexity(l, 25)).sum());