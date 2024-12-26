import { skip } from "./input.js";

let buffer = [0];

for (let currIdx = 0, i = 0; i < 2017; i++) {
    currIdx = (currIdx + skip) % buffer.length + 1;
    buffer.splice(currIdx, 0, i + 1);
}

console.log("Part 1", buffer[buffer.findIndex(v => v == 2017) + 1]);

let val = 0;
let currIdx = 0;
let bufferSize = 1;
for (let i = 0; i < 49999999; i++) {
    currIdx = (currIdx + skip) % bufferSize + 1;
    if (currIdx == 1) {
        val = i + 1;
    }
    bufferSize++;
}

console.log("Part 2", val);