import { input } from "./input.js";
import '../array.js';

const layouts = input.split('\n\n');
const locks = layouts.filter(l => l.startsWith('#####')).map(lock => {
    const lines = lock.split('\n').slice(1);
    const heights = new Array(5).fill(0);
    for (let i = 0; i < lines.length; i++) {
        for (let j = 0; j < heights.length; j++) {
            if (lines[i][j] == '#') {
                heights[j]++;
            }
        }
    }
    return heights;
});
const keys = layouts.filter(l => l.endsWith('#####')).map(key => {
    const lines = key.split('\n');
    const heights = new Array(5).fill(0);
    for (let i = lines.length - 2; i >= 0; i--) {
        for (let j = 0; j < heights.length; j++) {
            if (lines[i][j] == '#') {
                heights[j]++;
            }
        }
    }
    return heights;
});

let count = 0;
for (const lock of locks) {
    for (const key of keys) {
        if (lock.every((h, i) => key[i] + h <= 5)) {
            count++;
        }
    }
}

console.log(count);