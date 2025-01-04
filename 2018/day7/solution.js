import { input } from "./input.js";
import { parseInput } from '../utils.js';
import { PriorityQueue } from '../minHeap.js';

const steps = parseInput(input, l => ([l[5], l[36]]));
const stepsMap = new Map();
const reverseStepsMap = new Map();
for (const [s1, s2] of steps) {
    if (!stepsMap.has(s1)) {
        stepsMap.set(s1, []);
    }

    if (!reverseStepsMap.has(s2)) {
        reverseStepsMap.set(s2, []);
    }

    stepsMap.set(s1, [
        ...stepsMap.get(s1),
        s2
    ].toSorted());

    reverseStepsMap.set(s2, [
        ...reverseStepsMap.get(s2),
        s1
    ].toSorted());
}
const starts = [...stepsMap.keys()].filter(s => ![...stepsMap.values()].some(v => v.includes(s)));
const queue = new PriorityQueue((a, b) => a < b);
let part1 = '';
for (const s of starts) {
    queue.push(s);
}
while (queue.size() > 0) {
    const s = queue.pop();
    if (part1.includes(s)) continue;
    part1 += s;
    const nextSteps = stepsMap.get(s) ?? [];
    const unfinished = [...stepsMap.keys()].filter(s => !part1.includes(s));
    for (const ns of nextSteps) {
        if (unfinished.some(s => stepsMap.get(s).includes(ns))) continue;
        queue.push(ns);
    }
}
console.log(part1);

let workers = [];
let finished = [];
let part2 = 0;
while (finished.length < part1.length) {
    for (const s of part1) {
        if (finished.includes(s)) continue;
        const needToFinish = reverseStepsMap.get(s) ?? [];
        const canStart = needToFinish.length == 0 || needToFinish.every(s => finished.includes(s));
        if (canStart && !workers.some(w => w.s == s) && workers.length < 5) {
            workers.push({ s, t: s.charCodeAt(0) - 4 });
        }
    }
    workers.sort((a, b) => a.t - b.t);
    const { s, t } = workers.shift();
    workers = workers.map(w => ({ ...w, t: w.t - t }));
    finished.push(s);
    part2 += t;
}
console.log(part2);