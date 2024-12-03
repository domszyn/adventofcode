import { input } from "./input.js";
import '../array.js';

const multiply = input => {
    const regex = /mul\((\d{1,3}),(\d{1,3})\)/g;

    let matches = input.matchAll(regex);
    return [...matches].map(([,left,right])=> left*right).sum();
}

console.log("Part 1", multiply(input));

const filteredInput = input
    .replaceAll("do()", "\ndo()")
    .replaceAll("don't()", "\ndon't()")
    .split('\n')
    .filter(s => !s.startsWith("don't()"))
    .join("");

console.log("Part 2", multiply(filteredInput));
